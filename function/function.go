package function

func SomaFunc(a int, b int) (result int) {
	result = a + b
	return
}

func SomaVariadic(variables ...int) (result int) {
	// result = 0
	// for i := 0; i < len(variables); i++ {
	// 	result = result + variables[i]
	// }

	for _, value := range variables {
		result = result + value
	}
	return
}

