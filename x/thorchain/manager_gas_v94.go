package thorchain

import (
	"fmt"

	"gitlab.com/thorchain/thornode/common"
	"gitlab.com/thorchain/thornode/common/cosmos"
	"gitlab.com/thorchain/thornode/constants"
	"gitlab.com/thorchain/thornode/x/thorchain/keeper"
	"gitlab.com/thorchain/thornode/x/thorchain/types"
)

// GasMgrV94 implement GasManager interface which will store the gas related events happened in thorchain to memory
// emit GasEvent per block if there are any
type GasMgrV94 struct {
	gasEvent          *EventGas
	gas               common.Gas
	gasCount          map[common.Asset]int64
	constantsAccessor constants.ConstantValues
	keeper            keeper.Keeper
	mgr               Manager
}

// newGasMgrV94 create a new instance of GasMgrV1
func newGasMgrV94(constantsAccessor constants.ConstantValues, k keeper.Keeper) *GasMgrV94 {
	return &GasMgrV94{
		gasEvent:          NewEventGas(),
		gas:               common.Gas{},
		gasCount:          make(map[common.Asset]int64),
		constantsAccessor: constantsAccessor,
		keeper:            k,
	}
}

func (gm *GasMgrV94) reset() {
	gm.gasEvent = NewEventGas()
	gm.gas = common.Gas{}
	gm.gasCount = make(map[common.Asset]int64)
}

// BeginBlock need to be called when a new block get created , update the internal EventGas to new one
func (gm *GasMgrV94) BeginBlock(mgr Manager) {
	gm.mgr = mgr
	gm.reset()
}

// AddGasAsset to the EventGas
func (gm *GasMgrV94) AddGasAsset(_ common.Asset, gas common.Gas, increaseTxCount bool) {
	gm.gas = gm.gas.Adds_deprecated(gas)
	if !increaseTxCount {
		return
	}
	for _, coin := range gas {
		gm.gasCount[coin.Asset]++
	}
}

// GetGas return gas
func (gm *GasMgrV94) GetGas() common.Gas {
	return gm.gas
}

// GetFee retrieve the network fee information from kv store, and calculate the dynamic fee customer should pay
// the return value is the amount of fee in RUNE
func (gm *GasMgrV94) GetFee(ctx cosmos.Context, chain common.Chain, asset common.Asset) cosmos.Uint {
	outboundTxFee, err := gm.keeper.GetMimir(ctx, constants.OutboundTransactionFee.String())
	if outboundTxFee < 0 || err != nil {
		outboundTxFee = gm.constantsAccessor.GetInt64Value(constants.OutboundTransactionFee)
	}
	transactionFee := cosmos.NewUint(uint64(outboundTxFee))
	// if the asset is Native RUNE , then we could just return the transaction Fee
	// because transaction fee is always in native RUNE
	if asset.IsRune() && chain.Equals(common.THORChain) {
		return transactionFee
	}

	// if the asset is synthetic asset , it need to get the layer 1 asset pool and convert it
	// synthetic asset live on THORChain , thus it doesn't need to get the layer1 network fee
	if asset.IsSyntheticAsset() {
		return gm.getRuneInAssetValue(ctx, transactionFee, asset)
	}

	networkFee, err := gm.keeper.GetNetworkFee(ctx, chain)
	if err != nil {
		ctx.Logger().Error("fail to get network fee", "error", err)
		return transactionFee
	}
	if err := networkFee.Valid(); err != nil {
		ctx.Logger().Error("network fee is invalid", "error", err, "chain", chain)
		return transactionFee
	}

	pool, err := gm.keeper.GetPool(ctx, chain.GetGasAsset())
	if err != nil {
		ctx.Logger().Error("fail to get pool", "asset", asset, "error", err)
		return transactionFee
	}

	minOutboundUSD, err := gm.keeper.GetMimir(ctx, constants.MinimumL1OutboundFeeUSD.String())
	if minOutboundUSD < 0 || err != nil {
		minOutboundUSD = gm.constantsAccessor.GetInt64Value(constants.MinimumL1OutboundFeeUSD)
	}
	oneDollarRune := cosmos.ZeroUint()
	// since gm.mgr get set at BeginBlock , so here add a safeguard in case gm.mgr is nil
	if gm.mgr != nil {
		oneDollarRune = DollarInRune(ctx, gm.mgr)
	}
	minAsset := cosmos.ZeroUint()
	if !oneDollarRune.IsZero() {
		// since MinOutboundUSD is in USD value , thus need to figure out how much RUNE
		// here use GetShare instead GetSafeShare it is because minOutboundUSD can set to more than $1
		minOutboundInRune := common.GetUncappedShare(cosmos.NewUint(uint64(minOutboundUSD)),
			cosmos.NewUint(common.One),
			oneDollarRune)

		minAsset = pool.RuneValueInAsset(minOutboundInRune)
	}
	// Fee is calculated based on the network fee observed in previous block
	// THORNode is going to charge 3 times the fee it takes to send out the tx
	// 1.5 * fee will goes to vault
	// 1.5 * fee will become the max gas used to send out the tx
	fee := cosmos.RoundToDecimal(
		cosmos.NewUint(networkFee.TransactionSize*networkFee.TransactionFeeRate*3),
		pool.Decimals,
	)

	// Ensure fee is always more than minAsset
	if fee.LT(minAsset) {
		fee = minAsset
	}

	if asset.Equals(asset.GetChain().GetGasAsset()) && chain.Equals(asset.GetChain()) {
		return fee
	}

	// convert gas asset value into rune
	if pool.BalanceAsset.Equal(cosmos.ZeroUint()) || pool.BalanceRune.Equal(cosmos.ZeroUint()) {
		return transactionFee
	}

	fee = pool.AssetValueInRune(fee)
	if asset.IsRune() {
		return fee
	}

	// convert rune value into non-gas asset value
	pool, err = gm.keeper.GetPool(ctx, asset)
	if err != nil {
		ctx.Logger().Error("fail to get pool", "asset", asset, "error", err)
		return transactionFee
	}
	if pool.BalanceAsset.Equal(cosmos.ZeroUint()) || pool.BalanceRune.Equal(cosmos.ZeroUint()) {
		return transactionFee
	}
	return pool.RuneValueInAsset(fee)
}

