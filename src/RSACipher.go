package src

import (
	"fmt"
	"math/big"
)

type RSA struct {
	P_rime *big.Int
	Q_rime *big.Int

	N *big.Int // n is used as the modulus for both the public and private keys. = p_rime * q_rime

	Ctn *big.Int //Compute the Carmichael's totient function of the product as λ(n) = lcm(p − 1, q − 1)

	E *big.Int // Choose an integer e such that 1 < e < λ(n) and gcd(e, λ(n)) = 1
	D *big.Int //Compute d, the modular multiplicative inverse of e (mod λ(n)),

	public_key  string
	private_key string
}

func (r *RSA) Getprivatekey() string {
	return "the private key is {d,n} with values :\nd:" + r.D.String() + "\nn :" + r.N.String()
}

func (r *RSA) Getpublickey() string {
	return "the public key is {e,n} with values :\ne: " + r.E.String() + "\nn : " + r.N.String()
}

func (r *RSA) NewRsa(p, q big.Int) *RSA {
	var result, num, num2, tem big.Int
	r.D = big.NewInt(0)
	r.P_rime = &p
	r.Q_rime = &q

	result.Mul(r.P_rime, r.Q_rime)
	r.N = &result

	smallInt := big.NewInt(1)
	tem.Mul(num.Sub(r.P_rime, smallInt), num2.Sub(r.Q_rime, smallInt))
	r.Ctn = &tem

	r.private_key = r.Getprivatekey()
	r.public_key = r.Getpublickey()

	return r

}

func (r *RSA) GeneratePublicKey() string {

	fmt.Println("generating public key...")
	//fmt.Println("1. enter public key value ")
	fmt.Println(" generate random number for the public key...\n --")

	r.E = GenerateRandomE()

	return r.Getpublickey()
}

func (r *RSA) GeneratePrivateKey() string {
	result := modInverse(r.E, r.Ctn)
	r.D = &result
	return "generating private key successfully\n"
}
func (r *RSA) Encrypter(message big.Int) any {
	result := big.NewInt(10)
	return result.Exp(&message, r.E, r.N)
}

func (r *RSA) Decrypter(ciphertext big.Int) any {
	result := big.NewInt(10)
	return result.Exp(&ciphertext, r.D, r.N)
}
