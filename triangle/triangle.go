package triangle

type Kind int

func (k Kind) String() string {
	switch k {
	case 0:
		return "Equilateral"
	case 1:
		return "Isosceles"
	case 2:
		return "Scalene"
	default:
		return "isn't a triangle"
	}
}

func (k Kind) IsTriangle(a, b, c float64) bool {
	return a+b > c && a+c > b && b+c > a
}

func (k Kind) IsEquilateral(a, b, c float64) bool {
	return k.IsTriangle(a, b, c) && a == c && b == c
}

func (k Kind) IsIsosceles(a, b, c float64) bool {
	return k.IsTriangle(a, b, c) && (a == b || b == c || c == a)
}

func (k Kind) IsScalene(a, b, c float64) bool {
	return k.IsTriangle(a, b, c) && a != b && b != c && c != a
}

func (k Kind) TypeTriangle(a, b, c float64) Kind {
	switch {
	case k.IsEquilateral(a, b, c):
		return Equ
	case k.IsScalene(a, b, c):
		return Sca
	case k.IsIsosceles(a, b, c):
		return Iso
	default:
		return NaT
	}
}

const (
	NaT Kind = -1
	Equ Kind = 0
	Iso Kind = 1
	Sca Kind = 2
)

func KindFromSides(a, b, c float64) Kind {
	var k Kind
	return k.TypeTriangle(a, b, c)
}
