package complexnumbers

// Define the Number type here.
type Number struct {
	real float64
	imag float64
}

func (n Number) Real() float64 {
	return n.real
}

func (n Number) Imaginary() float64 {
	return n.imag
}

func (n1 Number) Add(n2 Number) Number {
	return Number{n1.real + n2.real, n1.imag + n2.imag}
}

func (n1 Number) Subtract(n2 Number) Number {
	return Number{n1.real - n2.real, n1.imag - n2.imag}
}

func (n1 Number) Multiply(n2 Number) Number {
	return Number{n1.real*n2.real - n1.imag*n2.imag, n1.imag*n2.real + n1.real*n2.imag}
}

func (n Number) Times(factor float64) Number {
	return n.Multiply(Number{factor, 0})
}

func (n1 Number) Divide(n2 Number) Number {
	real := (n1.real*n2.real + n1.imag*n2.imag) / (n2.real*n2.real + n2.imag*n2.imag)
	imag := (n1.imag*n2.real - n1.real*n2.imag) / (n2.real*n2.real + n2.imag*n2.imag)
	return Number{real, imag}
}

func (n Number) Conjugate() Number {
	panic("Please implement the Conjugate method")
}

func (n Number) Abs() float64 {
	panic("Please implement the Abs method")
}

func (n Number) Exp() Number {
	panic("Please implement the Exp method")
}
