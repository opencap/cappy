package formats

import (
	"encoding/base32"
	"golang.org/x/crypto/blake2b"
)

type NanoFormatter struct {
	prefix string
	enc *base32.Encoding
}

func NewNanoFormatter(prefix string) *NanoFormatter {
	return &NanoFormatter{
		prefix: prefix,
		enc: base32.NewEncoding("13456789abcdefghijkmnopqrstuwxyz"),
	}
}

func (f *NanoFormatter) padLeft(bs []byte) []byte {
	pad := 5 - len(bs) % 5
	slice := make([]byte, pad + len(bs))
	copy(slice[pad:], bs)
	return slice
}

func (f *NanoFormatter) checksum(bs []byte) []byte {
	hash, _ := blake2b.New(5, nil)
	hash.Write(bs)
	check := hash.Sum(nil)

	for i := len(check)/2-1; i >= 0; i-- {
		opp := len(check)-1-i
		check[i], check[opp] = check[opp], check[i]
	}

	return check
}

func (f *NanoFormatter) Format(bs []byte) string {
	var (
		address = f.enc.EncodeToString(f.padLeft(bs))[4:]
		checksum = f.enc.EncodeToString(f.checksum(bs))
	)
	return f.prefix + address + checksum
}

func (f *NanoFormatter) Parse(str string) []byte {
	bs, err := f.enc.DecodeString(str)
	if err != nil {
		return nil
	}

	return bs
}


