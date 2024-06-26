package thorchain

import (
	"errors"
	"fmt"
	"testing"

	"github.com/magiconair/properties/assert"
	. "gopkg.in/check.v1"

	"gitlab.com/thorchain/thornode/common"
	"gitlab.com/thorchain/thornode/common/cosmos"
	"gitlab.com/thorchain/thornode/constants"
	"gitlab.com/thorchain/thornode/x/thorchain/keeper"
	"gitlab.com/thorchain/thornode/x/thorchain/types"
)

type WithdrawSuiteV91 struct{}

var _ = Suite(&WithdrawSuiteV91{})

type WithdrawTestKeeperV91 struct {
	keeper.KVStoreDummy
	store       map[string]interface{}
	networkFees map[common.Chain]NetworkFee
	keeper      keeper.Keeper
}

func NewWithdrawTestKeeperV91(keeper keeper.Keeper) *WithdrawTestKeeperV91 {
	return &WithdrawTestKeeperV91{
		keeper:      keeper,
		store:       make(map[string]interface{}),
		networkFees: make(map[common.Chain]NetworkFee),
	}
}

// this one has an extra liquidity provider already set
func getWithdrawTestKeeper2V91(c *C, ctx cosmos.Context, k keeper.Keeper, runeAddress common.Address) keeper.Keeper {
	store := NewWithdrawTestKeeperV91(k)
	pool := Pool{
		BalanceRune:  cosmos.NewUint(100 * common.One),
		BalanceAsset: cosmos.NewUint(100 * common.One),
		Asset:        common.BNBAsset,
		LPUnits:      cosmos.NewUint(200 * common.One),
		SynthUnits:   cosmos.ZeroUint(),
		Status:       PoolAvailable,
	}
	c.Assert(store.SetPool(ctx, pool), IsNil)
	lp := LiquidityProvider{
		Asset:        pool.Asset,
		RuneAddress:  runeAddress,
		AssetAddress: runeAddress,
		Units:        cosmos.NewUint(100 * common.One),
		PendingRune:  cosmos.ZeroUint(),
	}
	store.SetLiquidityProvider(ctx, lp)
	return store
}

func (k *WithdrawTestKeeperV91) PoolExist(ctx cosmos.Context, asset common.Asset) bool {
	return !asset.Equals(common.Asset{Chain: common.BNBChain, Symbol: "NOTEXIST", Ticker: "NOTEXIST"})
}

func (k *WithdrawTestKeeperV91) GetPool(ctx cosmos.Context, asset common.Asset) (types.Pool, error) {
	if asset.Equals(common.Asset{Chain: common.BNBChain, Symbol: "NOTEXIST", Ticker: "NOTEXIST"}) {
		return types.Pool{}, nil
	}
	if val, ok := k.store[asset.String()]; ok {
		p, _ := val.(types.Pool)
		return p, nil
	}
	return types.Pool{
		BalanceRune:  cosmos.NewUint(100).MulUint64(common.One),
		BalanceAsset: cosmos.NewUint(100).MulUint64(common.One),
		LPUnits:      cosmos.NewUint(100).MulUint64(common.One),
		SynthUnits:   cosmos.ZeroUint(),
		Status:       PoolAvailable,
		Asset:        asset,
	}, nil
}

func (k *WithdrawTestKeeperV91) SetPool(ctx cosmos.Context, ps Pool) error {
	k.store[ps.Asset.String()] = ps
	return nil
}

func (k *WithdrawTestKeeperV91) GetGas(ctx cosmos.Context, asset common.Asset) ([]cosmos.Uint, error) {
	return []cosmos.Uint{cosmos.NewUint(37500), cosmos.NewUint(30000)}, nil
}

func (k *WithdrawTestKeeperV91) GetLiquidityProvider(ctx cosmos.Context, asset common.Asset, addr common.Address) (LiquidityProvider, error) {
	if asset.Equals(common.Asset{Chain: common.BNBChain, Symbol: "NOTEXISTSTICKER", Ticker: "NOTEXISTSTICKER"}) {
		return types.LiquidityProvider{}, errors.New("you asked for it")
	}
	if notExistLiquidityProviderAsset.Equals(asset) {
		return LiquidityProvider{}, errors.New("simulate error for test")
	}
	return k.keeper.GetLiquidityProvider(ctx, asset, addr)
}

func (k *WithdrawTestKeeperV91) GetNetworkFee(ctx cosmos.Context, chain common.Chain) (NetworkFee, error) {
	return k.networkFees[chain], nil
}

func (k *WithdrawTestKeeperV91) SaveNetworkFee(ctx cosmos.Context, chain common.Chain, networkFee NetworkFee) error {
	k.networkFees[chain] = networkFee
	return nil
}

func (k *WithdrawTestKeeperV91) SetLiquidityProvider(ctx cosmos.Context, lp LiquidityProvider) {
	k.keeper.SetLiquidityProvider(ctx, lp)
}

func (s *WithdrawSuiteV91) SetUpSuite(c *C) {
	SetupConfigForTest()
}

