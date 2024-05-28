package main

import (
	"fmt"
	"math/big"

	"github.com/Ahmed-Mahmoud-M/RSA-implementation/src"
)

func main() {
	rsa := &src.RSA{}
	var p big.Int

	var q big.Int
	p.SetString("61", 10) // base 10
	q.SetString("53", 10)
	rsa.NewRsa(p, q)

	fmt.Println(rsa.GeneratePublicKey())
	fmt.Print(rsa.GeneratePrivateKey())

	fmt.Println(rsa.Getprivatekey())
	fmt.Println(rsa.Encrypter(*big.NewInt(65)))
	fmt.Println(rsa.Decrypter(*big.NewInt(2790)))
}
