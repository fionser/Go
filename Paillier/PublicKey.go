package Paillier

import (
	"math/big"
    "fmt"
)

type PublicKey struct {
	N    *big.Int
	N_sq *big.Int
	G    *big.Int
}

func newPublicKey(n *big.Int) PublicKey {
	return PublicKey{N: Copy(n), N_sq: big.NewInt(0).Mul(n, n),
		G: big.NewInt(0).Add(n, big.NewInt(1))}
}

func (this *PublicKey) String() string {
    return fmt.Sprintf("n:%v, g:%v", this.N, this.G)
}
