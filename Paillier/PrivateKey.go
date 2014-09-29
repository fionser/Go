package Paillier

import (
	"math/big"
)

type PrivateKey struct {
	l *big.Int
	m *big.Int
}

func newPrivateKey(p, q, n *big.Int) PrivateKey {
	one := big.NewInt(1)
	p_1 := big.NewInt(0).Sub(p, one)
	q_1 := big.NewInt(0).Sub(q, one)
	l := big.NewInt(0).Mul(p_1, q_1)
	m := big.NewInt(0).ModInverse(l, n)
	return PrivateKey{l, m}
}
