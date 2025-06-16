package main

import (
	"fmt"
	"slices"
	"sort"
)

func main() {
	lista := []float64{5.2, 3.1, 7.4, 1.8, 9.0}

	sort.Float64s(lista)
	fmt.Println("Esta es una lista ordenada:", lista)

	slices.Reverse(lista)
	fmt.Println("Esta es una lista invertida:", lista)

	slice := make([]float64, len(lista))
	copy(slice, lista)
	fmt.Println("Esto es un slice copiado:", slice)

	slice = append(slice, 6.5)
	fmt.Println("Esto es un slice después de agregar un valor:", slice)

	indice := 2
	slice = append(slice[:indice], slice[indice+1:]...)
	fmt.Println("Este es un slice después de eliminar el valor en índice", indice, ":", slice)

	valorBuscar := 7.4
	encontrado := false
	for _, v := range slice {
		if v == valorBuscar {
			encontrado = true
			break
		}
	}

	if encontrado {
		fmt.Println("El ", valorBuscar, "está en el slice.")
	} else {
		fmt.Println("El ", valorBuscar, "no está en el slice.")
	}
}
