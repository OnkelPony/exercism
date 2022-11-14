package diffiehellman

import (
	"crypto/rand"
	"math/big"
)

// PrivateKey returns random value between (1, p)
func PrivateKey(p *big.Int) *big.Int {
	z := new(big.Int)
	z.Sub(p, big.NewInt(3))
	result, _ := rand.Int(rand.Reader, z)
	result.Add(result, big.NewInt(2))
	return result
}

// PublicKey returns g ** a mod p
func PublicKey(a *big.Int, p *big.Int, g int64) *big.Int {
	return SecretKey(a, big.NewInt(g), p)
}

// SecretKey returns B ** a mod p
func SecretKey(a *big.Int, B *big.Int, p *big.Int) *big.Int {
	result := new(big.Int)
	result.Exp(B, a, p)
	return result
}

// NewPair returns public and private keys based on p and g in parameter
func NewPair(p *big.Int, g int64) (*big.Int, *big.Int) {
	// a, A := new(big.Int), new(big.Int)
	a := PrivateKey(p)
	A := PublicKey(a, p, g)
	return a, A
}
