package Paillier

import (
	"math/big"
)

func GenerateKeyPair(bits int) (pri PrivateKey, pub PublicKey) {
	for {
		p := GeneratePrime(bits>>1 - 1)
		q := GeneratePrime(bits>>1 + 1)
		if p.Cmp(q) != 0 {
			n := big.NewInt(0).Mul(p, q)
			return newPrivateKey(p, q, n), newPublicKey(n)
		}
	}
}