// TestValidateWithdraw is to test validateWithdraw function
func (s WithdrawSuiteV91) TestValidateWithdraw(c *C) {
	accountAddr := GetRandomValidatorNode(NodeWhiteListed).NodeAddress
	runeAddress, err := common.NewAddress("bnb1g0xakzh03tpa54khxyvheeu92hwzypkdce77rm")
	if err != nil {
		c.Error("fail to create new BNB Address")
	}
	inputs := []struct {
		name          string
		msg           MsgWithdrawLiquidity
		expectedError error
	}{
		{
			name: "empty-rune-address",
			msg: MsgWithdrawLiquidity{
				WithdrawAddress: "",
				BasisPoints:     cosmos.NewUint(10000),
				Asset:           common.BNBAsset,
				Tx:              common.Tx{ID: "28B40BF105A112389A339A64BD1A042E6140DC9082C679586C6CF493A9FDE3FE"},
				Signer:          accountAddr,
			},
			expectedError: errors.New("empty withdraw address"),
		},
		{
			name: "empty-withdraw-basis-points",
			msg: MsgWithdrawLiquidity{
				WithdrawAddress: runeAddress,
				BasisPoints:     cosmos.ZeroUint(),
				Asset:           common.BNBAsset,
				Tx:              common.Tx{ID: "28B40BF105A112389A339A64BD1A042E6140DC9082C679586C6CF493A9FDE3FE"},
				Signer:          accountAddr,
			},
			expectedError: errors.New("withdraw basis points 0 is invalid"),
		},
		{
			name: "empty-request-txhash",
			msg: MsgWithdrawLiquidity{
				WithdrawAddress: runeAddress,
				BasisPoints:     cosmos.NewUint(10000),
				Asset:           common.BNBAsset,
				Tx:              common.Tx{},
				Signer:          accountAddr,
			},
			expectedError: errors.New("request tx hash is empty"),
		},
		{
			name: "empty-asset",
			msg: MsgWithdrawLiquidity{
				WithdrawAddress: runeAddress,
				BasisPoints:     cosmos.NewUint(10000),
				Asset:           common.Asset{},
				Tx:              common.Tx{ID: "28B40BF105A112389A339A64BD1A042E6140DC9082C679586C6CF493A9FDE3FE"},
				Signer:          accountAddr,
			},
			expectedError: errors.New("empty asset"),
		},
		{
			name: "invalid-basis-point",
			msg: MsgWithdrawLiquidity{
				WithdrawAddress: runeAddress,
				BasisPoints:     cosmos.NewUint(10001),
				Asset:           common.BNBAsset,
				Tx:              common.Tx{ID: "28B40BF105A112389A339A64BD1A042E6140DC9082C679586C6CF493A9FDE3FE"},
				Signer:          accountAddr,
			},
			expectedError: errors.New("withdraw basis points 10001 is invalid"),
		},
		{
			name: "invalid-pool-notexist",
			msg: MsgWithdrawLiquidity{
				WithdrawAddress: runeAddress,
				BasisPoints:     cosmos.NewUint(10000),
				Asset:           common.Asset{Chain: common.BNBChain, Ticker: "NOTEXIST", Symbol: "NOTEXIST"},
				Tx:              common.Tx{ID: "28B40BF105A112389A339A64BD1A042E6140DC9082C679586C6CF493A9FDE3FE"},
				Signer:          accountAddr,
			},
			expectedError: errors.New("pool-BNB.NOTEXIST doesn't exist"),
		},
		{
			name: "all-good",
			msg: MsgWithdrawLiquidity{
				WithdrawAddress: runeAddress,
				BasisPoints:     cosmos.NewUint(10000),
				Asset:           common.BNBAsset,
				Tx:              common.Tx{ID: "28B40BF105A112389A339A64BD1A042E6140DC9082C679586C6CF493A9FDE3FE"},
				Signer:          accountAddr,
			},
			expectedError: nil,
		},
	}

	for _, item := range inputs {
		ctx, _ := setupKeeperForTest(c)
		ps := &WithdrawTestKeeperV91{}
		c.Logf("name:%s", item.name)
		err := validateWithdrawV1(ctx, ps, item.msg)
		if item.expectedError != nil {
			c.Assert(err, NotNil)
			c.Assert(err.Error(), Equals, item.expectedError.Error())
			continue
		}
		c.Assert(err, IsNil)
	}
}

