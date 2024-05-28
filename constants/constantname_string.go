// Code generated by "stringer -type=ConstantName"; DO NOT EDIT.

package constants

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[EmissionCurve-0]
	_ = x[IncentiveCurve-1]
	_ = x[MaxRuneSupply-2]
	_ = x[BlocksPerYear-3]
	_ = x[OutboundTransactionFee-4]
	_ = x[NativeTransactionFee-5]
	_ = x[KillSwitchStart-6]
	_ = x[KillSwitchDuration-7]
	_ = x[PoolCycle-8]
	_ = x[MinRunePoolDepth-9]
	_ = x[MaxAvailablePools-10]
	_ = x[StagedPoolCost-11]
	_ = x[PendingLiquidityAgeLimit-12]
	_ = x[MinimumNodesForYggdrasil-13]
	_ = x[MinimumNodesForBFT-14]
	_ = x[DesiredValidatorSet-15]
	_ = x[AsgardSize-16]
	_ = x[DerivedDepthBasisPts-17]
	_ = x[DerivedMinDepth-18]
	_ = x[MaxAnchorSlip-19]
	_ = x[MaxAnchorBlocks-20]
	_ = x[DynamicMaxAnchorSlipBlocks-21]
	_ = x[DynamicMaxAnchorTarget-22]
	_ = x[DynamicMaxAnchorCalcInterval-23]
	_ = x[ChurnInterval-24]
	_ = x[ChurnRetryInterval-25]
	_ = x[ValidatorsChangeWindow-26]
	_ = x[LeaveProcessPerBlockHeight-27]
	_ = x[BadValidatorRedline-28]
	_ = x[LackOfObservationPenalty-29]
	_ = x[SigningTransactionPeriod-30]
	_ = x[DoubleSignMaxAge-31]
	_ = x[PauseBond-32]
	_ = x[PauseUnbond-33]
	_ = x[MinimumBondInRune-34]
	_ = x[FundMigrationInterval-35]
	_ = x[ArtificialRagnarokBlockHeight-36]
	_ = x[MaximumLiquidityRune-37]
	_ = x[StrictBondLiquidityRatio-38]
	_ = x[DefaultPoolStatus-39]
	_ = x[MaxOutboundAttempts-40]
	_ = x[SlashPenalty-41]
	_ = x[PauseOnSlashThreshold-42]
	_ = x[FailKeygenSlashPoints-43]
	_ = x[FailKeysignSlashPoints-44]
	_ = x[LiquidityLockUpBlocks-45]
	_ = x[ObserveSlashPoints-46]
	_ = x[MissBlockSignSlashPoints-47]
	_ = x[ObservationDelayFlexibility-48]
	_ = x[StopFundYggdrasil-49]
	_ = x[YggFundLimit-50]
	_ = x[YggFundRetry-51]
	_ = x[JailTimeKeygen-52]
	_ = x[JailTimeKeysign-53]
	_ = x[NodePauseChainBlocks-54]
	_ = x[EnableDerivedAssets-55]
	_ = x[MinSwapsPerBlock-56]
	_ = x[MaxSwapsPerBlock-57]
	_ = x[EnableOrderBooks-58]
	_ = x[MintSynths-59]
	_ = x[BurnSynths-60]
	_ = x[MaxSynthPerAssetDepth-61]
	_ = x[MaxSynthPerPoolDepth-62]
	_ = x[MaxSynthsForSaversYield-63]
	_ = x[VirtualMultSynths-64]
	_ = x[VirtualMultSynthsBasisPoints-65]
	_ = x[MinSlashPointsForBadValidator-66]
	_ = x[FullImpLossProtectionBlocks-67]
	_ = x[BondLockupPeriod-68]
	_ = x[MaxBondProviders-69]
	_ = x[NumberOfNewNodesPerChurn-70]
	_ = x[MinTxOutVolumeThreshold-71]
	_ = x[TxOutDelayRate-72]
	_ = x[TxOutDelayMax-73]
	_ = x[MaxTxOutOffset-74]
	_ = x[TNSRegisterFee-75]
	_ = x[TNSFeeOnSale-76]
	_ = x[TNSFeePerBlock-77]
	_ = x[StreamingSwapPause-78]
	_ = x[StreamingSwapMinBPFee-79]
	_ = x[StreamingSwapMaxLength-80]
	_ = x[StreamingSwapMaxLengthNative-81]
	_ = x[MinCR-82]
	_ = x[MaxCR-83]
	_ = x[LoanStreamingSwapsInterval-84]
	_ = x[PauseLoans-85]
	_ = x[LoanRepaymentMaturity-86]
	_ = x[LendingLever-87]
	_ = x[PermittedSolvencyGap-88]
	_ = x[NodeOperatorFee-89]
	_ = x[ValidatorMaxRewardRatio-90]
	_ = x[PoolDepthForYggFundingMin-91]
	_ = x[MaxNodeToChurnOutForLowVersion-92]
	_ = x[ChurnOutForLowVersionBlocks-93]
	_ = x[POLMaxNetworkDeposit-94]
	_ = x[POLMaxPoolMovement-95]
	_ = x[POLSynthUtilization-96]
	_ = x[POLTargetSynthPerPoolDepth-97]
	_ = x[POLBuffer-98]
	_ = x[RagnarokProcessNumOfLPPerIteration-99]
	_ = x[SwapOutDexAggregationDisabled-100]
	_ = x[SynthYieldBasisPoints-101]
	_ = x[SynthYieldCycle-102]
	_ = x[MinimumL1OutboundFeeUSD-103]
	_ = x[MinimumPoolLiquidityFee-104]
	_ = x[ILPCutoff-105]
	_ = x[ChurnMigrateRounds-106]
	_ = x[AllowWideBlame-107]
	_ = x[MaxAffiliateFeeBasisPoints-108]
	_ = x[TargetOutboundFeeSurplusRune-109]
	_ = x[MaxOutboundFeeMultiplierBasisPoints-110]
	_ = x[MinOutboundFeeMultiplierBasisPoints-111]
	_ = x[NativeOutboundFeeUSD-112]
	_ = x[NativeTransactionFeeUSD-113]
	_ = x[TNSRegisterFeeUSD-114]
	_ = x[TNSFeePerBlockUSD-115]
	_ = x[EnableUSDFees-116]
	_ = x[PreferredAssetOutboundFeeMultiplier-117]
	_ = x[FeeUSDRoundSignificantDigits-118]
	_ = x[MigrationVaultSecurityBps-119]
	_ = x[CloutReset-120]
	_ = x[CloutLimit-121]
	_ = x[KeygenRetryInterval-122]
	_ = x[SaversStreamingSwapsInterval-123]
	_ = x[RescheduleCoalesceBlocks-124]
	_ = x[SignerConcurrency-125]
	_ = x[L1SlipMinBps-126]
	_ = x[SynthSlipMinBps-127]
	_ = x[TradeAccountsSlipMinBps-128]
	_ = x[TradeAccountsEnabled-129]
	_ = x[EVMDisableContractWhitelist-130]
	_ = x[OperationalVotesMin-131]
}

