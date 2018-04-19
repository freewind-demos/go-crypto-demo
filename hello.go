package main

import (
	"github.com/tendermint/go-crypto"
	"fmt"
)

func main() {
	fmt.Println("-------------- generate private and public key ----------------------")
	privateKey := crypto.GenPrivKeyEd25519()
	fmt.Printf("Private key: %02X\n", privateKey)
	pubKey1 := privateKey.PubKey()
	fmt.Printf("Public key1: %02X\n", pubKey1)
	pubKey2 := privateKey.PubKey()
	fmt.Printf("Public key2: %02X\n", pubKey2)

}
