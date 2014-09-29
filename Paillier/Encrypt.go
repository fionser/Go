package Paillier

import (
	"math/big"
    crand "crypto/rand"
    "math/rand"
)

func Encrypt(pk PublicKey, plain int) *big.Int {
    return EncryptBig(pk, big.NewInt(int64(plain)))
}

func EncryptBig(pk PublicKey, plain *big.Int) *big.Int {
    copy := big.NewInt(0).Set(plain)
    return MutableEncryptBig(pk, copy)
}

func MutableEncryptBig(pk PublicKey, plain *big.Int) *big.Int {
    r, err := crand.Int(crand.Reader, pk.N)
    if err != nil {
        r = big.NewInt(rand.Int63())
    }
    x := r.Exp(r, pk.N, pk.N_sq)
    cipher := plain.Exp(pk.G, plain, pk.N_sq)
    cipher.Mul(cipher, x).Mod(cipher, pk.N_sq)
    return cipher
}
