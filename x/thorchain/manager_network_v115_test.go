package thorchain

import (
	. "gopkg.in/check.v1"

	"gitlab.com/thorchain/thornode/common"
	"gitlab.com/thorchain/thornode/common/cosmos"
	"gitlab.com/thorchain/thornode/constants"
	"gitlab.com/thorchain/thornode/x/thorchain/keeper"
	"gitlab.com/thorchain/thornode/x/thorchain/types"
)

type NetworkManagerV115TestSuite struct{}

var _ = Suite(&NetworkManagerV115TestSuite{})

func (s *NetworkManagerV115TestSuite) SetUpSuite(c *C) {
	SetupConfigForTest()
}

func (s *NetworkManagerV115TestSuite) TestRagnarokChain(c *C) {
	ctx, _ := setupKeeperForTest(c)
	ctx = ctx.WithBlockHeight(100000)

	activeVault := GetRandomVault()
	activeVault.StatusSince = ctx.BlockHeight() - 10
	activeVault.Coins = common.Coins{
		common.NewCoin(common.BNBAsset, cosmos.NewUint(100*common.One)),
	}
	retireVault := GetRandomVault()
	retireVault.Chains = common.Chains{common.BNBChain, common.BTCChain}.Strings()
	yggVault := GetRandomVault()
	yggVault.Type = YggdrasilVault
	yggVault.Coins = common.Coins{
		common.NewCoin(common.BTCAsset, cosmos.NewUint(3*common.One)),
		common.NewCoin(common.RuneAsset(), cosmos.NewUint(300*common.One)),
	}

	btcPool := NewPool()
	btcPool.Asset = common.BTCAsset
	btcPool.BalanceRune = cosmos.NewUint(1000 * common.One)
	btcPool.BalanceAsset = cosmos.NewUint(10 * common.One)
	btcPool.LPUnits = cosmos.NewUint(1600)

	bnbPool := NewPool()
	bnbPool.Asset = common.BNBAsset
	bnbPool.BalanceRune = cosmos.NewUint(1000 * common.One)
	bnbPool.BalanceAsset = cosmos.NewUint(10 * common.One)
	bnbPool.LPUnits = cosmos.NewUint(1600)

	addr := GetRandomRUNEAddress()
	lps := LiquidityProviders{
		{
			RuneAddress:       addr,
			AssetAddress:      GetRandomBTCAddress(),
			LastAddHeight:     5,
			Units:             btcPool.LPUnits.QuoUint64(2),
			PendingRune:       cosmos.ZeroUint(),
			PendingAsset:      cosmos.ZeroUint(),
			AssetDepositValue: cosmos.ZeroUint(),
			RuneDepositValue:  cosmos.ZeroUint(),
		},
		{
			RuneAddress:       GetRandomRUNEAddress(),
			AssetAddress:      GetRandomBTCAddress(),
			LastAddHeight:     10,
			Units:             btcPool.LPUnits.QuoUint64(2),
			PendingRune:       cosmos.ZeroUint(),
			PendingAsset:      cosmos.ZeroUint(),
			AssetDepositValue: cosmos.ZeroUint(),
			RuneDepositValue:  cosmos.ZeroUint(),
		},
	}

	keeper := &TestRagnarokChainKeeper{
		na:          GetRandomValidatorNode(NodeActive),
		activeVault: activeVault,
		retireVault: retireVault,
		yggVault:    yggVault,
		pools:       Pools{bnbPool, btcPool},
		lps:         lps,
	}

	mgr := NewDummyMgrWithKeeper(keeper)

	networkMgr := newNetworkMgrV115(keeper, mgr.TxOutStore(), mgr.EventMgr())

	// the first round should just recall yggdrasil fund
	err := networkMgr.manageChains(ctx, mgr)
	c.Assert(err, IsNil)
	c.Check(keeper.pools[1].Asset.Equals(common.BTCAsset), Equals, true)
	c.Check(keeper.pools[1].LPUnits.IsZero(), Equals, false, Commentf("%d\n", keeper.pools[1].LPUnits.Uint64()))
	c.Check(keeper.pools[0].LPUnits.Equal(cosmos.NewUint(1600)), Equals, true)
	for _, skr := range keeper.lps {
		c.Check(skr.Units.IsZero(), Equals, false)
	}

	// the first round should just recall yggdrasil fund
	ctx = ctx.WithBlockHeight(200000)
	err = networkMgr.manageChains(ctx, mgr)
	c.Assert(err, IsNil)
	c.Check(keeper.pools[1].Asset.Equals(common.BTCAsset), Equals, true)
	c.Check(keeper.pools[1].LPUnits.IsZero(), Equals, true, Commentf("%d\n", keeper.pools[1].LPUnits.Uint64()))
	c.Check(keeper.pools[0].LPUnits.Equal(cosmos.NewUint(1600)), Equals, true)
	for _, skr := range keeper.lps {
		c.Check(skr.Units.IsZero(), Equals, true)
	}
	// ensure we have requested for ygg funds to be returned
	txOutStore := mgr.TxOutStore()
	c.Assert(err, IsNil)
	items, err := txOutStore.GetOutboundItems(ctx)
	c.Assert(err, IsNil)

	// 1 ygg return + 4 withdrawals
	c.Check(items, HasLen, 3, Commentf("Len %d", items))
	c.Check(items[0].Memo, Equals, NewYggdrasilReturn(100000).String())
	c.Check(items[0].Chain.Equals(common.BTCChain), Equals, true)

	ctx, mgr1 := setupManagerForTest(c)
	helper := NewVaultGenesisSetupTestHelper(mgr1.Keeper())
	mgr.K = helper
	networkMgr1 := newNetworkMgrV115(helper, mgr1.TxOutStore(), mgr1.EventMgr())
	// fail to get active nodes should error out
	helper.failToListActiveAccounts = true
	c.Assert(networkMgr1.ragnarokChain(ctx, common.BNBChain, 1, mgr), NotNil)
	helper.failToListActiveAccounts = false

	// no active nodes , should error
	c.Assert(networkMgr1.ragnarokChain(ctx, common.BNBChain, 1, mgr), NotNil)
	c.Assert(helper.Keeper.SetNodeAccount(ctx, GetRandomValidatorNode(NodeActive)), IsNil)
	c.Assert(helper.Keeper.SetNodeAccount(ctx, GetRandomValidatorNode(NodeActive)), IsNil)

	// fail to get pools should error out
	helper.failGetPools = true
	c.Assert(networkMgr1.ragnarokChain(ctx, common.BNBChain, 1, mgr), NotNil)
	helper.failGetPools = false
}

