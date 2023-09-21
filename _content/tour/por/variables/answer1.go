//go:build OMIT

// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare três variáveis que são inicializadas com seu zero value e três
// declaradas com um valor literal. Declare variáveis do tipo string, int e
// bool. Exiba os valores dessas variáveis.
//
// Declare uma nova variável do tipo float32 e inicialize a variável
// convertendo o valor literal de Pi (3.14).
package main

import "fmt"

func main() {

	// Declare variáveis com seu zero value.
	var age int
	var name string
	var legal bool

	// Exiba o valor dessas variáveis.
	fmt.Println(age)
	fmt.Println(name)
	fmt.Println(legal)

	// Declare variáveis e inicialize.
	// Usando o operador de declaração de variável curto.
	month := 10
	dayOfWeek := "Tuesday"
	happy := true

	// Exiba o valor dessas variáveis.
	fmt.Println(month)
	fmt.Println(dayOfWeek)
	fmt.Println(happy)

	// Faça uma conversão de tipo.
	pi := float32(3.14)

	// Exiba o valor dessa variável.
	fmt.Printf("%T [%v]\n", pi, pi)
}