func (s WithdrawSuiteV91) TestCalculateUnsake(c *C) {
	inputs := []struct {
		name                  string
		poolUnit              cosmos.Uint
		poolRune              cosmos.Uint
		poolAsset             cosmos.Uint
		lpUnit                cosmos.Uint
		percentage            cosmos.Uint
		expectedWithdrawRune  cosmos.Uint
		expectedWithdrawAsset cosmos.Uint
		expectedUnitLeft      cosmos.Uint
		expectedErr           error
	}{
		{
			name:                  "zero-poolunit",
			poolUnit:              cosmos.ZeroUint(),
			poolRune:              cosmos.ZeroUint(),
			poolAsset:             cosmos.ZeroUint(),
			lpUnit:                cosmos.ZeroUint(),
			percentage:            cosmos.ZeroUint(),
			expectedWithdrawRune:  cosmos.ZeroUint(),
			expectedWithdrawAsset: cosmos.ZeroUint(),
			expectedUnitLeft:      cosmos.ZeroUint(),
			expectedErr:           errors.New("poolUnits can't be zero"),
		},

		{
			name:                  "zero-poolrune",
			poolUnit:              cosmos.NewUint(500 * common.One),
			poolRune:              cosmos.ZeroUint(),
			poolAsset:             cosmos.ZeroUint(),
			lpUnit:                cosmos.ZeroUint(),
			percentage:            cosmos.ZeroUint(),
			expectedWithdrawRune:  cosmos.ZeroUint(),
			expectedWithdrawAsset: cosmos.ZeroUint(),
			expectedUnitLeft:      cosmos.ZeroUint(),
			expectedErr:           errors.New("pool rune balance can't be zero"),
		},

		{
			name:                  "zero-poolasset",
			poolUnit:              cosmos.NewUint(500 * common.One),
			poolRune:              cosmos.NewUint(500 * common.One),
			poolAsset:             cosmos.ZeroUint(),
			lpUnit:                cosmos.ZeroUint(),
			percentage:            cosmos.ZeroUint(),
			expectedWithdrawRune:  cosmos.ZeroUint(),
			expectedWithdrawAsset: cosmos.ZeroUint(),
			expectedUnitLeft:      cosmos.ZeroUint(),
			expectedErr:           errors.New("pool asset balance can't be zero"),
		},
		{
			name:                  "negative-liquidity-provider-unit",
			poolUnit:              cosmos.NewUint(500 * common.One),
			poolRune:              cosmos.NewUint(500 * common.One),
			poolAsset:             cosmos.NewUint(5100 * common.One),
			lpUnit:                cosmos.ZeroUint(),
			percentage:            cosmos.ZeroUint(),
			expectedWithdrawRune:  cosmos.ZeroUint(),
			expectedWithdrawAsset: cosmos.ZeroUint(),
			expectedUnitLeft:      cosmos.ZeroUint(),
			expectedErr:           errors.New("liquidity provider unit can't be zero"),
		},

		{
			name:                  "percentage-larger-than-100",
			poolUnit:              cosmos.NewUint(500 * common.One),
			poolRune:              cosmos.NewUint(500 * common.One),
			poolAsset:             cosmos.NewUint(500 * common.One),
			lpUnit:                cosmos.NewUint(100 * common.One),
			percentage:            cosmos.NewUint(12000),
			expectedWithdrawRune:  cosmos.ZeroUint(),
			expectedWithdrawAsset: cosmos.ZeroUint(),
			expectedUnitLeft:      cosmos.ZeroUint(),
			expectedErr:           fmt.Errorf("withdraw basis point %s is not valid", cosmos.NewUint(12000)),
		},
		{
			name:                  "withdraw-1",
			poolUnit:              cosmos.NewUint(700 * common.One),
			poolRune:              cosmos.NewUint(700 * common.One),
			poolAsset:             cosmos.NewUint(700 * common.One),
			lpUnit:                cosmos.NewUint(200 * common.One),
			percentage:            cosmos.NewUint(10000),
			expectedUnitLeft:      cosmos.ZeroUint(),
			expectedWithdrawAsset: cosmos.NewUint(200 * common.One),
			expectedWithdrawRune:  cosmos.NewUint(200 * common.One),
			expectedErr:           nil,
		},
		{
			name:                  "withdraw-2",
			poolUnit:              cosmos.NewUint(100),
			poolRune:              cosmos.NewUint(15 * common.One),
			poolAsset:             cosmos.NewUint(155 * common.One),
			lpUnit:                cosmos.NewUint(100),
			percentage:            cosmos.NewUint(1000),
			expectedUnitLeft:      cosmos.NewUint(90),
			expectedWithdrawAsset: cosmos.NewUint(1550000000),
			expectedWithdrawRune:  cosmos.NewUint(150000000),
			expectedErr:           nil,
		},
	}

	for _, item := range inputs {
		c.Logf("name:%s", item.name)
		withDrawRune, withDrawAsset, unitAfter, err := calculateWithdrawV76(item.poolUnit, item.poolRune, item.poolAsset, item.lpUnit, cosmos.ZeroUint(), item.percentage, common.EmptyAsset)
		if item.expectedErr == nil {
			c.Assert(err, IsNil)
		} else {
			c.Assert(err.Error(), Equals, item.expectedErr.Error())
		}
		c.Logf("expected rune:%s,rune:%s", item.expectedWithdrawRune, withDrawRune)
		c.Check(item.expectedWithdrawRune.Uint64(), Equals, withDrawRune.Uint64(), Commentf("Expected %d, got %d", item.expectedWithdrawRune.Uint64(), withDrawRune.Uint64()))
		c.Check(item.expectedWithdrawAsset.Uint64(), Equals, withDrawAsset.Uint64(), Commentf("Expected %d, got %d", item.expectedWithdrawAsset.Uint64(), withDrawAsset.Uint64()))
		c.Check(item.expectedUnitLeft.Uint64(), Equals, unitAfter.Uint64())
	}
}

