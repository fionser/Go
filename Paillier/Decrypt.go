package Paillier

import (
	"math/big"
)

func Decrypt(priK PrivateKey, pk PublicKey, ciper *big.Int) *big.Int {
	x := big.NewInt(0).Exp(ciper, priK.l, pk.N_sq)
	x.Sub(x, big.NewInt(1))
	plain := big.NewInt(0).Quo(x, pk.N)
	plain.Mul(plain, priK.m)
	plain.Mod(plain, pk.N)
	return plain
}
