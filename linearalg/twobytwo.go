package linearalg 

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

func Add(M, N matrix2x2) matrix2x2 {
	return matrix2x2{M.A11 + N.A11, M.A12 + N.A12, M.A21 + N.A21, M.A22 + N.A22}
}

func (M matrix2x2) MatrixMult(N matrix2x2) matrix2x2 {
	return matrix2x2{M.A11*N.A11 + M.A12*N.A21, M.A11*N.A12 + M.A12*N.A22, 
	M.A21*N.A11 + M.A22*N.A21, M.A21*N.A12+M.A22+M.A22*N.A22}
}

func (M matrix2x2) VectorMult(W vector2) vector2 {
	return vector2{M.A11*W.V1 + M.A12*W.V2, M.A21*W.V1 + M.A22*W.V2}
}

func (M matrix2x2) Det() int {
	return M.A11*M.A22 - M.A12*M.A21
}

func (M matrix2x2) Adjugate() matrix2x2 {
	return matrix2x2{M.A22, -M.A12, -M.A21, M.A11}
}
