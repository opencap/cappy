package keys

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/opencap/cappy/internal/pkg/context"
	"github.com/opencap/cappy/internal/pkg/key"
	"golang.org/x/crypto/ed25519"
)

func Add() (k key.Key, err error) {
	var ks ed25519.PrivateKey
	_, ks, err = ed25519.GenerateKey(nil)
	if err != nil {
		return
	}

	k = key.Key(ks)
	if err = context.Instance().KeyManager().Write(k); err != nil {
		return
	}

	color.New(color.Bold).Println("\xF0\x9F\x8E\x89  We have created a key file for you!")
	fmt.Print("A key file is required to proof ownership of a domain when adding accounts and can be used for increased security using DNS signatures. ")
	fmt.Println("More information about DNS signatures can be found here: [insert link here]")
	fmt.Println()
	color.New(color.Bold).Print("Key ID: ")
	fmt.Println(k.Identifier())

	return
}
