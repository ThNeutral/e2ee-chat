package x3dh

import (
	"crypto/ecdh"
	"crypto/sha512"
)

var curve = ecdh.X25519

var hash = sha512.New

var info = []byte("x3dh")