// getRuneInAssetValue convert the transaction fee to asset value , when the given asset is synthetic , it will need to get
// the layer1 asset first , and then use the pool to convert
func (gm *GasMgrV94) getRuneInAssetValue(ctx cosmos.Context, transactionFee cosmos.Uint, asset common.Asset) cosmos.Uint {
	if asset.IsSyntheticAsset() {
		asset = asset.GetLayer1Asset()
	}
	pool, err := gm.keeper.GetPool(ctx, asset)
	if err != nil {
		ctx.Logger().Error("fail to get pool", "asset", asset, "error", err)
		return transactionFee
	}
	if pool.BalanceAsset.Equal(cosmos.ZeroUint()) || pool.BalanceRune.Equal(cosmos.ZeroUint()) {
		return transactionFee
	}

	return pool.RuneValueInAsset(transactionFee)
}

// GetGasRate return the gas rate
func (gm *GasMgrV94) GetGasRate(ctx cosmos.Context, chain common.Chain) cosmos.Uint {
	outboundTxFee, err := gm.keeper.GetMimir(ctx, constants.OutboundTransactionFee.String())
	if outboundTxFee < 0 || err != nil {
		outboundTxFee = gm.constantsAccessor.GetInt64Value(constants.OutboundTransactionFee)
	}
	transactionFee := cosmos.NewUint(uint64(outboundTxFee))
	if chain.Equals(common.THORChain) {
		return transactionFee
	}
	networkFee, err := gm.keeper.GetNetworkFee(ctx, chain)
	if err != nil {
		ctx.Logger().Error("fail to get network fee", "error", err)
		return transactionFee
	}
	if err := networkFee.Valid(); err != nil {
		ctx.Logger().Error("network fee is invalid", "error", err, "chain", chain)
		return transactionFee
	}
	return cosmos.RoundToDecimal(
		cosmos.NewUint(networkFee.TransactionFeeRate*3/2),
		chain.GetGasAssetDecimal(),
	)
}

func (gm *GasMgrV94) GetNetworkFee(ctx cosmos.Context, chain common.Chain) (types.NetworkFee, error) {
	outboundTxFee, err := gm.keeper.GetMimir(ctx, constants.OutboundTransactionFee.String())
	if outboundTxFee < 0 || err != nil {
		outboundTxFee = gm.constantsAccessor.GetInt64Value(constants.OutboundTransactionFee)
	}
	transactionFee := cosmos.NewUint(uint64(outboundTxFee))
	if chain.Equals(common.THORChain) {
		return types.NewNetworkFee(common.THORChain, 1, transactionFee.Uint64()), nil
	}

	return gm.keeper.GetNetworkFee(ctx, chain)
}

