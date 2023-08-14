package main

func main() {
	i := 3

	// Go não tem aritmética de ponteiros
	var p *int = nil

	p = &i // pegando o endereço de memória da variável i e colocando no ponteiro p

	*p++
	i++

	teste(p)

	println("pegar o valor do ponteiro", *p)
	println("pegar o valor da origem", i)

	println("pegar o endereço que o ponteiro está apontando", p)
	println("pegar o endereço de i", &i)
}

func teste(value *int) {
	*value++
}
