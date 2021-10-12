package error

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

func TestError() {
	res, err := http.Get("http://www.google.com")
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("resultado http", res.Body)
}

func SomaError(a int, b int) (int, error) {
	result := a + b
	if result > 10 {
		return 0, errors.New("Total maior do 10")
	}
	return result, nil
}