func (WithdrawSuiteV91) TestWithdraw(c *C) {
	ctx, mgr := setupManagerForTest(c)
	accountAddr := GetRandomValidatorNode(NodeWhiteListed).NodeAddress
	runeAddress := GetRandomRUNEAddress()
	ps := NewWithdrawTestKeeperV91(mgr.Keeper())
	ps2 := getWithdrawTestKeeperV91(c, ctx, mgr.Keeper(), runeAddress)

	remainGas := uint64(37500)
	testCases := []struct {
		name          string
		msg           MsgWithdrawLiquidity
		ps            keeper.Keeper
		runeAmount    cosmos.Uint
		assetAmount   cosmos.Uint
		expectedError error
	}{
		{
			name: "empty-rune-address",
			msg: MsgWithdrawLiquidity{
				WithdrawAddress: "",
				BasisPoints:     cosmos.NewUint(10000),
				Asset:           common.BNBAsset,
				Tx:              common.Tx{ID: "28B40BF105A112389A339A64BD1A042E6140DC9082C679586C6CF493A9FDE3FE"},
				Signer:          accountAddr,
			},
			ps:            ps,
			runeAmount:    cosmos.ZeroUint(),
			assetAmount:   cosmos.ZeroUint(),
			expectedError: errors.New("empty withdraw address"),
		},
		{
			name: "empty-request-txhash",
			msg: MsgWithdrawLiquidity{
				WithdrawAddress: runeAddress,
				BasisPoints:     cosmos.NewUint(10000),
				Asset:           common.BNBAsset,
				Tx:              common.Tx{},
				Signer:          accountAddr,
			},
			ps:            ps,
			runeAmount:    cosmos.ZeroUint(),
			assetAmount:   cosmos.ZeroUint(),
			expectedError: errors.New("request tx hash is empty"),
		},
		{
			name: "empty-asset",
			msg: MsgWithdrawLiquidity{
				WithdrawAddress: runeAddress,
				BasisPoints:     cosmos.NewUint(10000),
				Asset:           common.Asset{},
				Tx:              common.Tx{ID: "28B40BF105A112389A339A64BD1A042E6140DC9082C679586C6CF493A9FDE3FE"},
				Signer:          accountAddr,
			},
			ps:            ps,
			runeAmount:    cosmos.ZeroUint(),
			assetAmount:   cosmos.ZeroUint(),
			expectedError: errors.New("empty asset"),
		},
		{
			name: "invalid-basis-point",
			msg: MsgWithdrawLiquidity{
				WithdrawAddress: runeAddress,
				BasisPoints:     cosmos.NewUint(10001),
				Asset:           common.BNBAsset,
				Tx:              common.Tx{ID: "28B40BF105A112389A339A64BD1A042E6140DC9082C679586C6CF493A9FDE3FE"},
				Signer:          accountAddr,
			},
			ps:            ps,
			runeAmount:    cosmos.ZeroUint(),
			assetAmount:   cosmos.ZeroUint(),
			expectedError: errors.New("withdraw basis points 10001 is invalid"),
		},
		{
			name: "invalid-pool-notexist",
			msg: MsgWithdrawLiquidity{
				WithdrawAddress: runeAddress,
				BasisPoints:     cosmos.NewUint(10000),
				Asset:           common.Asset{Chain: common.BNBChain, Ticker: "NOTEXIST", Symbol: "NOTEXIST"},
				Tx:              common.Tx{ID: "28B40BF105A112389A339A64BD1A042E6140DC9082C679586C6CF493A9FDE3FE"},
				Signer:          accountAddr,
			},
			ps:            ps,
			runeAmount:    cosmos.ZeroUint(),
			assetAmount:   cosmos.ZeroUint(),
			expectedError: errors.New("pool-BNB.NOTEXIST doesn't exist"),
		},
		{
			name: "invalid-pool-liquidity-provider-notexist",
			msg: MsgWithdrawLiquidity{
				WithdrawAddress: runeAddress,
				BasisPoints:     cosmos.NewUint(10000),
				Asset:           common.Asset{Chain: common.BNBChain, Ticker: "NOTEXISTSTICKER", Symbol: "NOTEXISTSTICKER"},
				Tx:              common.Tx{ID: "28B40BF105A112389A339A64BD1A042E6140DC9082C679586C6CF493A9FDE3FE"},
				Signer:          accountAddr,
			},
			ps:            ps,
			runeAmount:    cosmos.ZeroUint(),
			assetAmount:   cosmos.ZeroUint(),
			expectedError: errors.New("you asked for it"),
		},
		{
			name: "nothing-to-withdraw",
			msg: MsgWithdrawLiquidity{
				WithdrawAddress: runeAddress,
				BasisPoints:     cosmos.ZeroUint(),
				Asset:           common.BNBAsset,
				Tx:              common.Tx{ID: "28B40BF105A112389A339A64BD1A042E6140DC9082C679586C6CF493A9FDE3FE"},
				Signer:          accountAddr,
			},
			ps:            ps,
			runeAmount:    cosmos.ZeroUint(),
			assetAmount:   cosmos.ZeroUint(),
			expectedError: errors.New("withdraw basis points 0 is invalid"),
		},
		{
			name: "all-good-half",
			msg: MsgWithdrawLiquidity{
				WithdrawAddress: runeAddress,
				BasisPoints:     cosmos.NewUint(5000),
				Asset:           common.BNBAsset,
				Tx:              common.Tx{ID: "28B40BF105A112389A339A64BD1A042E6140DC9082C679586C6CF493A9FDE3FE"},
				Signer:          accountAddr,
			},
			ps:            ps2,
			runeAmount:    cosmos.NewUint(50 * common.One),
			assetAmount:   cosmos.NewUint(50 * common.One),
			expectedError: nil,
		},
		{
			name: "all-good",
			msg: MsgWithdrawLiquidity{
				WithdrawAddress: runeAddress,
				BasisPoints:     cosmos.NewUint(10000),
				Asset:           common.BNBAsset,
				Tx:              common.Tx{ID: "28B40BF105A112389A339A64BD1A042E6140DC9082C679586C6CF493A9FDE3FE"},
				Signer:          accountAddr,
			},
			ps:            ps2,
			runeAmount:    cosmos.NewUint(50 * common.One),
			assetAmount:   cosmos.NewUint(50 * common.One).Sub(cosmos.NewUint(remainGas)),
			expectedError: nil,
		},
	}
	for _, tc := range testCases {
		c.Logf("name:%s", tc.name)
		mgr.K = tc.ps
		c.Assert(tc.ps.SaveNetworkFee(ctx, common.BNBChain, NetworkFee{
			Chain:              common.BNBChain,
			TransactionSize:    1,
			TransactionFeeRate: bnbSingleTxFee.Uint64(),
		}), IsNil)
		r, asset, _, _, _, err := withdrawV91(ctx, tc.msg, mgr)
		if tc.expectedError != nil {
			c.Assert(err, NotNil)
			c.Check(err.Error(), Equals, tc.expectedError.Error())
			c.Check(r.Uint64(), Equals, tc.runeAmount.Uint64())
			c.Check(asset.Uint64(), Equals, tc.assetAmount.Uint64())
			continue
		}
		c.Assert(err, IsNil)
		c.Assert(r.Uint64(), Equals, tc.runeAmount.Uint64(), Commentf("%d != %d", r.Uint64(), tc.runeAmount.Uint64()))
		c.Assert(asset.Equal(tc.assetAmount), Equals, true, Commentf("expect:%s, however got:%s", tc.assetAmount.String(), asset.String()))
	}
}

