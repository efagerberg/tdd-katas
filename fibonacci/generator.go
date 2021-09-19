package fibonacci

func fibonacci() func() int {
	a := 0
	b := 1
	f := func() int {
		s := a
		a = b
		b = s + b
		return s
	}

	return f
}
