package library

func ModN(i, N int) int {
	m := i % N 
	if m < 0 {
		m += N 
	}
	return m
}

func FastPower(N, g, A int) int {
	var b int 
	a := g 
	b = 1 
	for A > 0 {
		if A % 2 == 1 {
			b = ModN(b*a, N)
		} 
		a = ModN(a*a, N) 
		A = A / 2 
	}
	return b
}