func (WithdrawSuiteV91) TestWithdrawAsym(c *C) {
	accountAddr := GetRandomValidatorNode(NodeWhiteListed).NodeAddress
	runeAddress := GetRandomRUNEAddress()

	testCases := []struct {
		name          string
		msg           MsgWithdrawLiquidity
		runeAmount    cosmos.Uint
		assetAmount   cosmos.Uint
		expectedError error
	}{
		{
			name: "all-good-asymmetric-rune",
			msg: MsgWithdrawLiquidity{
				WithdrawAddress: runeAddress,
				BasisPoints:     cosmos.NewUint(10000),
				Asset:           common.BNBAsset,
				Tx:              common.Tx{ID: "28B40BF105A112389A339A64BD1A042E6140DC9082C679586C6CF493A9FDE3FE"},
				WithdrawalAsset: common.RuneAsset(),
				Signer:          accountAddr,
			},
			runeAmount:    cosmos.NewUint(6250000000),
			assetAmount:   cosmos.ZeroUint(),
			expectedError: nil,
		},
		{
			name: "all-good-asymmetric-asset",
			msg: MsgWithdrawLiquidity{
				WithdrawAddress: runeAddress,
				BasisPoints:     cosmos.NewUint(10000),
				Asset:           common.BNBAsset,
				Tx:              common.Tx{ID: "28B40BF105A112389A339A64BD1A042E6140DC9082C679586C6CF493A9FDE3FE"},
				WithdrawalAsset: common.BNBAsset,
				Signer:          accountAddr,
			},
			runeAmount:    cosmos.ZeroUint(),
			assetAmount:   cosmos.NewUint(6250000000),
			expectedError: nil,
		},
	}
	for _, tc := range testCases {
		c.Logf("name:%s", tc.name)
		ctx, mgr := setupManagerForTest(c)
		ps := getWithdrawTestKeeper2V91(c, ctx, mgr.Keeper(), runeAddress)
		mgr.K = ps
		c.Assert(ps.SaveNetworkFee(ctx, common.BNBChain, NetworkFee{
			Chain:              common.BNBChain,
			TransactionSize:    1,
			TransactionFeeRate: bnbSingleTxFee.Uint64(),
		}), IsNil)
		r, asset, _, _, _, err := withdrawV91(ctx, tc.msg, mgr)
		if tc.expectedError != nil {
			c.Assert(err, NotNil)
			c.Check(err.Error(), Equals, tc.expectedError.Error())
			c.Check(r.Uint64(), Equals, tc.runeAmount.Uint64())
			c.Check(asset.Uint64(), Equals, tc.assetAmount.Uint64())
			continue
		}
		c.Assert(err, IsNil)
		c.Assert(r.Uint64(), Equals, tc.runeAmount.Uint64(), Commentf("%d != %d", r.Uint64(), tc.runeAmount.Uint64()))
		c.Assert(asset.Equal(tc.assetAmount), Equals, true, Commentf("expect:%s, however got:%s", tc.assetAmount.String(), asset.String()))
	}
}

func (WithdrawSuiteV91) TestWithdrawPendingRuneOrAsset(c *C) {
	accountAddr := GetRandomValidatorNode(NodeActive).NodeAddress
	ctx, mgr := setupManagerForTest(c)
	pool := Pool{
		BalanceRune:  cosmos.NewUint(100 * common.One),
		BalanceAsset: cosmos.NewUint(100 * common.One),
		Asset:        common.BNBAsset,
		LPUnits:      cosmos.NewUint(200 * common.One),
		Status:       PoolAvailable,
	}
	c.Assert(mgr.Keeper().SetPool(ctx, pool), IsNil)
	lp := LiquidityProvider{
		Asset:              common.BNBAsset,
		RuneAddress:        GetRandomRUNEAddress(),
		AssetAddress:       GetRandomBNBAddress(),
		LastAddHeight:      1024,
		LastWithdrawHeight: 0,
		Units:              cosmos.ZeroUint(),
		PendingRune:        cosmos.NewUint(1024),
		PendingAsset:       cosmos.ZeroUint(),
		PendingTxID:        GetRandomTxHash(),
	}
	mgr.Keeper().SetLiquidityProvider(ctx, lp)
	msg := MsgWithdrawLiquidity{
		WithdrawAddress: lp.RuneAddress,
		BasisPoints:     cosmos.NewUint(10000),
		Asset:           common.BNBAsset,
		Tx:              common.Tx{ID: "28B40BF105A112389A339A64BD1A042E6140DC9082C679586C6CF493A9FDE3FE"},
		WithdrawalAsset: common.BNBAsset,
		Signer:          accountAddr,
	}
	runeAmt, assetAmt, _, unitsLeft, gas, err := withdrawV91(ctx, msg, mgr)
	c.Assert(err, IsNil)
	c.Assert(runeAmt.Equal(cosmos.NewUint(1024)), Equals, true)
	c.Assert(assetAmt.IsZero(), Equals, true)
	c.Assert(unitsLeft.IsZero(), Equals, true)
	c.Assert(gas.IsZero(), Equals, true)

	lp1 := LiquidityProvider{
		Asset:              common.BNBAsset,
		RuneAddress:        GetRandomRUNEAddress(),
		AssetAddress:       GetRandomBNBAddress(),
		LastAddHeight:      1024,
		LastWithdrawHeight: 0,
		Units:              cosmos.ZeroUint(),
		PendingRune:        cosmos.ZeroUint(),
		PendingAsset:       cosmos.NewUint(1024),
		PendingTxID:        GetRandomTxHash(),
	}
	mgr.Keeper().SetLiquidityProvider(ctx, lp1)
	msg1 := MsgWithdrawLiquidity{
		WithdrawAddress: lp1.RuneAddress,
		BasisPoints:     cosmos.NewUint(10000),
		Asset:           common.BNBAsset,
		Tx:              common.Tx{ID: "28B40BF105A112389A339A64BD1A042E6140DC9082C679586C6CF493A9FDE3FE"},
		WithdrawalAsset: common.BNBAsset,
		Signer:          accountAddr,
	}
	runeAmt, assetAmt, _, unitsLeft, gas, err = withdrawV91(ctx, msg1, mgr)
	c.Assert(err, IsNil)
	c.Assert(assetAmt.Equal(cosmos.NewUint(1024)), Equals, true)
	c.Assert(runeAmt.IsZero(), Equals, true)
	c.Assert(unitsLeft.IsZero(), Equals, true)
	c.Assert(gas.IsZero(), Equals, true)
}

