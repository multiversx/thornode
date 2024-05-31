package mvxSupport

import (
	"encoding/hex"

	. "gopkg.in/check.v1"
)

type PubkeySuite struct{}

var _ = Suite(&PubkeySuite{})

func (s *PubkeySuite) TestInvalidBytesSlice(c *C) {
	encoded, err := EncodeBytesToMvXAddress(nil)
	c.Assert(encoded, Equals, "")
	c.Assert(err.Error(), Equals, "invalid bytes slice, expected 32, got 0")

	encoded, err = EncodeBytesToMvXAddress(make([]byte, 0))
	c.Assert(encoded, Equals, "")
	c.Assert(err.Error(), Equals, "invalid bytes slice, expected 32, got 0")

	encoded, err = EncodeBytesToMvXAddress(make([]byte, 1))
	c.Assert(encoded, Equals, "")
	c.Assert(err.Error(), Equals, "invalid bytes slice, expected 32, got 1")

	encoded, err = EncodeBytesToMvXAddress(make([]byte, 31))
	c.Assert(encoded, Equals, "")
	c.Assert(err.Error(), Equals, "invalid bytes slice, expected 32, got 31")

	encoded, err = EncodeBytesToMvXAddress(make([]byte, 33))
	c.Assert(encoded, Equals, "")
	c.Assert(err.Error(), Equals, "invalid bytes slice, expected 32, got 33")

	encoded, err = EncodeBytesToMvXAddress(make([]byte, 500))
	c.Assert(encoded, Equals, "")
	c.Assert(err.Error(), Equals, "invalid bytes slice, expected 32, got 500")
}

func (s *PubkeySuite) TestConvert(c *C) {
	aliceBytes, err := hex.DecodeString("0139472eff6886771a982f3083da5d421f24c29181e63888228dc81ca60d69e1")
	c.Assert(err, IsNil)

	alice, err := EncodeBytesToMvXAddress(aliceBytes)
	c.Assert(err, IsNil)
	c.Assert(alice, Equals, "erd1qyu5wthldzr8wx5c9ucg8kjagg0jfs53s8nr3zpz3hypefsdd8ssycr6th")

	bobBytes, err := hex.DecodeString("8049d639e5a6980d1cd2392abcce41029cda74a1563523a202f09641cc2618f8")
	c.Assert(err, IsNil)

	bob, err := EncodeBytesToMvXAddress(bobBytes)
	c.Assert(err, IsNil)
	c.Assert(bob, Equals, "erd1spyavw0956vq68xj8y4tenjpq2wd5a9p2c6j8gsz7ztyrnpxrruqzu66jx")
}
