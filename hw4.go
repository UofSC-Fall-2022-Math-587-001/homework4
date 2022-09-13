package hillcipher 

import (
	"github.com/UofSC-Fall-2022-Math-587-001/homework4/linearalg"
)

// Recall that a Hill cipher has as a key a pair (A,v) where A is a nxn invertible 
// matrix over some field F_p and v is a n-dimensional vector. The message space 
// is F_p. To encrypt a message m, we send it to Am + v. To decrypt enciphered 
// text c, we evaluate A^{-1}(c - v) where A^{-1} is the inverse to the matrix A. 

// We import a very basic linear algebra library from a subfolder. 

type HillCipherKey struct {
	// Fill in here
}

func (H HillCipherKey) encrypt(W linearalg.Vector2, N int) linearalg.Vector2 {
	// Fill in here 
	return linearalg.Vector2{}
}

func (H HillCipherKey) decrypt(W linearalg.Vector2, N int) linearalg.Vector2 {
	// Fill in here 
	return linearalg.Vector2{}
}