func (s *NetworkManagerV115TestSuite) TestUpdateNetwork(c *C) {
	ctx, mgr := setupManagerForTest(c)
	ver := GetCurrentVersion()
	constAccessor := constants.GetConstantValues(ver)
	helper := NewVaultGenesisSetupTestHelper(mgr.Keeper())
	mgr.K = helper
	networkMgr := newNetworkMgrV115(helper, mgr.TxOutStore(), mgr.EventMgr())

	// fail to get Network should return error
	helper.failGetNetwork = true
	c.Assert(networkMgr.UpdateNetwork(ctx, constAccessor, mgr.gasMgr, mgr.eventMgr), NotNil)
	helper.failGetNetwork = false

	// TotalReserve is zero , should not doing anything
	vd := NewNetwork()
	err := mgr.Keeper().SetNetwork(ctx, vd)
	c.Assert(err, IsNil)
	c.Assert(networkMgr.UpdateNetwork(ctx, constAccessor, mgr.GasMgr(), mgr.EventMgr()), IsNil)

	c.Assert(networkMgr.UpdateNetwork(ctx, constAccessor, mgr.GasMgr(), mgr.EventMgr()), IsNil)

	p := NewPool()
	p.Asset = common.BNBAsset
	p.BalanceRune = cosmos.NewUint(common.One * 100)
	p.BalanceAsset = cosmos.NewUint(common.One * 100)
	p.Status = PoolAvailable
	c.Assert(helper.SetPool(ctx, p), IsNil)
	// no active node , thus no bond
	c.Assert(networkMgr.UpdateNetwork(ctx, constAccessor, mgr.GasMgr(), mgr.EventMgr()), IsNil)

	// with liquidity fee , and bonds
	c.Assert(helper.Keeper.AddToLiquidityFees(ctx, common.BNBAsset, cosmos.NewUint(50*common.One)), IsNil)

	c.Assert(networkMgr.UpdateNetwork(ctx, constAccessor, mgr.GasMgr(), mgr.EventMgr()), IsNil)
	// add bond
	c.Assert(helper.Keeper.SetNodeAccount(ctx, GetRandomValidatorNode(NodeActive)), IsNil)
	c.Assert(helper.Keeper.SetNodeAccount(ctx, GetRandomValidatorNode(NodeActive)), IsNil)
	c.Assert(networkMgr.UpdateNetwork(ctx, constAccessor, mgr.GasMgr(), mgr.EventMgr()), IsNil)

	// fail to get total liquidity fee should result an error
	helper.failGetTotalLiquidityFee = true
	if common.RuneAsset().Equals(common.RuneNative) {
		FundModule(c, ctx, helper, ReserveName, 100)
	}
	c.Assert(networkMgr.UpdateNetwork(ctx, constAccessor, mgr.GasMgr(), mgr.EventMgr()), NotNil)
	helper.failGetTotalLiquidityFee = false

	helper.failToListActiveAccounts = true
	c.Assert(networkMgr.UpdateNetwork(ctx, constAccessor, mgr.GasMgr(), mgr.EventMgr()), NotNil)
}

func (s *NetworkManagerV115TestSuite) TestCalcBlockRewards(c *C) {
	mgr := NewDummyMgr()
	networkMgr := newNetworkMgrV115(keeper.KVStoreDummy{}, mgr.TxOutStore(), mgr.EventMgr())

	ver := GetCurrentVersion()
	constAccessor := constants.GetConstantValues(ver)
	emissionCurve := constAccessor.GetInt64Value(constants.EmissionCurve)
	incentiveCurve := constAccessor.GetInt64Value(constants.IncentiveCurve)
	blocksPerYear := constAccessor.GetInt64Value(constants.BlocksPerYear)

	bondR, poolR, lpD, lpShare := networkMgr.calcBlockRewards(cosmos.NewUint(1000*common.One), cosmos.NewUint(2000*common.One), cosmos.NewUint(1000*common.One), cosmos.ZeroUint(), emissionCurve, incentiveCurve, blocksPerYear)
	c.Check(bondR.Uint64(), Equals, uint64(1585), Commentf("%d", bondR.Uint64()))
	c.Check(poolR.Uint64(), Equals, uint64(1586), Commentf("%d", poolR.Uint64()))
	c.Check(lpD.Uint64(), Equals, uint64(0), Commentf("%d", lpD.Uint64()))
	c.Check(lpShare.Uint64(), Equals, uint64(5002), Commentf("%d", lpShare.Uint64()))

	bondR, poolR, lpD, lpShare = networkMgr.calcBlockRewards(cosmos.NewUint(1000*common.One), cosmos.NewUint(2000*common.One), cosmos.NewUint(1000*common.One), cosmos.NewUint(3000), emissionCurve, incentiveCurve, blocksPerYear)
	c.Check(bondR.Uint64(), Equals, uint64(3085), Commentf("%d", bondR.Uint64()))
	c.Check(poolR.Uint64(), Equals, uint64(86), Commentf("%d", poolR.Uint64()))
	c.Check(lpD.Uint64(), Equals, uint64(0), Commentf("%d", lpD.Uint64()))
	c.Check(lpShare.Uint64(), Equals, uint64(5001), Commentf("%d", lpShare.Uint64()))

	bondR, poolR, lpD, lpShare = networkMgr.calcBlockRewards(cosmos.NewUint(1000*common.One), cosmos.NewUint(2000*common.One), cosmos.ZeroUint(), cosmos.ZeroUint(), emissionCurve, incentiveCurve, blocksPerYear)
	c.Check(bondR.Uint64(), Equals, uint64(0), Commentf("%d", bondR.Uint64()))
	c.Check(poolR.Uint64(), Equals, uint64(0), Commentf("%d", poolR.Uint64()))
	c.Check(lpD.Uint64(), Equals, uint64(0), Commentf("%d", lpD.Uint64()))
	c.Check(lpShare.Uint64(), Equals, uint64(0), Commentf("%d", lpShare.Uint64()))

	bondR, poolR, lpD, lpShare = networkMgr.calcBlockRewards(cosmos.NewUint(1000*common.One), cosmos.NewUint(1000*common.One), cosmos.NewUint(1000*common.One), cosmos.ZeroUint(), emissionCurve, incentiveCurve, blocksPerYear)
	c.Check(bondR.Uint64(), Equals, uint64(3171), Commentf("%d", bondR.Uint64()))
	c.Check(poolR.Uint64(), Equals, uint64(0), Commentf("%d", poolR.Uint64()))
	c.Check(lpD.Uint64(), Equals, uint64(0), Commentf("%d", lpD.Uint64()))
	c.Check(lpShare.Uint64(), Equals, uint64(0), Commentf("%d", lpShare.Uint64()))

	bondR, poolR, lpD, lpShare = networkMgr.calcBlockRewards(cosmos.ZeroUint(), cosmos.NewUint(1000*common.One), cosmos.NewUint(1000*common.One), cosmos.ZeroUint(), emissionCurve, incentiveCurve, blocksPerYear)
	c.Check(bondR.Uint64(), Equals, uint64(0), Commentf("%d", bondR.Uint64()))
	c.Check(poolR.Uint64(), Equals, uint64(3171), Commentf("%d", poolR.Uint64()))
	c.Check(lpD.Uint64(), Equals, uint64(0), Commentf("%d", lpD.Uint64()))
	c.Check(lpShare.Uint64(), Equals, uint64(10_000), Commentf("%d", lpShare.Uint64()))

	bondR, poolR, lpD, lpShare = networkMgr.calcBlockRewards(cosmos.NewUint(2001*common.One), cosmos.NewUint(1000*common.One), cosmos.NewUint(1000*common.One), cosmos.ZeroUint(), emissionCurve, incentiveCurve, blocksPerYear)
	c.Check(bondR.Uint64(), Equals, uint64(3171), Commentf("%d", bondR.Uint64()))
	c.Check(poolR.Uint64(), Equals, uint64(0), Commentf("%d", poolR.Uint64()))
	c.Check(lpD.Uint64(), Equals, uint64(0), Commentf("%d", lpD.Uint64()))
	c.Check(lpShare.Uint64(), Equals, uint64(0), Commentf("%d", lpShare.Uint64()))
}

func (s *NetworkManagerV115TestSuite) TestCalcPoolDeficit(c *C) {
	pool1Fees := cosmos.NewUint(1000)
	pool2Fees := cosmos.NewUint(3000)
	totalFees := cosmos.NewUint(4000)

	mgr := NewDummyMgr()
	networkMgr := newNetworkMgrV115(keeper.KVStoreDummy{}, mgr.TxOutStore(), mgr.EventMgr())

	lpDeficit := cosmos.NewUint(1120)
	amt1 := networkMgr.calcPoolDeficit(lpDeficit, totalFees, pool1Fees)
	amt2 := networkMgr.calcPoolDeficit(lpDeficit, totalFees, pool2Fees)

	c.Check(amt1.Equal(cosmos.NewUint(280)), Equals, true, Commentf("%d", amt1.Uint64()))
	c.Check(amt2.Equal(cosmos.NewUint(840)), Equals, true, Commentf("%d", amt2.Uint64()))
}

