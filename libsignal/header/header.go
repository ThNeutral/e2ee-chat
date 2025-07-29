package header

import (
	"chat/libsignal/dh"
	"encoding/json"
)

const (
	PublicKeyName           = "pub_key"
	PreviousChainLengthName = "pn"
	MessageNumber           = "n"
	AssociatedDataKey       = "ad"
)

func New(pair *dh.KeyPair, pn, n int) map[string]any {
	m := make(map[string]any)

	m[PublicKeyName] = pair.Public()
	m[PreviousChainLengthName] = pn
	m[MessageNumber] = n

	return m
}

func Concat(associatedData []byte, header map[string]any) []byte {
	header[AssociatedDataKey] = associatedData

	bytes, _ := json.Marshal(header)
	return bytes
}
