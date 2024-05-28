//go:build !stagenet && !mocknet
// +build !stagenet,!mocknet

package aggregators

import (
	"gitlab.com/thorchain/thornode/common"
)

func DexAggregatorsV120() []Aggregator {
	return []Aggregator{
		// TSAggregatorPancakeSwap Ethereum V2
		{common.ETHChain, `0x35CF22003c90126528fbe95b21bB3ADB2ca8c53D`, 400_000},
		// TSAggregatorWoofi Avalanche V2
		{common.AVAXChain, `0x5505BE604dFA8A1ad402A71f8A357fba47F9bf5a`, 400_000},
		// TSAggregatorGeneric
		{common.ETHChain, `0xd31f7e39afECEc4855fecc51b693F9A0Cec49fd2`, 400_000},
		// RangoThorchainOutputAggUniV2
		{common.ETHChain, `0x2a7813412b8da8d18Ce56FE763B9eb264D8e28a8`, 400_000},
		// RangoThorchainOutputAggUniV3
		{common.ETHChain, `0xbB8De86F3b041B3C084431dcf3159fE4827c5F0D`, 400_000},
		// PangolinAggregator
		{common.AVAXChain, `0x7a68c37D8AFA3078f3Ad51D98eA23Fe57a8Ae21a`, 400_000},
		// TSAggregatorUniswapV2 - short notation
		{common.ETHChain, `0x86904eb2b3c743400d03f929f2246efa80b91215`, 400_000},
		// TSAggregatorSushiswap - short notation
		{common.ETHChain, `0xbf365e79aa44a2164da135100c57fdb6635ae870`, 400_000},
		// TSAggregatorUniswapV3 100 - short notation
		{common.ETHChain, `0xbd68cbe6c247e2c3a0e36b8f0e24964914f26ee8`, 400_000},
		// TSAggregatorUniswapV3 500 - short notation
		{common.ETHChain, `0xe4ddca21881bac219af7f217703db0475d2a9f02`, 400_000},
		// TSAggregatorUniswapV3 3000 - short notation
		{common.ETHChain, `0x11733abf0cdb43298f7e949c930188451a9a9ef2`, 400_000},
		// TSAggregatorUniswapV3 10000 - short notation
		{common.ETHChain, `0xb33874810e5395eb49d8bd7e912631db115d5a03`, 400_000},
		// TSAggregatorPangolin
		{common.AVAXChain, `0x942c6dA485FD6cEf255853ef83a149d43A73F18a`, 400_000},
		// TSAggregatorTraderJoe
		{common.AVAXChain, `0x3b7DbdD635B99cEa39D3d95Dbd0217F05e55B212`, 400_000},
		// TSAggregatorAvaxGeneric
		{common.AVAXChain, `0x7C38b8B2efF28511ECc14a621e263857Fb5771d3`, 400_000},
		// XDEFIAggregatorEthGeneric
		{common.ETHChain, `0x53E4DD4072A9a8ed56289e048f5BD5AA51c9Bf6E`, 400_000},
		// XDEFIAggregatorEthUniswapV2
		{common.ETHChain, `0xeEe520b0DA1F8a9e4a0480F92CC4c5f6C027ef1E`, 400_000},
		// XDEFIAggregatorAvaxGeneric
		{common.AVAXChain, `0xd0269244A876F7Bc600D1f38B03a9916864b73C6`, 400_000},
		// XDEFIAggregatorAvaxTraderJoe
		{common.AVAXChain, `0x4ab34123A077aE294A39844f3e8df418d2A3D8c4`, 400_000},
		// XDEFIAggregatorUniswapV3 100 - short notation
		{common.ETHChain, `0x88100E08e5287bA3445F95d448ABfF3113d82a4C`, 400_000},
		// XDEFIAggregatorUniswapV3 500 - short notation
		{common.ETHChain, `0xC1faA12981160945903E0725888828E2d6a15821`, 400_000},
		// XDEFIAggregatorUniswapV3 3000 - short notation
		{common.ETHChain, `0x7E019988299cd8038091D8d7fe38f7a1dd3f90F1`, 400_000},
		// XDEFIAggregatorUniswapV3 10000 - short notation
		{common.ETHChain, `0x95B6b888a9fCc5BCA4A3004Df5E9498B63195F48`, 400_000},
		// TSAggregatorGeneric
		{common.BSCChain, `0xB6fA6f1DcD686F4A573Fd243a6FABb4ba36Ba98c`, 400_000},
		// TSAggregatorPancakeV2 BinanceSmartChain
		{common.BSCChain, `0x30912B38618D3D37De3191A4FFE982C65a9aEC2E`, 400_000},
		// TSAggregatorStargate Ethereum gen2 V1
		{common.ETHChain, `0x1204b5Bf0D6d48E718B1d9753A4166A7258B8432`, 800_000},
		// LayerZero Executor Ethereum
		{common.ETHChain, `0xe93685f3bBA03016F02bD1828BaDD6195988D950`, 800_000},
		// LayerZero Executor Avalanche
		{common.AVAXChain, `0xe93685f3bBA03016F02bD1828BaDD6195988D950`, 800_000},
		// LayerZero Executor BinanceSmartChain
		{common.BSCChain, `0xe93685f3bBA03016F02bD1828BaDD6195988D950`, 800_000},
		// TSLedgerAdapter
		{common.ETHChain, `0xaD545d047E34862121695F8E37aeEB71324a9E78`, 500_000},
		// SquidRouter MultiCall Ethereum
		{common.ETHChain, `0x4fd39C9E151e50580779bd04B1f7eCc310079fd3`, 800_000},
		// TSAggregatorStargate Ethereum gen2 V2
		{common.ETHChain, `0x48f68ff093b3b3A80D2FC97488EaD97E16b86283`, 800_000},
		// TSAggregatorUniswapV2 Ethereum gen2 V2.5 - tax tokens
		{common.ETHChain, `0x0fA226e8BCf45ec2f3c3163D2d7ba0d2aAD2eBcF`, 800_000},
		// RangoDiamond Ethereum
		{common.ETHChain, `0x69460570c93f9DE5E2edbC3052bf10125f0Ca22d`, 400_000},
		// RangoDiamond BSC
		{common.BSCChain, `0x69460570c93f9DE5E2edbC3052bf10125f0Ca22d`, 400_000},
		// RangoDiamond Avax
		{common.AVAXChain, `0x69460570c93f9DE5E2edbC3052bf10125f0Ca22d`, 400_000},
		// RangoThorchainOutputAggUniV3_COMPACT_Fee500
		{common.ETHChain, `0x70F75937546fB26c6FD3956eBBfb285f41526186`, 400_000},
		// RangoThorchainOutputAggUniV3_COMPACT_Fee3000
		{common.ETHChain, `0xd1687354CBA0e56facd0c44eD0F69D97F5734Dc1`, 400_000},
		// RangoThorchainOutputAggUniV3_COMPACT_Fee10000
		{common.ETHChain, `0xaFa4cBA6db85515f66E3ed7d6784e8cf5b689E2D`, 400_000},
		// RangoThorchainOutputAggUniV2_COMPACT_SUSHI
		{common.ETHChain, `0x0964347B0019eb227c901220ce7d66BB01479220`, 400_000},
		// RangoThorchainOutputAggUniV2_COMPACT_UNI
		{common.ETHChain, `0x6f281993AB68216F8898c593C4578C8a4a76F063`, 400_000},
		// RangoThorchainOutputAggUniV2_COMPACT_PANCAKE
		{common.BSCChain, `0xd0d7A5374ed70D5cB9E9034871F1d89F79De07Dd`, 400_000},
		// RangoThorchainOutputAggUniV3SwapRouter2_COMPACT_Fee500
		{common.BSCChain, `0x5bCAC8ac5f65623f8e151d676605EdE52E0Db532`, 400_000},
		// RangoThorchainOutputAggUniV3SwapRouter2_COMPACT_Fee3000
		{common.BSCChain, `0x36C29dC30E6728BC5524806EeA8897F6d8b9edE3`, 400_000},
		// RangoThorchainOutputAggUniV3SwapRouter2_COMPACT_Fee10000
		{common.BSCChain, `0xd1127EB3bc10a00434FfaD4fBA534212F1ba1165`, 400_000},
		// RangoThorchainOutputAggUniV2_COMPACT_TRADERJOE
		{common.AVAXChain, `0x892Fb7C2A23772f4A2FFC3DC82419147dC22021C`, 400_000},
		// RangoThorchainOutputAggUniV2_COMPACT_PANGOLIN
		{common.AVAXChain, `0xBd039a45e656221E28594d2761DDed8F6712AE46`, 400_000},
	}
}
