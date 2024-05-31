//go:build !stagenet && !mocknet
// +build !stagenet,!mocknet

package common

import (
	"encoding/hex"

	xed25519 "golang.org/x/crypto/ed25519"
	. "gopkg.in/check.v1"

	"github.com/cosmos/cosmos-sdk/crypto/codec"
	"github.com/tendermint/tendermint/crypto/ed25519"
	"github.com/tendermint/tendermint/crypto/secp256k1"
	"gitlab.com/thorchain/thornode/common/cosmos"
)

func (s *PubKeyTestSuite) TestPubKeyGetAddress(c *C) {
	for _, d := range s.keyData {
		privB, _ := hex.DecodeString(d.priv)
		pubB, _ := hex.DecodeString(d.pub)
		priv := secp256k1.PrivKey(privB)
		pubKey := priv.PubKey()
		pubT, _ := pubKey.(secp256k1.PubKey)
		pub := pubT[:]

		c.Assert(hex.EncodeToString(pub), Equals, hex.EncodeToString(pubB))

		tmp, err := codec.FromTmPubKeyInterface(pubKey)
		c.Assert(err, IsNil)
		pubBech32, err := cosmos.Bech32ifyPubKey(cosmos.Bech32PubKeyTypeAccPub, tmp)
		c.Assert(err, IsNil)

		pk, err := NewPubKey(pubBech32)
		c.Assert(err, IsNil)

		addrETH, err := pk.GetAddress(ETHChain)
		c.Assert(err, IsNil)
		c.Assert(addrETH.String(), Equals, d.addrETH.mainnet)

		addrBNB, err := pk.GetAddress(BNBChain)
		c.Assert(err, IsNil)
		c.Assert(addrBNB.String(), Equals, d.addrBNB.mainnet)

		addrBTC, err := pk.GetAddress(BTCChain)
		c.Assert(err, IsNil)
		c.Assert(addrBTC.String(), Equals, d.addrBTC.mainnet)

		addrLTC, err := pk.GetAddress(LTCChain)
		c.Assert(err, IsNil)
		c.Assert(addrLTC.String(), Equals, d.addrLTC.mainnet)

		addrBCH, err := pk.GetAddress(BCHChain)
		c.Assert(err, IsNil)
		c.Assert(addrBCH.String(), Equals, d.addrBCH.mainnet)

		addrDOGE, err := pk.GetAddress(DOGEChain)
		c.Assert(err, IsNil)
		c.Assert(addrDOGE.String(), Equals, d.addrDOGE.mainnet)
	}
}

func (s *PubKeyTestSuite) TestPubKeyGetAddressForEd25519(c *C) {
	for _, d := range s.keyDataEd25519 {
		privB, _ := hex.DecodeString(d.priv)
		pubB, _ := hex.DecodeString(d.pub)
		priv := ed25519.PrivKey(xed25519.NewKeyFromSeed(privB))
		pubKey := priv.PubKey()
		pubT, _ := pubKey.(ed25519.PubKey)
		pub := pubT[:]

		c.Assert(hex.EncodeToString(pub), Equals, hex.EncodeToString(pubB))

		tmp, err := codec.FromTmPubKeyInterface(pubKey)
		c.Assert(err, IsNil)
		pubBech32, err := cosmos.Bech32ifyPubKey(cosmos.Bech32PubKeyTypeAccPub, tmp)
		c.Assert(err, IsNil)

		pk, err := NewPubKey(pubBech32)
		c.Assert(err, IsNil)

		addrMVX, err := pk.GetAddress(MVXChain)
		c.Assert(err, IsNil)
		c.Assert(addrMVX.String(), Equals, d.addrMVX.mainnet)
	}
}