// GetMaxGas will calculate the maximum gas fee a tx can use
func (gm *GasMgrV94) GetMaxGas(ctx cosmos.Context, chain common.Chain) (common.Coin, error) {
	gasAsset := chain.GetGasAsset()
	var amount cosmos.Uint

	nf, err := gm.keeper.GetNetworkFee(ctx, chain)
	if err != nil {
		return common.NoCoin, fmt.Errorf("fail to get network fee for chain(%s): %w", chain, err)
	}
	if chain.IsBNB() {
		amount = cosmos.NewUint(nf.TransactionSize * nf.TransactionFeeRate)
	} else {
		amount = cosmos.NewUint(nf.TransactionSize * nf.TransactionFeeRate).MulUint64(3).QuoUint64(2)
	}
	gasCoin := common.NewCoin(gasAsset, amount)
	chainGasAssetPrecision := chain.GetGasAssetDecimal()
	gasCoin.Amount = cosmos.RoundToDecimal(amount, chainGasAssetPrecision)
	gasCoin.Decimals = chainGasAssetPrecision
	return gasCoin, nil
}

// EndBlock emit the events
func (gm *GasMgrV94) EndBlock(ctx cosmos.Context, keeper keeper.Keeper, eventManager EventManager) {
	gm.ProcessGas(ctx, keeper)

	if len(gm.gasEvent.Pools) == 0 {
		return
	}
	if err := eventManager.EmitGasEvent(ctx, gm.gasEvent); nil != err {
		ctx.Logger().Error("fail to emit gas event", "error", err)
	}
	gm.reset() // do not remove, will cause consensus failures
}

// ProcessGas to subsidise the pool with RUNE for the gas they have spent
func (gm *GasMgrV94) ProcessGas(ctx cosmos.Context, keeper keeper.Keeper) {
	if keeper.RagnarokInProgress(ctx) {
		// ragnarok is in progress , stop
		return
	}
	vault, err := keeper.GetNetwork(ctx)
	if err != nil {
		ctx.Logger().Error("fail to get network data", "error", err)
		return
	}
	for _, gas := range gm.gas {
		// if the coin is zero amount, don't need to do anything
		if gas.Amount.IsZero() {
			continue
		}

		pool, err := keeper.GetPool(ctx, gas.Asset)
		if err != nil {
			ctx.Logger().Error("fail to get pool", "pool", gas.Asset, "error", err)
			continue
		}
		if err := pool.Valid(); err != nil {
			ctx.Logger().Error("invalid pool", "pool", gas.Asset, "error", err)
			continue
		}
		runeGas := pool.AssetValueInRune(gas.Amount) // Convert to Rune (gas will never be RUNE)
		if runeGas.IsZero() {
			continue
		}
		// If Rune owed now exceeds the Total Reserve, return it all
		if runeGas.LT(keeper.GetRuneBalanceOfModule(ctx, ReserveName)) {
			coin := common.NewCoin(common.RuneNative, runeGas)
			if err := keeper.SendFromModuleToModule(ctx, ReserveName, AsgardName, common.NewCoins(coin)); err != nil {
				ctx.Logger().Error("fail to transfer funds from reserve to asgard", "pool", gas.Asset, "error", err)
				continue
			}
			pool.BalanceRune = pool.BalanceRune.Add(runeGas) // Add to the pool
		} else {
			// since we don't have enough in the reserve to cover the gas used,
			// no rune is added to the pool, sorry LPs!
			runeGas = cosmos.ZeroUint()
		}
		pool.BalanceAsset = common.SafeSub(pool.BalanceAsset, gas.Amount)

		if err := keeper.SetPool(ctx, pool); err != nil {
			ctx.Logger().Error("fail to set pool", "pool", gas.Asset, "error", err)
			continue
		}

		gasPool := GasPool{
			Asset:    gas.Asset,
			AssetAmt: gas.Amount,
			RuneAmt:  runeGas,
			Count:    gm.gasCount[gas.Asset],
		}
		gm.gasEvent.UpsertGasPool(gasPool)
	}

	if err := keeper.SetNetwork(ctx, vault); err != nil {
		ctx.Logger().Error("fail to set network data", "error", err)
	}
}

func (gm *GasMgrV94) CalcOutboundFeeMultiplier(ctx cosmos.Context, targetSurplusRune, gasSpentRune, gasWithheldRune, maxMultiplier, minMultiplier cosmos.Uint) cosmos.Uint {
	return cosmos.ZeroUint()
}

func (gm *GasMgrV94) GetAssetOutboundFee(ctx cosmos.Context, asset common.Asset, inRune bool) (cosmos.Uint, error) {
	return cosmos.ZeroUint(), nil
}