func (s *WithdrawSuiteV91) TestWithdrawPendingLiquidityShouldRoundToPoolDecimals(c *C) {
	accountAddr := GetRandomValidatorNode(NodeActive).NodeAddress
	ctx, mgr := setupManagerForTest(c)
	pool := Pool{
		BalanceRune:  cosmos.NewUint(100 * common.One),
		BalanceAsset: cosmos.NewUint(100 * common.One),
		Asset:        common.LUNAAsset,
		LPUnits:      cosmos.NewUint(200 * common.One),
		Status:       PoolAvailable,
		Decimals:     int64(6),
	}
	c.Assert(mgr.Keeper().SetPool(ctx, pool), IsNil)
	v := GetCurrentVersion()
	constantAccessor := constants.GetConstantValues(v)
	addHandler := NewAddLiquidityHandler(mgr)
	// create a LP record that has pending asset
	lpAddr := GetRandomTHORAddress()
	c.Assert(addHandler.addLiquidity(ctx,
		common.LUNAAsset,
		cosmos.ZeroUint(),
		cosmos.NewUint(339448125567),
		lpAddr,
		GetRandomBTCAddress(),
		GetRandomTxHash(),
		true,
		constantAccessor), IsNil)

	newctx := ctx.WithBlockHeight(ctx.BlockHeight() + 17280*2)
	msg2 := MsgWithdrawLiquidity{
		WithdrawAddress: lpAddr,
		BasisPoints:     cosmos.NewUint(10000),
		Asset:           common.LUNAAsset,
		Tx:              common.Tx{ID: "28B40BF105A112389A339A64BD1A042E6140DC9082C679586C6CF493A9FDE3FE"},
		WithdrawalAsset: common.LUNAAsset,
		Signer:          accountAddr,
	}
	runeAmt, assetAmt, protectoinRuneAmt, unitsClaimed, _, err := withdrawV91(newctx, msg2, mgr)
	c.Assert(err, IsNil)
	c.Assert(assetAmt.Equal(cosmos.NewUint(339448125500)), Equals, true, Commentf("%d", assetAmt.Uint64()))
	c.Assert(runeAmt.IsZero(), Equals, true)
	c.Assert(protectoinRuneAmt.IsZero(), Equals, true)
	c.Assert(unitsClaimed.IsZero(), Equals, true)
}

func getWithdrawTestKeeperV91(c *C, ctx cosmos.Context, k keeper.Keeper, runeAddress common.Address) keeper.Keeper {
	store := NewWithdrawTestKeeperV91(k)
	pool := Pool{
		BalanceRune:  cosmos.NewUint(100 * common.One),
		BalanceAsset: cosmos.NewUint(100 * common.One),
		Asset:        common.BNBAsset,
		LPUnits:      cosmos.NewUint(100 * common.One),
		SynthUnits:   cosmos.ZeroUint(),
		Status:       PoolAvailable,
	}
	c.Assert(store.SetPool(ctx, pool), IsNil)
	lp := LiquidityProvider{
		Asset:              pool.Asset,
		RuneAddress:        runeAddress,
		AssetAddress:       runeAddress,
		LastAddHeight:      0,
		LastWithdrawHeight: 0,
		Units:              cosmos.NewUint(100 * common.One),
		PendingRune:        cosmos.ZeroUint(),
		PendingAsset:       cosmos.ZeroUint(),
		PendingTxID:        "",
		RuneDepositValue:   cosmos.NewUint(100 * common.One),
		AssetDepositValue:  cosmos.NewUint(100 * common.One),
	}
	store.SetLiquidityProvider(ctx, lp)
	return store
}

