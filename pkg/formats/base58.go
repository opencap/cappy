package formats

import (
	"github.com/btcsuite/btcutil/base58"
	"crypto/sha256"
)

type Base58Formatter struct {
	prefix byte
}

func NewBase58Formatter(prefix byte) *Base58Formatter {
	return &Base58Formatter{
		prefix: prefix,
	}
}

func (f *Base58Formatter) checksum(bs []byte) []byte {
	hashed1 := sha256.Sum256(bs)
	hashed2 := sha256.Sum256(hashed1[:])
	return hashed2[:4]
}

func (f *Base58Formatter) Format(bs []byte) string {
	addr := []byte{f.prefix}
	addr = append(addr, bs...)

	checksum := f.checksum(addr)
	addr = append(addr, checksum...)

	return base58.Encode(addr)
}

func (f *Base58Formatter) Parse(str string) []byte {
	addr := base58.Decode(str)
	if len(addr) != 25 {
		return nil
	}

	if addr[0] != f.prefix {
		return nil
	}
	addr = addr[1:]

	checksum1 := addr[len(addr)-4:]
	addr = addr[:len(addr)-4]
	checksum2 := f.checksum(addr)

	if checksum1[0] != checksum2[0] || checksum1[1] != checksum2[1] || checksum1[2] != checksum2[2] || checksum1[3] != checksum2[3] {
		return nil
	}

	return addr
}