const _ConstantName_name = "EmissionCurveIncentiveCurveMaxRuneSupplyBlocksPerYearOutboundTransactionFeeNativeTransactionFeeKillSwitchStartKillSwitchDurationPoolCycleMinRunePoolDepthMaxAvailablePoolsStagedPoolCostPendingLiquidityAgeLimitMinimumNodesForYggdrasilMinimumNodesForBFTDesiredValidatorSetAsgardSizeDerivedDepthBasisPtsDerivedMinDepthMaxAnchorSlipMaxAnchorBlocksDynamicMaxAnchorSlipBlocksDynamicMaxAnchorTargetDynamicMaxAnchorCalcIntervalChurnIntervalChurnRetryIntervalValidatorsChangeWindowLeaveProcessPerBlockHeightBadValidatorRedlineLackOfObservationPenaltySigningTransactionPeriodDoubleSignMaxAgePauseBondPauseUnbondMinimumBondInRuneFundMigrationIntervalArtificialRagnarokBlockHeightMaximumLiquidityRuneStrictBondLiquidityRatioDefaultPoolStatusMaxOutboundAttemptsSlashPenaltyPauseOnSlashThresholdFailKeygenSlashPointsFailKeysignSlashPointsLiquidityLockUpBlocksObserveSlashPointsMissBlockSignSlashPointsObservationDelayFlexibilityStopFundYggdrasilYggFundLimitYggFundRetryJailTimeKeygenJailTimeKeysignNodePauseChainBlocksEnableDerivedAssetsMinSwapsPerBlockMaxSwapsPerBlockEnableOrderBooksMintSynthsBurnSynthsMaxSynthPerAssetDepthMaxSynthPerPoolDepthMaxSynthsForSaversYieldVirtualMultSynthsVirtualMultSynthsBasisPointsMinSlashPointsForBadValidatorFullImpLossProtectionBlocksBondLockupPeriodMaxBondProvidersNumberOfNewNodesPerChurnMinTxOutVolumeThresholdTxOutDelayRateTxOutDelayMaxMaxTxOutOffsetTNSRegisterFeeTNSFeeOnSaleTNSFeePerBlockStreamingSwapPauseStreamingSwapMinBPFeeStreamingSwapMaxLengthStreamingSwapMaxLengthNativeMinCRMaxCRLoanStreamingSwapsIntervalPauseLoansLoanRepaymentMaturityLendingLeverPermittedSolvencyGapNodeOperatorFeeValidatorMaxRewardRatioPoolDepthForYggFundingMinMaxNodeToChurnOutForLowVersionChurnOutForLowVersionBlocksPOLMaxNetworkDepositPOLMaxPoolMovementPOLSynthUtilizationPOLTargetSynthPerPoolDepthPOLBufferRagnarokProcessNumOfLPPerIterationSwapOutDexAggregationDisabledSynthYieldBasisPointsSynthYieldCycleMinimumL1OutboundFeeUSDMinimumPoolLiquidityFeeILPCutoffChurnMigrateRoundsAllowWideBlameMaxAffiliateFeeBasisPointsTargetOutboundFeeSurplusRuneMaxOutboundFeeMultiplierBasisPointsMinOutboundFeeMultiplierBasisPointsNativeOutboundFeeUSDNativeTransactionFeeUSDTNSRegisterFeeUSDTNSFeePerBlockUSDEnableUSDFeesPreferredAssetOutboundFeeMultiplierFeeUSDRoundSignificantDigitsMigrationVaultSecurityBpsCloutResetCloutLimitKeygenRetryIntervalSaversStreamingSwapsIntervalRescheduleCoalesceBlocksSignerConcurrencyL1SlipMinBpsSynthSlipMinBpsTradeAccountsSlipMinBpsTradeAccountsEnabledEVMDisableContractWhitelistOperationalVotesMin"

