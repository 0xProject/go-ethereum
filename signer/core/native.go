// +build !js

package core

import (
	"fmt"

	"github.com/ethereum/go-ethereum/consensus/clique"
	"github.com/ethereum/go-ethereum/core/types"
)

// cliqueHeaderHashAndRlp returns the hash which is used as input for the proof-of-authority
// signing. It is the hash of the entire header apart from the 65 byte signature
// contained at the end of the extra data.
//
// The method requires the extra data to be at least 65 bytes -- the original implementation
// in clique.go panics if this is the case, thus it's been reimplemented here to avoid the panic
// and simply return an error instead
func cliqueHeaderHashAndRlp(header *types.Header) (hash, rlp []byte, err error) {
	if len(header.Extra) < 65 {
		err = fmt.Errorf("clique header extradata too short, %d < 65", len(header.Extra))
		return
	}
	rlp = clique.CliqueRLP(header)
	hash = clique.SealHash(header).Bytes()
	return hash, rlp, err
}