func TestCalcImpLossV91(t *testing.T) {
	testCases := []struct {
		name                  string
		pool                  Pool
		lp                    types.LiquidityProvider
		withdrawBasisPoint    int64
		protectionBasisPoints int64
		expectedILP           cosmos.Uint
		expectedDepositValue  cosmos.Uint
		expectedRedeemValue   cosmos.Uint
	}{
		{
			name: "little-asset-large-rune",
			pool: Pool{
				Asset:        common.BTCAsset,
				BalanceAsset: cosmos.NewUint(100 * common.One),
				BalanceRune:  cosmos.NewUint(100 * common.One),
				LPUnits:      cosmos.NewUint(100 * common.One),
				SynthUnits:   cosmos.ZeroUint(),
			},
			lp: types.LiquidityProvider{
				Units:             cosmos.NewUint(12345678),
				AssetDepositValue: cosmos.NewUint(common.One),
				RuneDepositValue:  cosmos.NewUint(common.One * 50),
			},
			withdrawBasisPoint:    10000,
			protectionBasisPoints: 10000,
			expectedILP:           cosmos.NewUint(5075308644),
			expectedDepositValue:  cosmos.NewUint(5100000000),
			expectedRedeemValue:   cosmos.NewUint(24691356),
		},
		{
			name: "symmetrical-add",
			pool: Pool{
				Asset:        common.BTCAsset,
				BalanceAsset: cosmos.NewUint(100 * common.One),
				BalanceRune:  cosmos.NewUint(100 * common.One),
				LPUnits:      cosmos.NewUint(100 * common.One),
				SynthUnits:   cosmos.ZeroUint(),
			},
			lp: types.LiquidityProvider{
				Units:             cosmos.NewUint(common.One * 50),
				AssetDepositValue: cosmos.NewUint(common.One * 50),
				RuneDepositValue:  cosmos.NewUint(common.One * 50),
			},
			withdrawBasisPoint:    10000,
			protectionBasisPoints: 10000,
			expectedILP:           cosmos.ZeroUint(),
			expectedDepositValue:  cosmos.NewUint(10000000000),
			expectedRedeemValue:   cosmos.NewUint(10000000000),
		},
		{
			name: "price-less-than-one",
			pool: Pool{
				Asset:        common.BTCAsset,
				BalanceAsset: cosmos.NewUint(100 * common.One),
				BalanceRune:  cosmos.NewUint(10 * common.One),
				LPUnits:      cosmos.NewUint(100 * common.One),
				SynthUnits:   cosmos.ZeroUint(),
			},
			lp: types.LiquidityProvider{
				Units:             cosmos.NewUint(common.One * 10),
				AssetDepositValue: cosmos.NewUint(common.One * 50),
				RuneDepositValue:  cosmos.NewUint(common.One),
			},
			withdrawBasisPoint:    10000,
			protectionBasisPoints: 10000,
			expectedILP:           cosmos.NewUint(400000000),
			expectedDepositValue:  cosmos.NewUint(600000000),
			expectedRedeemValue:   cosmos.NewUint(200000000),
		},
		{
			name: "half-coverage",
			pool: Pool{
				Asset:        common.BTCAsset,
				BalanceAsset: cosmos.NewUint(100 * common.One),
				BalanceRune:  cosmos.NewUint(100 * common.One),
				LPUnits:      cosmos.NewUint(100 * common.One),
				SynthUnits:   cosmos.ZeroUint(),
			},
			lp: types.LiquidityProvider{
				Units:             cosmos.NewUint(12345678),
				AssetDepositValue: cosmos.NewUint(common.One),
				RuneDepositValue:  cosmos.NewUint(common.One * 50),
			},
			withdrawBasisPoint:    10000,
			protectionBasisPoints: 5000,
			expectedILP:           cosmos.NewUint(2537654322),
			expectedDepositValue:  cosmos.NewUint(5100000000),
			expectedRedeemValue:   cosmos.NewUint(24691356),
		},
		{
			name: "no-panic",
			pool: Pool{
				Asset:        common.BTCAsset,
				BalanceAsset: cosmos.NewUint(100 * common.One),
				BalanceRune:  cosmos.NewUint(100 * common.One),
				LPUnits:      getUintFromString("56641523781457101833101355"),
				SynthUnits:   getUintFromString("14364006570060942596417271288417281565500"),
			},
			lp: types.LiquidityProvider{
				Units:             getUintFromString("56641523781433101833101355"),
				AssetDepositValue: cosmos.NewUint(common.One),
				RuneDepositValue:  cosmos.NewUint(common.One * 50),
			},
			withdrawBasisPoint:    10000,
			protectionBasisPoints: 5000,
			expectedILP:           cosmos.ZeroUint(),
			expectedDepositValue:  cosmos.ZeroUint(),
			expectedRedeemValue:   cosmos.ZeroUint(),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t1 *testing.T) {
			ilpRune, depositValue, redeemValue := calcImpLossV91(tc.lp, cosmos.NewUint(uint64(tc.withdrawBasisPoint)), tc.protectionBasisPoints, tc.pool)
			assert.Equal(t1, ilpRune.String(), tc.expectedILP.String())
			assert.Equal(t1, depositValue.String(), tc.expectedDepositValue.String())
			assert.Equal(t1, redeemValue.String(), tc.expectedRedeemValue.String())
		})
	}
}

func getUintFromString(input string) cosmos.Uint {
	result, _ := cosmos.ParseUint(input)
	return result
}