func (*NetworkManagerV115TestSuite) TestProcessGenesisSetup(c *C) {
	ctx, mgr := setupManagerForTest(c)
	helper := NewVaultGenesisSetupTestHelper(mgr.Keeper())
	ctx = ctx.WithBlockHeight(1)
	mgr.K = helper
	networkMgr := newNetworkMgrV115(helper, mgr.TxOutStore(), mgr.EventMgr())
	// no active account
	c.Assert(networkMgr.EndBlock(ctx, mgr), NotNil)

	nodeAccount := GetRandomValidatorNode(NodeActive)
	c.Assert(mgr.Keeper().SetNodeAccount(ctx, nodeAccount), IsNil)
	c.Assert(networkMgr.EndBlock(ctx, mgr), IsNil)
	// make sure asgard vault get created
	vaults, err := mgr.Keeper().GetAsgardVaults(ctx)
	c.Assert(err, IsNil)
	c.Assert(vaults, HasLen, 1)

	// fail to get asgard vaults should return an error
	helper.failToGetAsgardVaults = true
	c.Assert(networkMgr.EndBlock(ctx, mgr), NotNil)
	helper.failToGetAsgardVaults = false

	// vault already exist , it should not do anything , and should not error
	c.Assert(networkMgr.EndBlock(ctx, mgr), IsNil)

	ctx, mgr = setupManagerForTest(c)
	helper = NewVaultGenesisSetupTestHelper(mgr.Keeper())
	ctx = ctx.WithBlockHeight(1)
	mgr.K = helper
	networkMgr = newNetworkMgrV115(helper, mgr.TxOutStore(), mgr.EventMgr())
	helper.failToListActiveAccounts = true
	c.Assert(networkMgr.EndBlock(ctx, mgr), NotNil)
	helper.failToListActiveAccounts = false

	helper.failToSetVault = true
	c.Assert(networkMgr.EndBlock(ctx, mgr), NotNil)
	helper.failToSetVault = false

	helper.failGetRetiringAsgardVault = true
	ctx = ctx.WithBlockHeight(1024)
	c.Assert(networkMgr.EndBlock(ctx, mgr), NotNil)
	helper.failGetRetiringAsgardVault = false

	helper.failGetActiveAsgardVault = true
	c.Assert(networkMgr.EndBlock(ctx, mgr), NotNil)
	helper.failGetActiveAsgardVault = false
}

func (*NetworkManagerV115TestSuite) TestGetTotalActiveBond(c *C) {
	ctx, mgr := setupManagerForTest(c)
	helper := NewVaultGenesisSetupTestHelper(mgr.Keeper())
	mgr.K = helper
	networkMgr := newNetworkMgrV115(helper, mgr.TxOutStore(), mgr.EventMgr())
	helper.failToListActiveAccounts = true
	bond, err := networkMgr.getTotalActiveBond(ctx)
	c.Assert(err, NotNil)
	c.Assert(bond.Equal(cosmos.ZeroUint()), Equals, true)
	helper.failToListActiveAccounts = false
	c.Assert(helper.Keeper.SetNodeAccount(ctx, GetRandomValidatorNode(NodeActive)), IsNil)
	bond, err = networkMgr.getTotalActiveBond(ctx)
	c.Assert(err, IsNil)
	c.Assert(bond.Uint64() > 0, Equals, true)
}

func (*NetworkManagerV115TestSuite) TestGetTotalLiquidityRune(c *C) {
	ctx, mgr := setupManagerForTest(c)
	helper := NewVaultGenesisSetupTestHelper(mgr.Keeper())
	mgr.K = helper
	networkMgr := newNetworkMgrV115(helper, mgr.TxOutStore(), mgr.EventMgr())
	p := NewPool()
	p.Asset = common.BNBAsset
	p.BalanceRune = cosmos.NewUint(common.One * 100)
	p.BalanceAsset = cosmos.NewUint(common.One * 100)
	p.Status = PoolAvailable
	c.Assert(helper.SetPool(ctx, p), IsNil)
	pools, totalLiquidity, err := networkMgr.getTotalProvidedLiquidityRune(ctx)
	c.Assert(err, IsNil)
	c.Assert(pools, HasLen, 1)
	c.Assert(totalLiquidity.Equal(p.BalanceRune), Equals, true)
}

func (*NetworkManagerV115TestSuite) TestPayPoolRewards(c *C) {
	ctx, mgr := setupManagerForTest(c)
	helper := NewVaultGenesisSetupTestHelper(mgr.Keeper())
	mgr.K = helper
	networkMgr := newNetworkMgrV115(helper, mgr.TxOutStore(), mgr.EventMgr())
	p := NewPool()
	p.Asset = common.BNBAsset
	p.BalanceRune = cosmos.NewUint(common.One * 100)
	p.BalanceAsset = cosmos.NewUint(common.One * 100)
	p.Status = PoolAvailable
	c.Assert(helper.SetPool(ctx, p), IsNil)
	c.Assert(networkMgr.payPoolRewards(ctx, []cosmos.Uint{cosmos.NewUint(100 * common.One)}, Pools{p}), IsNil)
	helper.failToSetPool = true
	c.Assert(networkMgr.payPoolRewards(ctx, []cosmos.Uint{cosmos.NewUint(100 * common.One)}, Pools{p}), NotNil)
}

func (*NetworkManagerV115TestSuite) TestFindChainsToRetire(c *C) {
	ctx, mgr := setupManagerForTest(c)
	helper := NewVaultGenesisSetupTestHelper(mgr.Keeper())
	mgr.K = helper
	networkMgr := newNetworkMgrV115(helper, mgr.TxOutStore(), mgr.EventMgr())
	// fail to get active asgard vault
	helper.failGetActiveAsgardVault = true
	chains, err := networkMgr.findChainsToRetire(ctx)
	c.Assert(err, NotNil)
	c.Assert(chains, HasLen, 0)
	helper.failGetActiveAsgardVault = false

	// fail to get retire asgard vault
	helper.failGetRetiringAsgardVault = true
	chains, err = networkMgr.findChainsToRetire(ctx)
	c.Assert(err, NotNil)
	c.Assert(chains, HasLen, 0)
	helper.failGetRetiringAsgardVault = false
}

func (*NetworkManagerV115TestSuite) TestRecallChainFunds(c *C) {
	ctx, mgr := setupManagerForTest(c)
	helper := NewVaultGenesisSetupTestHelper(mgr.Keeper())
	mgr.K = helper
	networkMgr := newNetworkMgrV115(helper, mgr.TxOutStore(), mgr.EventMgr())
	helper.failToListActiveAccounts = true
	c.Assert(networkMgr.RecallChainFunds(ctx, common.BNBChain, mgr, common.PubKeys{}), NotNil)
	helper.failToListActiveAccounts = false

	helper.failGetActiveAsgardVault = true
	c.Assert(networkMgr.RecallChainFunds(ctx, common.BNBChain, mgr, common.PubKeys{}), NotNil)
	helper.failGetActiveAsgardVault = false
}

func (s *NetworkManagerV115TestSuite) TestRecoverPoolDeficit(c *C) {
	ctx, mgr := setupManagerForTest(c)
	helper := NewVaultGenesisSetupTestHelper(mgr.Keeper())
	mgr.K = helper
	networkMgr := newNetworkMgrV115(helper, mgr.TxOutStore(), mgr.EventMgr())

	pools := Pools{
		Pool{
			Asset:        common.BNBAsset,
			BalanceRune:  cosmos.NewUint(common.One * 2000),
			BalanceAsset: cosmos.NewUint(common.One * 2000),
			Status:       PoolAvailable,
		},
	}
	c.Assert(helper.Keeper.SetPool(ctx, pools[0]), IsNil)

	totalLiquidityFees := cosmos.NewUint(50 * common.One)
	c.Assert(helper.Keeper.AddToLiquidityFees(ctx, common.BNBAsset, totalLiquidityFees), IsNil)

	lpDeficit := cosmos.NewUint(totalLiquidityFees.Uint64())

	bondBefore := helper.Keeper.GetRuneBalanceOfModule(ctx, BondName)
	asgardBefore := helper.Keeper.GetRuneBalanceOfModule(ctx, AsgardName)
	reserveBefore := helper.Keeper.GetRuneBalanceOfModule(ctx, ReserveName)

	poolAmts, err := networkMgr.deductPoolRewardDeficit(ctx, pools, totalLiquidityFees, lpDeficit)
	c.Assert(err, IsNil)
	c.Assert(len(poolAmts), Equals, 1)

	bondAfter := helper.Keeper.GetRuneBalanceOfModule(ctx, BondName)
	asgardAfter := helper.Keeper.GetRuneBalanceOfModule(ctx, AsgardName)
	reserveAfter := helper.Keeper.GetRuneBalanceOfModule(ctx, ReserveName)

	// bond module is not touched
	c.Assert(bondAfter.String(), Equals, bondBefore.String())

	// deficit moves from asgard to reserve
	c.Assert(asgardAfter.String(), Equals, asgardBefore.Sub(lpDeficit).String())
	c.Assert(reserveAfter.String(), Equals, reserveBefore.Add(lpDeficit).String())

	// deficit rune is deducted from the pool record
	pool, err := helper.Keeper.GetPool(ctx, common.BNBAsset)
	c.Assert(err, IsNil)
	c.Assert(pool.BalanceRune.String(), Equals, pools[0].BalanceRune.Sub(lpDeficit).String())
}

