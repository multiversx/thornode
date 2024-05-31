package mvxSupport

import (
	"fmt"

	"github.com/btcsuite/btcutil/bech32"
)

type config struct {
	prefix   string
	fromBits byte
	toBits   byte
	pad      bool
}

var bech32Config = config{
	fromBits: byte(8),
	toBits:   byte(5),
	pad:      true,
}

const bytesSliceExpectedSize = 32
const mvxHRP = "erd"

// EncodeBytesToMvXAddress will encode the provided bytes to a MultiversX bech32 address (erd1....)
// The bytes slice provided as parameter should be exactly 32 bytes long, the function errors otherwise
func EncodeBytesToMvXAddress(bytes []byte) (string, error) {
	if len(bytes) != bytesSliceExpectedSize {
		return "", fmt.Errorf("invalid bytes slice, expected %d, got %d", bytesSliceExpectedSize, len(bytes))
	}

	conv, err := bech32.ConvertBits(bytes, bech32Config.fromBits, bech32Config.toBits, bech32Config.pad)
	if err != nil {
		return "", fmt.Errorf("%w: %s", fmt.Errorf("invalid fromBits or toBits when converting bits"), err.Error())
	}

	encodedBytes, err := bech32.Encode(mvxHRP, conv)
	if err != nil {
		return "", fmt.Errorf("can't convert bech32 string: %w", err)
	}

	return encodedBytes, nil
}