func (WithdrawSuiteV91) TestWithdrawSynth(c *C) {
	accountAddr := GetRandomValidatorNode(NodeActive).NodeAddress
	ctx, mgr := setupManagerForTest(c)
	asset := common.BTCAsset.GetSyntheticAsset()

	coin := common.NewCoin(asset, cosmos.NewUint(100*common.One))
	c.Assert(mgr.Keeper().MintToModule(ctx, ModuleName, coin), IsNil)
	c.Assert(mgr.Keeper().SendFromModuleToModule(ctx, ModuleName, AsgardName, common.NewCoins(coin)), IsNil)

	pool := Pool{
		BalanceRune:  cosmos.ZeroUint(),
		BalanceAsset: coin.Amount,
		Asset:        asset,
		LPUnits:      cosmos.NewUint(200 * common.One),
		Status:       PoolAvailable,
	}
	c.Assert(mgr.Keeper().SetPool(ctx, pool), IsNil)
	lp := LiquidityProvider{
		Asset:              asset,
		RuneAddress:        common.NoAddress,
		AssetAddress:       GetRandomRUNEAddress(),
		LastAddHeight:      0,
		LastWithdrawHeight: 0,
		Units:              cosmos.NewUint(100 * common.One),
		PendingRune:        cosmos.ZeroUint(),
		PendingAsset:       cosmos.ZeroUint(),
		PendingTxID:        GetRandomTxHash(),
	}
	mgr.Keeper().SetLiquidityProvider(ctx, lp)
	msg := MsgWithdrawLiquidity{
		WithdrawAddress: lp.AssetAddress,
		BasisPoints:     cosmos.NewUint(MaxWithdrawBasisPoints / 2),
		Asset:           asset,
		Tx:              common.Tx{ID: "28B40BF105A112389A339A64BD1A042E6140DC9082C679586C6CF493A9FDE3FE"},
		WithdrawalAsset: common.EmptyAsset,
		Signer:          accountAddr,
	}
	runeAmt, assetAmt, _, unitsLeft, gas, err := withdrawV91(ctx, msg, mgr)
	c.Assert(err, IsNil)
	c.Check(assetAmt.Uint64(), Equals, uint64(25*common.One), Commentf("%d", assetAmt.Uint64()))
	c.Check(runeAmt.IsZero(), Equals, true)
	c.Check(unitsLeft.Uint64(), Equals, uint64(50*common.One), Commentf("%d", unitsLeft.Uint64()))
	c.Check(gas.IsZero(), Equals, true)

	pool, err = mgr.Keeper().GetPool(ctx, asset)
	c.Assert(err, IsNil)
	c.Check(pool.BalanceRune.Uint64(), Equals, uint64(0), Commentf("%d", pool.BalanceRune.Uint64()))
	c.Check(pool.BalanceAsset.Uint64(), Equals, uint64(75*common.One), Commentf("%d", pool.BalanceAsset.Uint64()))
	c.Check(pool.LPUnits.Uint64(), Equals, uint64(150*common.One), Commentf("%d", pool.LPUnits.Uint64())) // LP units did decreased
}

func (WithdrawSuiteV91) TestWithdrawSynthSingleLP(c *C) {
	accountAddr := GetRandomValidatorNode(NodeActive).NodeAddress
	ctx, mgr := setupManagerForTest(c)
	asset := common.BTCAsset.GetSyntheticAsset()

	coin := common.NewCoin(asset, cosmos.NewUint(30*common.One))
	c.Assert(mgr.Keeper().MintToModule(ctx, ModuleName, coin), IsNil)
	c.Assert(mgr.Keeper().SendFromModuleToModule(ctx, ModuleName, AsgardName, common.NewCoins(coin)), IsNil)

	pool := Pool{
		BalanceRune:  cosmos.ZeroUint(),
		BalanceAsset: coin.Amount,
		Asset:        asset,
		LPUnits:      cosmos.NewUint(200 * common.One),
		Status:       PoolAvailable,
	}
	c.Assert(mgr.Keeper().SetPool(ctx, pool), IsNil)
	lp := LiquidityProvider{
		Asset:              asset,
		RuneAddress:        common.NoAddress,
		AssetAddress:       GetRandomRUNEAddress(),
		LastAddHeight:      0,
		LastWithdrawHeight: 0,
		Units:              cosmos.NewUint(200 * common.One),
		PendingRune:        cosmos.ZeroUint(),
		PendingAsset:       cosmos.ZeroUint(),
		PendingTxID:        GetRandomTxHash(),
	}
	mgr.Keeper().SetLiquidityProvider(ctx, lp)
	msg := MsgWithdrawLiquidity{
		WithdrawAddress: lp.AssetAddress,
		BasisPoints:     cosmos.NewUint(MaxWithdrawBasisPoints),
		Asset:           asset,
		Tx:              common.Tx{ID: "28B40BF105A112389A339A64BD1A042E6140DC9082C679586C6CF493A9FDE3FE"},
		WithdrawalAsset: common.EmptyAsset,
		Signer:          accountAddr,
	}
	runeAmt, assetAmt, _, unitsLeft, gas, err := withdrawV91(ctx, msg, mgr)
	c.Assert(err, IsNil)
	c.Check(assetAmt.Uint64(), Equals, coin.Amount.Uint64(), Commentf("%d", assetAmt.Uint64()))
	c.Check(runeAmt.IsZero(), Equals, true)
	c.Check(unitsLeft.Uint64(), Equals, uint64(200*common.One), Commentf("%d", unitsLeft.Uint64()))
	c.Check(gas.IsZero(), Equals, true)

	pool, err = mgr.Keeper().GetPool(ctx, asset)
	c.Check(err, IsNil)
	c.Check(pool.BalanceRune.Uint64(), Equals, uint64(0), Commentf("%d", pool.BalanceRune.Uint64()))
	c.Check(pool.BalanceAsset.Uint64(), Equals, uint64(0), Commentf("%d", pool.BalanceAsset.Uint64()))
	c.Check(pool.LPUnits.Uint64(), Equals, uint64(0), Commentf("%d", pool.LPUnits.Uint64())) // LP units did decreased
}