func (s *NetworkManagerV115TestSuite) TestSaverYieldFunc(c *C) {
	var err error
	ctx, mgr := setupManagerForTest(c)
	net := newNetworkMgrV115(mgr.Keeper(), mgr.TxOutStore(), mgr.EventMgr())
	mgr.Keeper().SetMimir(ctx, constants.SynthYieldCycle.String(), 5_000)

	// mint synths
	coin := common.NewCoin(common.BTCAsset.GetSyntheticAsset(), cosmos.NewUint(10*common.One))
	c.Assert(mgr.Keeper().MintToModule(ctx, ModuleName, coin), IsNil)
	c.Assert(mgr.Keeper().SendFromModuleToModule(ctx, ModuleName, AsgardName, common.NewCoins(coin)), IsNil)

	spool := NewPool()
	spool.Asset = common.BTCAsset.GetSyntheticAsset()
	spool.BalanceAsset = coin.Amount
	spool.LPUnits = cosmos.NewUint(100)
	c.Assert(mgr.Keeper().SetPool(ctx, spool), IsNil)

	// first pool
	pool := NewPool()
	pool.Asset = common.BTCAsset
	pool.BalanceRune = cosmos.NewUint(100 * common.One)
	pool.BalanceAsset = cosmos.NewUint(100 * common.One)
	pool.LPUnits = cosmos.NewUint(100)
	pool.CalcUnits(mgr.GetVersion(), coin.Amount)
	c.Assert(mgr.Keeper().SetPool(ctx, pool), IsNil)

	c.Assert(net.paySaverYield(ctx, common.BTCAsset, cosmos.NewUint(50*common.One)), IsNil)
	spool, err = mgr.Keeper().GetPool(ctx, spool.Asset)
	c.Assert(err, IsNil)
	c.Assert(spool.BalanceAsset.String(), Equals, "1113100000", Commentf("%d", spool.BalanceAsset.Uint64()))
}

func (s *NetworkManagerV115TestSuite) TestSaverYieldCall(c *C) {
	var err error
	ctx, mgr := setupManagerForTest(c)
	ver := GetCurrentVersion()
	constAccessor := constants.GetConstantValues(ver)

	na := GetRandomValidatorNode(NodeActive)
	na.Bond = cosmos.NewUint(500000 * common.One)
	c.Assert(mgr.Keeper().SetNodeAccount(ctx, na), IsNil)

	coin := common.NewCoin(common.BTCAsset.GetSyntheticAsset(), cosmos.NewUint(10*common.One))
	spool := NewPool()
	spool.Asset = common.BTCAsset.GetSyntheticAsset()
	spool.BalanceAsset = coin.Amount
	spool.LPUnits = cosmos.NewUint(100)
	c.Assert(mgr.Keeper().SetPool(ctx, spool), IsNil)

	// first pool
	pool := NewPool()
	pool.Asset = common.BTCAsset
	pool.BalanceRune = cosmos.NewUint(100 * common.One)
	pool.BalanceAsset = cosmos.NewUint(100 * common.One)
	pool.LPUnits = cosmos.NewUint(100)
	pool.CalcUnits(mgr.GetVersion(), coin.Amount)
	c.Assert(mgr.Keeper().SetPool(ctx, pool), IsNil)

	networkMgr := newNetworkMgrV115(mgr.Keeper(), mgr.TxOutStore(), mgr.EventMgr())

	// test no fees collected
	c.Assert(networkMgr.UpdateNetwork(ctx, constAccessor, mgr.gasMgr, mgr.eventMgr), IsNil)
	spool, err = mgr.Keeper().GetPool(ctx, spool.Asset.GetSyntheticAsset())
	c.Assert(err, IsNil)
	c.Check(spool.BalanceAsset.Uint64(), Equals, uint64(7155446454), Commentf("%d", spool.BalanceAsset.Uint64()))

	// mgr.Keeper().SetMimir(ctx, constants.IncentiveCurve.String(), 50)
	c.Assert(mgr.Keeper().AddToLiquidityFees(ctx, pool.Asset, cosmos.NewUint(50*common.One)), IsNil)
	c.Assert(networkMgr.UpdateNetwork(ctx, constAccessor, mgr.gasMgr, mgr.eventMgr), IsNil)
	spool, err = mgr.Keeper().GetPool(ctx, spool.Asset.GetSyntheticAsset())
	c.Assert(err, IsNil)
	c.Check(spool.BalanceAsset.String(), Equals, "7834021738", Commentf("%d", spool.BalanceAsset.Uint64()))

	// check we don't give yield when synth utilization is too high
	// add some synths
	coins := cosmos.NewCoins(cosmos.NewCoin("btc/btc", cosmos.NewInt(101*common.One))) // 51% utilization
	c.Assert(mgr.coinKeeper.MintCoins(ctx, ModuleName, coins), IsNil)
	c.Assert(mgr.Keeper().AddToLiquidityFees(ctx, pool.Asset, cosmos.NewUint(50*common.One)), IsNil)
	c.Assert(networkMgr.UpdateNetwork(ctx, constAccessor, mgr.gasMgr, mgr.eventMgr), IsNil)
	spool, err = mgr.Keeper().GetPool(ctx, spool.Asset.GetSyntheticAsset())
	c.Assert(err, IsNil)
	c.Check(spool.BalanceAsset.String(), Equals, "7834021738", Commentf("%d", spool.BalanceAsset.Uint64()))
}