var _ConstantName_index = [...]uint16{0, 13, 27, 40, 53, 75, 95, 110, 128, 137, 153, 170, 184, 208, 232, 250, 269, 279, 299, 314, 327, 342, 368, 390, 418, 431, 449, 471, 497, 516, 540, 564, 580, 589, 600, 617, 638, 667, 687, 711, 728, 747, 759, 780, 801, 823, 844, 862, 886, 913, 930, 942, 954, 968, 983, 1003, 1022, 1038, 1054, 1070, 1080, 1090, 1111, 1131, 1154, 1171, 1199, 1228, 1255, 1271, 1287, 1311, 1334, 1348, 1361, 1375, 1389, 1401, 1415, 1433, 1454, 1476, 1504, 1509, 1514, 1540, 1550, 1571, 1583, 1603, 1618, 1641, 1666, 1696, 1723, 1743, 1761, 1780, 1806, 1815, 1849, 1878, 1899, 1914, 1937, 1960, 1969, 1987, 2001, 2027, 2055, 2090, 2125, 2145, 2168, 2185, 2202, 2215, 2250, 2278, 2303, 2313, 2323, 2342, 2370, 2394, 2411, 2423, 2438, 2461, 2481, 2508, 2527}

func (i ConstantName) String() string {
	if i < 0 || i >= ConstantName(len(_ConstantName_index)-1) {
		return "ConstantName(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ConstantName_name[_ConstantName_index[i]:_ConstantName_index[i+1]]
}
