// funcao tambem é chamdo de procedimento ou sub-rotina
// parte do codigo
//ela pega um dado de entrada e retorna um dado de saida

package main

import "fmt"

func media(lista []float64) float64 { // função que calcula a média de uma sala

	total := 0.0 // variavel que vai receber a soma das notas

	for _, valor := range lista { // for que vai somar as notas

		total += valor
	}

	return total / float64(len(lista)) // retorna a media

}

func main() { // programa que calcula a média de uma sala

	 lista := []float64{98, 93, 77, 50, 80} // lista de notas
	 fmt.Println(media(lista)) // imprime a media

	// total := 0.0 // variavel que vai receber a soma das notas

	// for _, valor := range lista { // for que vai somar as notas

	// 	total += valor
	// }

	// fmt.Println(total / float64(len(lista))) // imprime a media

}