func (s *NetworkManagerV115TestSuite) TestRagnarokPool(c *C) {
	ctx, k := setupKeeperForTest(c)
	ctx = ctx.WithBlockHeight(100000)
	na := GetRandomValidatorNode(NodeActive)
	c.Assert(k.SetNodeAccount(ctx, na), IsNil)
	activeVault := GetRandomVault()
	activeVault.StatusSince = ctx.BlockHeight() - 10
	activeVault.Coins = common.Coins{
		common.NewCoin(common.BNBAsset, cosmos.NewUint(100*common.One)),
	}
	c.Assert(k.SetVault(ctx, activeVault), IsNil)
	retireVault := GetRandomVault()
	retireVault.Chains = common.Chains{common.BNBChain, common.BTCChain}.Strings()
	yggVault := GetRandomVault()
	yggVault.PubKey = na.PubKeySet.Secp256k1
	yggVault.Type = YggdrasilVault
	yggVault.Coins = common.Coins{
		common.NewCoin(common.BTCAsset, cosmos.NewUint(3*common.One)),
	}
	c.Assert(k.SetVault(ctx, yggVault), IsNil)
	btcPool := NewPool()
	btcPool.Asset = common.BTCAsset
	btcPool.BalanceRune = cosmos.NewUint(1000 * common.One)
	btcPool.BalanceAsset = cosmos.NewUint(10 * common.One)
	btcPool.LPUnits = cosmos.NewUint(1600)
	btcPool.Status = PoolAvailable
	c.Assert(k.SetPool(ctx, btcPool), IsNil)
	bnbPool := NewPool()
	bnbPool.Asset = common.BNBAsset
	bnbPool.BalanceRune = cosmos.NewUint(1000 * common.One)
	bnbPool.BalanceAsset = cosmos.NewUint(10 * common.One)
	bnbPool.LPUnits = cosmos.NewUint(1600)
	bnbPool.Status = PoolAvailable
	c.Assert(k.SetPool(ctx, bnbPool), IsNil)
	addr := GetRandomRUNEAddress()
	lps := LiquidityProviders{
		{
			Asset:             common.BTCAsset,
			RuneAddress:       addr,
			AssetAddress:      GetRandomBTCAddress(),
			LastAddHeight:     5,
			Units:             btcPool.LPUnits.QuoUint64(2),
			PendingRune:       cosmos.ZeroUint(),
			PendingAsset:      cosmos.ZeroUint(),
			AssetDepositValue: cosmos.ZeroUint(),
			RuneDepositValue:  cosmos.ZeroUint(),
		},
		{
			Asset:             common.BTCAsset,
			RuneAddress:       GetRandomRUNEAddress(),
			AssetAddress:      GetRandomBTCAddress(),
			LastAddHeight:     10,
			Units:             btcPool.LPUnits.QuoUint64(2),
			PendingRune:       cosmos.ZeroUint(),
			PendingAsset:      cosmos.ZeroUint(),
			AssetDepositValue: cosmos.ZeroUint(),
			RuneDepositValue:  cosmos.ZeroUint(),
		},
	}
	k.SetLiquidityProvider(ctx, lps[0])
	k.SetLiquidityProvider(ctx, lps[1])
	mgr := NewDummyMgrWithKeeper(k)
	networkMgr := newNetworkMgrV115(k, mgr.TxOutStore(), mgr.EventMgr())

	ctx = ctx.WithBlockHeight(1)
	// block height not correct , doesn't take any actions
	err := networkMgr.checkPoolRagnarok(ctx, mgr)
	c.Assert(err, IsNil)
	for _, a := range []common.Asset{common.BTCAsset, common.BNBAsset} {
		tempPool, err := k.GetPool(ctx, a)
		c.Assert(err, IsNil)
		c.Assert(tempPool.Status, Equals, PoolAvailable)
	}
	interval := mgr.GetConstants().GetInt64Value(constants.FundMigrationInterval)
	// mimir didn't set , it should not take any actions
	ctx = ctx.WithBlockHeight(interval * 5)
	err = networkMgr.checkPoolRagnarok(ctx, mgr)
	c.Assert(err, IsNil)

	// happy path
	networkMgr.k.SetMimir(ctx, "RAGNAROK-BTC-BTC", 1)
	// first round , it should recall yggdrasil
	err = networkMgr.checkPoolRagnarok(ctx, mgr)
	c.Assert(err, IsNil)
	items, _ := mgr.txOutStore.GetOutboundItems(ctx)
	c.Assert(items, HasLen, 1)
	c.Assert(items[0].Memo, Equals, "YGGDRASIL-:200")

	// second round, ragnarok
	ctx = ctx.WithBlockHeight(interval * 6)
	err = networkMgr.checkPoolRagnarok(ctx, mgr)
	c.Assert(err, IsNil)
	items, _ = mgr.txOutStore.GetOutboundItems(ctx)
	c.Assert(items, HasLen, 3)

	tempPool, err := k.GetPool(ctx, common.BTCAsset)
	c.Assert(err, IsNil)
	c.Assert(tempPool.Status, Equals, PoolSuspended)

	tempPool, err = k.GetPool(ctx, common.BNBAsset)
	c.Assert(err, IsNil)
	c.Assert(tempPool.Status, Equals, PoolAvailable)

	// when there are none gas token pool , and it is active , gas asset token pool should not be ragnarok
	busdPool := NewPool()
	busdAsset, err := common.NewAsset("BNB.BUSD-BD1")
	c.Assert(err, IsNil)
	busdPool.Asset = busdAsset
	busdPool.BalanceRune = cosmos.NewUint(1000 * common.One)
	busdPool.BalanceAsset = cosmos.NewUint(10 * common.One)
	busdPool.LPUnits = cosmos.NewUint(1600)
	busdPool.Status = PoolAvailable
	c.Assert(k.SetPool(ctx, busdPool), IsNil)

	networkMgr.k.SetMimir(ctx, "RAGNAROK-BNB-BNB", 1)
	err = networkMgr.checkPoolRagnarok(ctx, mgr)
	c.Assert(err, IsNil)
	tempPool, err = k.GetPool(ctx, common.BNBAsset)
	c.Assert(err, IsNil)
	c.Assert(tempPool.Status, Equals, PoolAvailable)
}

func (s *NetworkManagerV115TestSuite) TestCleanupAsgardIndex(c *C) {
	ctx, k := setupKeeperForTest(c)
	vault1 := NewVault(1024, ActiveVault, AsgardVault, GetRandomPubKey(), common.Chains{common.BNBChain}.Strings(), []ChainContract{})
	c.Assert(k.SetVault(ctx, vault1), IsNil)
	vault2 := NewVault(1024, RetiringVault, AsgardVault, GetRandomPubKey(), common.Chains{common.BNBChain}.Strings(), []ChainContract{})
	c.Assert(k.SetVault(ctx, vault2), IsNil)
	vault3 := NewVault(1024, InitVault, AsgardVault, GetRandomPubKey(), common.Chains{common.BNBChain}.Strings(), []ChainContract{})
	c.Assert(k.SetVault(ctx, vault3), IsNil)
	vault4 := NewVault(1024, InactiveVault, AsgardVault, GetRandomPubKey(), common.Chains{common.BNBChain}.Strings(), []ChainContract{})
	c.Assert(k.SetVault(ctx, vault4), IsNil)
	mgr := NewDummyMgrWithKeeper(k)
	networkMgr := newNetworkMgrV115(k, mgr.TxOutStore(), mgr.EventMgr())
	c.Assert(networkMgr.cleanupAsgardIndex(ctx), IsNil)
	containsVault := func(vaults Vaults, pubKey common.PubKey) bool {
		for _, item := range vaults {
			if item.PubKey.Equals(pubKey) {
				return true
			}
		}
		return false
	}
	asgards, err := k.GetAsgardVaults(ctx)
	c.Assert(err, IsNil)
	c.Assert(containsVault(asgards, vault1.PubKey), Equals, true)
	c.Assert(containsVault(asgards, vault2.PubKey), Equals, true)
	c.Assert(containsVault(asgards, vault3.PubKey), Equals, true)
	c.Assert(containsVault(asgards, vault4.PubKey), Equals, false)
}

