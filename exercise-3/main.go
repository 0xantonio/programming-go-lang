package main

import "fmt"

// How does Go handle pointers?
// What is the difference between passing a pointer and passing a value?

// What is the difference between these two functions?


func main() {
	ciao := "Ciao"
	fmt.Println("Posizione in memoria di ciao:")
	fmt.Println(&ciao)
	fmt.Println("Valore di ciao:")
	fmt.Println(ciao)
	
	puntatoreACiao := &ciao

	fmt.Println("Posizione in memoria del puntatore puntatoreACiao:")
	fmt.Println(&puntatoreACiao)
	cambiaStringaPuntatore(puntatoreACiao)
	fmt.Println(ciao)
	
	goodbye := "Goodbye"
	fmt.Println("Posizione in memoria di goodbye:")
	fmt.Println(&goodbye)
	fmt.Println("Valore di goodbye:")
	fmt.Println(goodbye)
	cambiaStringa(goodbye)
	fmt.Println(goodbye)

}

func cambiaStringaPuntatore(str *string) {
	// str è un nuovo puntatore ad una stringa (una copia di puntatoreACiao)
	// *str è il valore puntato da str
	// str è l'indirizzo di memoria di str
	fmt.Println("Posizione in memoria del puntatore str:")
	fmt.Println(&str)

	fmt.Println("Posizione in memoria di str:")
	fmt.Println(str)
	fmt.Println("Valore di str:")
	fmt.Println(*str)

	*str = "Ciao Mondo"

	fmt.Println(str)
	fmt.Println(*str)
}


func cambiaStringa(str string) {
	// str è una nuova stringa (una copia di goodbye)
	fmt.Println("Posizione in memoria di str:")
	fmt.Println(&str)
	fmt.Println("Valore di str:")
	fmt.Println(str)

	str = "Ciao Mondo"

	fmt.Println(str)
}
