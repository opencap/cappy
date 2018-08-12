package key

import (
	"fmt"
	"golang.org/x/crypto/ed25519"
	"io"
	"io/ioutil"
)

type Key ed25519.PrivateKey

func (k Key) WriteTo(w io.Writer) (err error) {
	_, err = w.Write(k)
	return
}

func (k *Key) ReadFrom(r io.Reader) error {
	bs, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	if len(bs) != ed25519.PrivateKeySize {
		return fmt.Errorf("invalid key size")
	}

	*k = bs
	return nil
}

func (k Key) Public() string {
	kp := k[ed25519.PrivateKeySize-ed25519.PublicKeySize:]
	return fmt.Sprintf("%x", kp)
}

func (k Key) Identifier() string {
	id := k[ed25519.PrivateKeySize-4:]
	return fmt.Sprintf("0x%X", id)
}

func Identifier(key ed25519.PublicKey) string {
	id := key[ed25519.PublicKeySize-4:]
	return fmt.Sprintf("0x%X", id)
}