func (*NetworkManagerV115TestSuite) TestPOLLiquidityAdd(c *C) {
	ctx, mgr := setupManagerForTest(c)

	net := newNetworkMgrV115(mgr.Keeper(), NewTxStoreDummy(), NewDummyEventMgr())
	max := cosmos.NewUint(10000)

	polAddress, err := mgr.Keeper().GetModuleAddress(ReserveName)
	c.Assert(err, IsNil)
	asgardAddress, err := mgr.Keeper().GetModuleAddress(AsgardName)
	c.Assert(err, IsNil)
	na := GetRandomValidatorNode(NodeActive)
	signer := na.NodeAddress
	c.Assert(mgr.Keeper().SetNodeAccount(ctx, na), IsNil)

	btcPool := NewPool()
	btcPool.Asset = common.BTCAsset
	btcPool.BalanceRune = cosmos.NewUint(2000 * common.One)
	btcPool.BalanceAsset = cosmos.NewUint(20 * common.One)
	btcPool.LPUnits = cosmos.NewUint(1600)
	c.Assert(mgr.Keeper().SetPool(ctx, btcPool), IsNil)

	// hit max
	util := cosmos.NewUint(1500)
	target := cosmos.NewUint(1000)
	c.Assert(net.addPOLLiquidity(ctx, btcPool, polAddress, asgardAddress, signer, max, util, target, mgr), IsNil)
	lp, err := mgr.Keeper().GetLiquidityProvider(ctx, btcPool.Asset, polAddress)
	c.Assert(err, IsNil)
	c.Check(lp.Units.Uint64(), Equals, uint64(7), Commentf("%d", lp.Units.Uint64()))

	// doesn't hit max
	util = cosmos.NewUint(1050)
	c.Assert(net.addPOLLiquidity(ctx, btcPool, polAddress, asgardAddress, signer, max, util, target, mgr), IsNil)
	lp, err = mgr.Keeper().GetLiquidityProvider(ctx, btcPool.Asset, polAddress)
	c.Assert(err, IsNil)
	c.Check(lp.Units.Uint64(), Equals, uint64(10), Commentf("%d", lp.Units.Uint64()))

	// no change needed
	util = cosmos.NewUint(1000)
	c.Assert(net.addPOLLiquidity(ctx, btcPool, polAddress, asgardAddress, signer, max, util, target, mgr), IsNil)
	lp, err = mgr.Keeper().GetLiquidityProvider(ctx, btcPool.Asset, polAddress)
	c.Assert(err, IsNil)
	c.Check(lp.Units.Uint64(), Equals, uint64(10), Commentf("%d", lp.Units.Uint64()))

	// not enough balance in the reserve module
	max = cosmos.NewUint(1000000)
	util = cosmos.NewUint(50_000)
	btcPool.BalanceRune = cosmos.NewUint(90000000000 * common.One)
	c.Assert(net.addPOLLiquidity(ctx, btcPool, polAddress, asgardAddress, signer, max, util, target, mgr), IsNil)
	lp, err = mgr.Keeper().GetLiquidityProvider(ctx, btcPool.Asset, polAddress)
	c.Assert(err, IsNil)
	c.Check(lp.Units.Uint64(), Equals, uint64(10), Commentf("%d", lp.Units.Uint64()))
}

func (*NetworkManagerV115TestSuite) TestPOLLiquidityWithdraw(c *C) {
	ctx, mgr := setupManagerForTest(c)

	net := newNetworkMgrV115(mgr.Keeper(), NewTxStoreDummy(), NewDummyEventMgr())
	max := cosmos.NewUint(10000)

	polAddress, err := mgr.Keeper().GetModuleAddress(ReserveName)
	c.Assert(err, IsNil)
	asgardAddress, err := mgr.Keeper().GetModuleAddress(AsgardName)
	c.Assert(err, IsNil)
	na := GetRandomValidatorNode(NodeActive)
	signer := na.NodeAddress
	c.Assert(mgr.Keeper().SetNodeAccount(ctx, na), IsNil)

	vault := GetRandomVault()
	c.Assert(mgr.Keeper().SetVault(ctx, vault), IsNil)

	btcPool := NewPool()
	btcPool.Asset = common.BTCAsset
	btcPool.BalanceRune = cosmos.NewUint(2000 * common.One)
	btcPool.BalanceAsset = cosmos.NewUint(20 * common.One)
	btcPool.LPUnits = cosmos.NewUint(1600)
	c.Assert(mgr.Keeper().SetPool(ctx, btcPool), IsNil)

	lps := LiquidityProviders{
		{
			Asset:             btcPool.Asset,
			RuneAddress:       GetRandomBNBAddress(),
			AssetAddress:      GetRandomBTCAddress(),
			LastAddHeight:     5,
			Units:             btcPool.LPUnits.QuoUint64(2),
			PendingRune:       cosmos.ZeroUint(),
			PendingAsset:      cosmos.ZeroUint(),
			AssetDepositValue: cosmos.ZeroUint(),
			RuneDepositValue:  cosmos.ZeroUint(),
		},
		{
			Asset:             btcPool.Asset,
			RuneAddress:       polAddress,
			AssetAddress:      common.NoAddress,
			LastAddHeight:     10,
			Units:             btcPool.LPUnits.QuoUint64(2),
			PendingRune:       cosmos.ZeroUint(),
			PendingAsset:      cosmos.ZeroUint(),
			AssetDepositValue: cosmos.ZeroUint(),
			RuneDepositValue:  cosmos.ZeroUint(),
		},
	}
	for _, lp := range lps {
		mgr.Keeper().SetLiquidityProvider(ctx, lp)
	}

	// hit max
	util := cosmos.NewUint(500)
	target := cosmos.NewUint(1000)
	c.Assert(net.removePOLLiquidity(ctx, btcPool, polAddress, asgardAddress, signer, max, util, target, mgr), IsNil)
	lp, err := mgr.Keeper().GetLiquidityProvider(ctx, btcPool.Asset, polAddress)
	c.Assert(err, IsNil)
	c.Check(lp.Units.Uint64(), Equals, uint64(792), Commentf("%d", lp.Units.Uint64()))
	// To withdraw max 1% (100 basis points) of the pool RUNE depth, asymmetrically withdraw as RUNE 0.5% of all pool units.
	// 0.5% of 1600 is 8; 800 minus 8 is 792.

	// doesn't hit max
	util = cosmos.NewUint(950)
	c.Assert(net.removePOLLiquidity(ctx, btcPool, polAddress, asgardAddress, signer, max, util, target, mgr), IsNil)
	lp, err = mgr.Keeper().GetLiquidityProvider(ctx, btcPool.Asset, polAddress)
	c.Assert(err, IsNil)
	c.Check(lp.Units.Uint64(), Equals, uint64(788), Commentf("%d", lp.Units.Uint64()))
	// To withdraw 0.5% of the pool RUNE depth, asymmetrically withdraw as RUNE 0.25% of all pool units.
	// 0.25% of 1592 is 3.98 which rounds to 4; 792 minus 4 is 788.

	// no change needed
	util = cosmos.NewUint(1000)
	c.Assert(net.removePOLLiquidity(ctx, btcPool, polAddress, asgardAddress, signer, max, util, target, mgr), IsNil)
	lp, err = mgr.Keeper().GetLiquidityProvider(ctx, btcPool.Asset, polAddress)
	c.Assert(err, IsNil)
	c.Check(lp.Units.Uint64(), Equals, uint64(788), Commentf("%d", lp.Units.Uint64()))
}

