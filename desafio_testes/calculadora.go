package desafio_testes

func sum(i ...int) int {
	result := 0
	for _, number := range i {
		result += number
	}
	return result
}

func substraction(i ...int) int {
	result := 0
	for _, number := range i {
		result -= number
	}
	return result
}

func multiplication(i ...int) int {
	result := 1
	for _, number := range i {
		result *= number
	}
	return result
}

func division(numerator float64, denominator float64) float64 {
	if denominator == 0 {
		return 0
	}

	return numerator / denominator
}
