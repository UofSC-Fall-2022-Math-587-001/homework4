package linearalg

import "github.com/UofSC-Fall-2022-Math-587-001/homework4/library"

type Matrix2x2 struct {
	A11 int
	A12 int
	A21 int
	A22 int
}

type Vector2 struct {
	V1 int
	V2 int
}

func (W Vector2) ModN(N int) Vector2 {
	return Vector2{library.ModN(W.V1,N), library.ModN(W.V2, N)}
}

func (M Matrix2x2) ModN(N int) Matrix2x2 {
	return Matrix2x2{library.ModN(M.A11,N),library.ModN(M.A12,N),library.ModN(M.A21,N),library.ModN(M.A22,N)}
}

func (W Vector2) VectorAdd(U Vector2, N int) Vector2 {
	X := Vector2{W.V1 + U.V1, W.V2 + U.V2}
	return X.ModN(N)
}

func (W Vector2) ScalarMult(c, N int) Vector2 {
	U := Vector2{c * W.V1, c * W.V2}
	return U.ModN(N)
}

func (M Matrix2x2) VectorMult(W Vector2, N int) Vector2 {
	U := Vector2{M.A11*W.V1 + M.A12*W.V2, M.A21*W.V1 + M.A22*W.V2}
	return U.ModN(N)
}

func (M Matrix2x2) det() int {
	return M.A11*M.A22 - M.A12*M.A21
}

func (M Matrix2x2) Inverse(N int) Matrix2x2 {
	det := M.det()
	if det % N == 0 {
		return Matrix2x2{0,0,0,0}
	}
	detinv := library.FastPower(N,det,N-2)
	Minv := Matrix2x2{detinv*M.A22,-detinv*M.A12,-detinv*M.A21, 
	detinv*M.A11}
	return Minv.ModN(N)
}
