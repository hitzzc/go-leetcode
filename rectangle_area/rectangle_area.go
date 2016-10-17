package rectangle_area

func computeArea(A int, B int, C int, D int, E int, F int, G int, H int) int {
	if C <= E || G <= A || D <= F || H <= B {
		return (C-A)*(D-B) + (G-E)*(H-F)
	}
	left := max(A, E)
	bottom := max(B, F)
	right := min(C, G)
	top := min(D, H)
	return (C-A)*(D-B) + (G-E)*(H-F) - (right-left)*(top-bottom)
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
