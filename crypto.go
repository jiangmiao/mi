package mi

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
)

func HEX(bs []byte) string {
	return hex.EncodeToString(bs)
}

func MD5(bs []byte) string {
	v := md5.Sum(bs)
	return HEX(v[:])
}

func SHA256(bs []byte) string {
	return HEX(SHA256BIN(bs))
}

func SHA256BIN(bs []byte) []byte {
	r := sha256.Sum256(bs)
	return r[:]
}

func SHA512(bs []byte) string {
	v := sha512.Sum512(bs)
	return HEX(v[:])
}