func (*NetworkManagerV115TestSuite) TestFairMergePOLCycle(c *C) {
	ctx, mgr := setupManagerForTest(c)
	net := newNetworkMgrV115(mgr.Keeper(), NewTxStoreDummy(), NewDummyEventMgr())

	// cycle should do nothing when target is 0
	err := net.POLCycle(ctx, mgr)
	c.Assert(err, IsNil)
	pol, err := mgr.Keeper().GetPOL(ctx)
	c.Assert(err, IsNil)
	c.Assert(pol.RuneDeposited.Uint64(), Equals, uint64(0))
	c.Assert(pol.RuneWithdrawn.Uint64(), Equals, uint64(0))

	// cycle should error when target is greater than 0 with no node accounts
	mgr.Keeper().SetMimir(ctx, constants.POLTargetSynthPerPoolDepth.String(), 1000) // 10% liability
	err = net.POLCycle(ctx, mgr)
	c.Assert(err, ErrorMatches, "dev err: no active node accounts")

	// create dummy bnb pool
	pool := NewPool()
	pool.Asset = common.BNBAsset
	pool.BalanceRune = cosmos.NewUint(100 * common.One)
	pool.BalanceAsset = cosmos.NewUint(100 * common.One)
	pool.Status = PoolAvailable
	pool.LPUnits = cosmos.NewUint(100 * common.One)
	err = mgr.Keeper().SetPool(ctx, pool)
	c.Assert(err, IsNil)

	// cycle should error since there are no pol enabled pools
	err = mgr.Keeper().SetNodeAccount(ctx, GetRandomValidatorNode(NodeActive))
	c.Assert(err, IsNil)
	err = net.POLCycle(ctx, mgr)
	c.Assert(err, ErrorMatches, "no POL pools")

	// cycle should silently succeed when there is a pool enabled
	mgr.Keeper().SetMimir(ctx, "POL-BNB-BNB", 1)
	err = net.POLCycle(ctx, mgr)
	c.Assert(err, IsNil)

	// pol should still be zero since there are no synths
	pol, err = mgr.Keeper().GetPOL(ctx)
	c.Assert(err, IsNil)
	c.Assert(pol.RuneDeposited.Uint64(), Equals, uint64(0))
	c.Assert(pol.RuneWithdrawn.Uint64(), Equals, uint64(0))

	// add some synths
	coins := cosmos.NewCoins(cosmos.NewCoin("bnb/bnb", cosmos.NewInt(20*common.One))) // 20% utilization, 10% liability
	err = mgr.coinKeeper.MintCoins(ctx, ModuleName, coins)
	c.Assert(err, IsNil)
	err = mgr.Keeper().SetPool(ctx, pool)
	c.Assert(err, IsNil)

	// synth liability should be 10%
	synthSupply := mgr.Keeper().GetTotalSupply(ctx, pool.Asset.GetSyntheticAsset())
	pool.CalcUnits(mgr.GetVersion(), synthSupply)
	liability := common.GetUncappedShare(pool.SynthUnits, pool.GetPoolUnits(), cosmos.NewUint(10_000))
	c.Assert(liability.String(), Equals, "1000")

	// cycle should succeed, still no rune deposited since max is 0
	err = net.POLCycle(ctx, mgr)
	c.Assert(err, IsNil)

	// pol should still be zero
	pol, err = mgr.Keeper().GetPOL(ctx)
	c.Assert(err, IsNil)
	c.Assert(pol.RuneDeposited.String(), Equals, "0")
	c.Assert(pol.RuneWithdrawn.String(), Equals, "0")

	// synth liability should still be 10%
	synthSupply = mgr.Keeper().GetTotalSupply(ctx, pool.Asset.GetSyntheticAsset())
	pool.CalcUnits(mgr.GetVersion(), synthSupply)
	liability = common.GetUncappedShare(pool.SynthUnits, pool.GetPoolUnits(), cosmos.NewUint(10_000))
	c.Assert(liability.String(), Equals, "1000")

	// set pol utilization to 5% should deposit up to the max
	mgr.Keeper().SetMimir(ctx, constants.POLMaxNetworkDeposit.String(), common.One)
	mgr.Keeper().SetMimir(ctx, constants.POLTargetSynthPerPoolDepth.String(), 500)
	mgr.Keeper().SetMimir(ctx, constants.POLMaxPoolMovement.String(), 10000)
	err = net.POLCycle(ctx, mgr)
	c.Assert(err, IsNil)
	pol, err = mgr.Keeper().GetPOL(ctx)
	c.Assert(err, IsNil)
	c.Assert(pol.RuneDeposited.String(), Equals, "100000000")
	c.Assert(pol.RuneWithdrawn.String(), Equals, "0")

	// there needs to be one vault or the withdraw handler fails
	vault := NewVault(0, ActiveVault, types.VaultType_AsgardVault, GetRandomPubKey(), []string{"BNB"}, nil)
	err = mgr.Keeper().SetVault(ctx, vault)
	c.Assert(err, IsNil)

	// synth liability should still be 10%
	synthSupply = mgr.Keeper().GetTotalSupply(ctx, pool.Asset.GetSyntheticAsset())
	pool.CalcUnits(mgr.GetVersion(), synthSupply)
	liability = common.GetUncappedShare(pool.SynthUnits, pool.GetPoolUnits(), cosmos.NewUint(10_000))
	c.Assert(liability.String(), Equals, "1000")

	// withdraw entire pol position
	mgr.Keeper().SetMimir(ctx, constants.POLTargetSynthPerPoolDepth.String(), 10000)
	err = net.POLCycle(ctx, mgr)
	c.Assert(err, IsNil)
	pol, err = mgr.Keeper().GetPOL(ctx)
	c.Assert(err, IsNil)
	c.Assert(pol.RuneDeposited.String(), Equals, "100000000")
	c.Assert(pol.RuneWithdrawn.String(), Equals, "99451741") // minus slip

	// synth liability should still be 10%
	synthSupply = mgr.Keeper().GetTotalSupply(ctx, pool.Asset.GetSyntheticAsset())
	pool.CalcUnits(mgr.GetVersion(), synthSupply)
	liability = common.GetUncappedShare(pool.SynthUnits, pool.GetPoolUnits(), cosmos.NewUint(10_000))
	c.Assert(liability.String(), Equals, "1000")

	// deposit entire pol position
	mgr.Keeper().SetMimir(ctx, constants.POLTargetSynthPerPoolDepth.String(), 500)
	err = net.POLCycle(ctx, mgr)
	c.Assert(err, IsNil)
	pol, err = mgr.Keeper().GetPOL(ctx)
	c.Assert(err, IsNil)
	c.Assert(pol.RuneDeposited.String(), Equals, "200005483")
	c.Assert(pol.RuneWithdrawn.String(), Equals, "99451741")

	// withdraw entire pol position 1 basis point of rune depth at a time
	mgr.Keeper().SetMimir(ctx, constants.POLTargetSynthPerPoolDepth.String(), 10000)
	mgr.Keeper().SetMimir(ctx, constants.POLMaxPoolMovement.String(), 1)
	err = net.POLCycle(ctx, mgr)
	c.Assert(err, IsNil)
	pol, err = mgr.Keeper().GetPOL(ctx)
	c.Assert(err, IsNil)
	c.Assert(pol.RuneDeposited.String(), Equals, "200005483")
	c.Assert(pol.RuneWithdrawn.String(), Equals, "99461736")
	// another basis point
	err = net.POLCycle(ctx, mgr)
	c.Assert(err, IsNil)
	pol, err = mgr.Keeper().GetPOL(ctx)
	c.Assert(err, IsNil)
	c.Assert(pol.RuneDeposited.String(), Equals, "200005483")
	c.Assert(pol.RuneWithdrawn.String(), Equals, "99471729")

	// set the buffer to 100% to stop any movement
	mgr.Keeper().SetMimir(ctx, constants.POLBuffer.String(), 10000)
	err = net.POLCycle(ctx, mgr)
	c.Assert(err, IsNil)
	pol, err = mgr.Keeper().GetPOL(ctx)
	c.Assert(err, IsNil)
	c.Assert(pol.RuneDeposited.String(), Equals, "200005483")
	c.Assert(pol.RuneWithdrawn.String(), Equals, "99471729")

	// current liability is at 10%, so buffer at 40% and target of 50% should still not move
	mgr.Keeper().SetMimir(ctx, constants.POLBuffer.String(), 4000)
	mgr.Keeper().SetMimir(ctx, constants.POLTargetSynthPerPoolDepth.String(), 5000)
	err = net.POLCycle(ctx, mgr)
	c.Assert(err, IsNil)
	pol, err = mgr.Keeper().GetPOL(ctx)
	c.Assert(err, IsNil)
	c.Assert(pol.RuneDeposited.String(), Equals, "200005483")
	c.Assert(pol.RuneWithdrawn.String(), Equals, "99471729")

	// any smaller buffer should withdraw one basis point of rune
	mgr.Keeper().SetMimir(ctx, constants.POLBuffer.String(), 3999)
	err = net.POLCycle(ctx, mgr)
	c.Assert(err, IsNil)
	pol, err = mgr.Keeper().GetPOL(ctx)
	c.Assert(err, IsNil)
	c.Assert(pol.RuneDeposited.String(), Equals, "200005483")
	c.Assert(pol.RuneWithdrawn.String(), Equals, "99481722")

	// withdraw everything
	mgr.Keeper().SetMimir(ctx, constants.POLTargetSynthPerPoolDepth.String(), 10000)
	mgr.Keeper().SetMimir(ctx, constants.POLBuffer.String(), 0)
	mgr.Keeper().SetMimir(ctx, constants.POLMaxPoolMovement.String(), 10000)
	err = net.POLCycle(ctx, mgr)
	c.Assert(err, IsNil)
	pol, err = mgr.Keeper().GetPOL(ctx)
	c.Assert(err, IsNil)
	c.Assert(pol.RuneDeposited.String(), Equals, "200005483")
	c.Assert(pol.RuneWithdrawn.String(), Equals, "198909097")

	// should be nothing left to withdraw again
	err = net.POLCycle(ctx, mgr)
	c.Assert(err, IsNil)
	pol, err = mgr.Keeper().GetPOL(ctx)
	c.Assert(err, IsNil)
	c.Assert(pol.RuneDeposited.String(), Equals, "200005483")
	c.Assert(pol.RuneWithdrawn.String(), Equals, "198909097")
}

