//go:build !stagenet && !mocknet
// +build !stagenet,!mocknet

package aggregators

import (
	"gitlab.com/thorchain/thornode/common"
)

func DexAggregatorsV101() []Aggregator {
	return []Aggregator{
		// TSAggregatorGeneric
		{common.ETHChain, `0xd31f7e39afECEc4855fecc51b693F9A0Cec49fd2`, 400_000},
		// TSAggregator2LegUniswapV2 USDC
		{common.ETHChain, `0x3660dE6C56cFD31998397652941ECe42118375DA`, 400_000},
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
		{common.ETHChain, `0x82ff5842776ea577fd5247d3e97c88f7a0386620`, 400_000},
		// XDEFIAggregatorAvaxGeneric
		{common.AVAXChain, `0xd0269244A876F7Bc600D1f38B03a9916864b73C6`, 400_000},
		// XDEFIAggregatorAvaxPangolin
		{common.AVAXChain, `0x1462F79C2D4B308EB8CB6938F4d2cC9f85Dde31A`, 400_000},
	}
}
