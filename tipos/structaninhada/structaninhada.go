package main

import "fmt"

type item struct {
	produtoID int
	qtde      int
	preco     float64
}

type pedido struct {
	userID int
	itens  []item
}

func (pedido pedido) valorTotal() float64 {
	total := 0.0
	for _, item := range pedido.itens {
		total += item.preco * float64(item.qtde)
	}
	return total
}

func main() {
	pedido := pedido{
		userID: 1,
		itens: []item{
			item{produtoID: 1, qtde: 2, preco: 12.10},
			item{produtoID: 2, qtde: 1, preco: 23.49},
			item{produtoID: 11, qtde: 100, preco: 3.13},
		},
	}

	fmt.Printf("Valor total do pedido eh %.2f\n", pedido.valorTotal())
}
