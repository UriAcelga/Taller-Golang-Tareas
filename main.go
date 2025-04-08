package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	f, err := os.Open("puntajes.txt")
	if err != nil {
		fmt.Println("Error al abrir puntajes.txt: ", err)
		return
	}
	defer f.Close()

	n, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println("Error al leer puntajes: ", err)
	}

	puntajes := map[string]int{
		"1": 0,
		"2": 0,
		"3": 0,
		"4": 0,
		"5": 0,
	}

	for _, v := range string(n) {
		puntajes[string(v)]++
	}

	fmt.Print(
		"Puntajes:",
		"\n1: ", puntajes["1"],
		"\n2: ", puntajes["2"],
		"\n3: ", puntajes["3"],
		"\n4: ", puntajes["4"],
		"\n5: ", puntajes["5"],
		"\n",
	)

	malosPuntajes := puntajes["1"] + puntajes["2"]
	buenosPuntajes := puntajes["4"] + puntajes["5"]

	if buenosPuntajes > malosPuntajes {
		fmt.Println("¡Buen Resultado!")
	} else {
		fmt.Println("Resultado Mejorable")
	}

}
