package main

import (
	"crypto/ecdsa"
	"fmt"
	"log"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main(){
	privateKey,err :=crypto.GenerateKey()
	// privateKey,err:=crypto.HexToECDSA("fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")
	if err != nil {
        log.Fatal(err)
    }
	privateKeyBytes := crypto.FromECDSA(privateKey)
	fmt.Println("privateKey:")
	fmt.Println(hexutil.Encode(privateKeyBytes)[2:])
	publicKey := privateKey.Public()
	publicKeyECDSA, ok:= publicKey.(*ecdsa.PublicKey)
	if !ok {
        log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
    }
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println("publicKeyKey:")
	fmt.Println(hexutil.Encode(publicKeyBytes)[4:])
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println("address:")
	fmt.Println(address)
}
