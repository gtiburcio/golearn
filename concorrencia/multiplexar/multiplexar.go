package main

import (
	"fmt"

	"github.com/tiburcio/html"
)

func encaminhar(origem <-chan string, destino chan string) {
	for {
		destino <- <-origem
	}
}

func juntar(entrada1, entrada2 <-chan string) <-chan string {
	c := make(chan string)
	go encaminhar(entrada1, c)
	go encaminhar(entrada2, c)
	return c
}

func main() {
	c := juntar(
		html.Titulo("https://www.cod3r.com.br", "https://www.google.com.br"),
		html.Titulo("https://www.facebook.com.br", "https://www.youtube.com.br"),
	)

	fmt.Println(<-c)
}
