package ipfs

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"github.com/ipfs/go-ipfs/repo/config"
	libp2p "gx/ipfs/QmPGxZ1DP2w45WcogpW1h43BvseXbfke9N91qotpoQcUeS/go-libp2p-crypto"
	peer "gx/ipfs/QmWUswjn261LSyVxWAEpMVtPdy8zmKBJJfBpG3Qdpa8ZsE/go-libp2p-peer"
)

func IdentityFromKey(privkey []byte) (config.Identity, error) {

	ident := config.Identity{}
	sk, err := libp2p.UnmarshalPrivateKey(privkey)
	if err != nil {
		return ident, err
	}
	skbytes, err := sk.Bytes()
	if err != nil {
		return ident, err
	}
	ident.PrivKey = base64.StdEncoding.EncodeToString(skbytes)

	id, err := peer.IDFromPublicKey(sk.GetPublic())
	if err != nil {
		return ident, err
	}
	ident.PeerID = id.Pretty()
	return ident, nil
}

func IdentityKeyFromSeed(seed []byte, bits int) ([]byte, error) {
	hmac := hmac.New(sha256.New, []byte("OpenBazaar seed"))
	hmac.Write(seed)
	reader := bytes.NewReader(hmac.Sum(nil))
	sk, _, err := libp2p.GenerateKeyPairWithReader(libp2p.Ed25519, bits, reader)
	if err != nil {
		return nil, err
	}
	encodedKey, err := sk.Bytes()
	if err != nil {
		return nil, err
	}
	return encodedKey, nil
}
