package src

import (
	"fmt"
	"math/big"
	"math/rand"
)

func gcd(a *big.Int, b any) *big.Int {
	var bigB *big.Int

	switch b := b.(type) {
	case int:
		bigB = big.NewInt(int64(b))
	case *big.Int:
		bigB = b
	default:
		panic("Unsupported type for b")
	}

	result := new(big.Int)
	result.GCD(nil, nil, a, bigB)
	return result
}

func publicKeyCondition(e, Ctn *big.Int) (bool, bool) {
	firstcondition := false
	secondcondition := false
	ones := *big.NewInt(1)
	/* first condition check if 1<e < CTN */
	if e.Cmp(&ones) == 1 && Ctn.Cmp(e) == 1 {
		firstcondition = true
	}

	/* second condition check if gcd(e,ctn) ==1 */
	fmt.Println(gcd(Ctn, e))
	if gcd(Ctn, e).Cmp(&ones) == 0 {
		secondcondition = true
	}
	return firstcondition, secondcondition
}

/*
Common choices for
e in practice are relatively small prime numbers like 3, 17, or 65537. These numbers are chosen because they offer a good balance between security and computational efficiency.
*/

func GenerateRandomE() *big.Int {

	selected := []int{3, 7, 65537}
	selected = append(selected, sieveOfEratosthenes(65537)...)
	n := rand.Int() % len(selected)

	return big.NewInt(int64(selected[n]))
}

// return list of primes less than N
func sieveOfEratosthenes(N int) (primes []int) {
	b := make([]bool, N)
	for i := 2; i < N; i++ {
		if b[i] {
			continue
		}
		primes = append(primes, i)
		for k := i * i; k < N; k += i {
			b[k] = true
		}
	}
	return
}

func modInverse(e, phi *big.Int) big.Int {

	x := new(big.Int)
	y := new(big.Int)

	// Extended Euclidean Algorithm to find d such that ed ≡ 1 (mod phi)
	gcd := new(big.Int).GCD(x, y, e, phi)

	// d is only valid if gcd(e, phi) == 1
	if gcd.Cmp(big.NewInt(1)) != 0 {
		return *big.NewInt(0) // no modular inverse if e and phi are not coprime
	}

	// Make sure d is positive
	if x.Cmp(big.NewInt(0)) < 0 {
		x.Add(x, phi)
	}

	return *x
}
