package Paillier

import (
	"crypto/rand"
	"math/big"
    "log"
)

func Copy(a *big.Int) *big.Int {
	return big.NewInt(0).Set(a)
}

func GeneratePrime(bits int) *big.Int {
	p, err := rand.Prime(rand.Reader, bits)
	if err == nil {
		return p
	} else {
        log.Fatal("Security Random Failed")
		return nil
	}
}

func Add(pk PublicKey, a, b *big.Int) *big.Int {
    c := big.NewInt(0).Mul(a, b)
	c.Mod(c, pk.N_sq)
	return c
}

// Homomorphically add b into a and return a
func MutableAdd(pk PublicKey, a, b *big.Int) *big.Int {
    a.Mul(a, b)
    a.Mod(a, pk.N_sq)
    return a
}

func MulConst(pk PublicKey, ctx *big.Int, cst int64) *big.Int {
    return MulBigConst(pk, ctx, big.NewInt(cst))
}

func MulBigConst(pk PublicKey, ctx *big.Int, cst *big.Int) *big.Int {
	return big.NewInt(0).Exp(ctx, cst, pk.N_sq)
}

func AddConst(pk PublicKey, ctx *big.Int, cst int64) *big.Int {
    return MutableAddBigConst(pk, ctx, big.NewInt(cst))
}

func AddBigConst(pk PublicKey, ctx, cst *big.Int) *big.Int {
    return MutableAddBigConst(pk, ctx, big.NewInt(0).Set(cst))
}

func MutableAddBigConst(pk PublicKey, ctx, cst *big.Int) *big.Int {
    cst.Exp(pk.G, cst, pk.N_sq)
    cst.Mul(cst, ctx)
    return cst.Mod(cst, pk.N_sq)
}

func SubBigConst(pk PublicKey, ctx, cst *big.Int) *big.Int {
    res := big.NewInt(0).Exp(pk.G, cst, pk.N_sq)
    res.ModInverse(res, pk.N_sq)
    return res.Mul(res, ctx)
}

func ModInverse(ctx, n*big.Int) *big.Int {
    return big.NewInt(0).ModInverse(ctx, n)
}

func MulMod(a, b, n *big.Int) *big.Int {
    c := big.NewInt(0).Set(a)
    c.Mul(c, b)
    c.Mod(c, n)
    return c
}
