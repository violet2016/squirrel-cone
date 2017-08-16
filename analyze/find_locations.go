package analyze

import (
	"bytes"
	"encoding/hex"
)

func SearchForSig(region []byte) func(string) int {
	return func(sig string) int {
		sigHex, err := hex.DecodeString(sig)
		if err != nil {
			return -1
		}
		return bytes.Index(region, sigHex)
	}
}