func (s *NetworkManagerV115TestSuite) TestSpawnDerivedAssets(c *C) {
	ctx, mgr := setupManagerForTest(c)

	nmgr := newNetworkMgrV115(mgr.Keeper(), NewTxStoreDummy(), NewDummyEventMgr())

	vault := GetRandomVault()
	c.Assert(mgr.Keeper().SetVault(ctx, vault), IsNil)

	mgr.Keeper().SetMimir(ctx, "DerivedDepthBasisPts", 10_000)
	mgr.Keeper().SetMimir(ctx, "TorAnchor-BNB-BUSD-BD1", 1) // enable BUSD pool as a TOR anchor
	maxAnchorSlip := fetchConfigInt64(ctx, mgr, constants.MaxAnchorSlip)
	busd, err := common.NewAsset("BNB.BUSD-BD1")
	c.Assert(err, IsNil)

	pool := NewPool()
	pool.Asset = busd
	pool.Status = PoolAvailable
	pool.BalanceRune = cosmos.NewUint(187493559385369)
	pool.BalanceAsset = cosmos.NewUint(925681680182301)
	pool.Decimals = 8
	c.Assert(mgr.Keeper().SetPool(ctx, pool), IsNil)

	bnb, err := common.NewAsset("BNB.BNB")
	c.Assert(err, IsNil)

	pool = NewPool()
	pool.Asset = bnb
	pool.Status = PoolAvailable
	pool.BalanceRune = cosmos.NewUint(110119961610327)
	pool.BalanceAsset = cosmos.NewUint(2343330836117)
	pool.Decimals = 8
	c.Assert(mgr.Keeper().SetPool(ctx, pool), IsNil)

	// happy path
	err = nmgr.spawnDerivedAssets(ctx, mgr)
	c.Assert(err, IsNil)
	usd, err := mgr.Keeper().GetPool(ctx, common.TOR)
	c.Assert(err, IsNil)
	c.Check(usd.BalanceAsset.Uint64(), Equals, uint64(925681680182301), Commentf("%d", usd.BalanceAsset.Uint64()))
	c.Check(usd.BalanceRune.Uint64(), Equals, uint64(187493559385369), Commentf("%d", usd.BalanceRune.Uint64()))
	dbnb, _ := common.NewAsset("THOR.BNB")
	bnbPool, err := mgr.Keeper().GetPool(ctx, dbnb)
	c.Assert(err, IsNil)
	c.Check(bnbPool.BalanceAsset.Uint64(), Equals, uint64(2343330836117), Commentf("%d", bnbPool.BalanceAsset.Uint64()))
	c.Check(bnbPool.BalanceRune.Uint64(), Equals, uint64(110119961610327), Commentf("%d", bnbPool.BalanceRune.Uint64()))

	// happy path, but some trade volume triggers a lower pool depth
	newctx := ctx.WithBlockHeight(ctx.BlockHeight() - 1)
	err = mgr.Keeper().AddToSwapSlip(newctx, busd, cosmos.NewInt(maxAnchorSlip/4))
	c.Assert(err, IsNil)
	err = nmgr.spawnDerivedAssets(ctx, mgr)
	c.Assert(err, IsNil)
	usd, err = mgr.Keeper().GetPool(ctx, common.TOR)
	c.Assert(err, IsNil)
	c.Check(usd.Status.String(), Equals, "Available")
	c.Check(usd.BalanceAsset.Uint64(), Equals, uint64(694261260136726), Commentf("%d", usd.BalanceAsset.Uint64()))
	c.Check(usd.BalanceRune.Uint64(), Equals, uint64(140620169539027), Commentf("%d", usd.BalanceRune.Uint64()))

	// unhappy path, too much liquidity fees collected in the anchor pools, goes to 1% depth
	err = mgr.Keeper().AddToSwapSlip(newctx, busd, cosmos.NewInt(10_000))
	c.Assert(err, IsNil)
	err = nmgr.spawnDerivedAssets(ctx, mgr)
	c.Assert(err, IsNil)
	usd, err = mgr.Keeper().GetPool(ctx, common.TOR)
	c.Assert(err, IsNil)
	c.Assert(usd.Status.String(), Equals, "Available")
	c.Assert(usd.BalanceAsset.Uint64(), Equals, uint64(9256816801824), Commentf("%d", usd.BalanceAsset.Uint64()))
	c.Assert(usd.BalanceRune.Uint64(), Equals, uint64(1874935593854), Commentf("%d", usd.BalanceRune.Uint64()))
	// ensure layer1 bnb pool is NOT suspended
	bnbPool, err = mgr.Keeper().GetPool(ctx, busd)
	c.Assert(err, IsNil)
	c.Assert(bnbPool.Status.String(), Equals, "Available")
	c.Assert(bnbPool.BalanceAsset.Uint64(), Equals, uint64(925681680182301), Commentf("%d", bnbPool.BalanceAsset.Uint64()))
	c.Assert(bnbPool.BalanceRune.Uint64(), Equals, uint64(187493559385369), Commentf("%d", bnbPool.BalanceRune.Uint64()))
}

func (s *NetworkManagerV115TestSuite) TestSpawnDerivedAssetsBasisPoints(c *C) {
	ctx, mgr := setupManagerForTest(c)

	nmgr := newNetworkMgrV115(mgr.Keeper(), NewTxStoreDummy(), NewDummyEventMgr())

	vault := GetRandomVault()
	c.Assert(mgr.Keeper().SetVault(ctx, vault), IsNil)

	mgr.Keeper().SetMimir(ctx, "TorAnchor-BNB-BUSD-BD1", 1) // enable BUSD pool as a TOR anchor
	busd, err := common.NewAsset("BNB.BUSD-BD1")
	c.Assert(err, IsNil)

	pool := NewPool()
	pool.Asset = busd
	pool.Status = PoolAvailable
	pool.BalanceRune = cosmos.NewUint(187493559385369)
	pool.BalanceAsset = cosmos.NewUint(925681680182301)
	pool.Decimals = 8
	c.Assert(mgr.Keeper().SetPool(ctx, pool), IsNil)

	// test that DerivedDepthBasisPts affects the pool depth
	mgr.Keeper().SetMimir(ctx, "DerivedDepthBasisPts", 20000)
	err = nmgr.spawnDerivedAssets(ctx, mgr)
	c.Assert(err, IsNil)
	usd, err := mgr.Keeper().GetPool(ctx, common.TOR)
	c.Assert(err, IsNil)
	c.Assert(usd.Status.String(), Equals, "Available")
	c.Check(usd.BalanceAsset.Uint64(), Equals, uint64(1851363360364602), Commentf("%d", usd.BalanceAsset.Uint64()))
	c.Check(usd.BalanceRune.Uint64(), Equals, uint64(374987118770738), Commentf("%d", usd.BalanceRune.Uint64()))

	// test that DerivedDepthBasisPts set to zero will cause the pools to
	// become suspended
	mgr.Keeper().SetMimir(ctx, "DerivedDepthBasisPts", 0)
	err = nmgr.spawnDerivedAssets(ctx, mgr)
	c.Assert(err, IsNil)
	usd, err = mgr.Keeper().GetPool(ctx, common.TOR)
	c.Assert(err, IsNil)
	c.Assert(usd.Status.String(), Equals, "Suspended")
	c.Assert(usd.BalanceAsset.Uint64(), Equals, uint64(1851363360364602), Commentf("%d", usd.BalanceAsset.Uint64()))
	c.Assert(usd.BalanceRune.Uint64(), Equals, uint64(374987118770738), Commentf("%d", usd.BalanceRune.Uint64()))
}
