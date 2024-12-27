package main

import (
	"errors"
	"fmt"
	"time"

	mt "github.com/carvalho2099/meet"
)

const constante string = "constante"

// struct
type Cliente struct {
	Nome     string
	Idade    int
	Endereco Endereco
	Email    string
}

type Endereco struct {
	Rua    string
	Numero int
	Cep    string
	Estado string
}

func (p Cliente) Apresentar() {
	println("Olá", p.Nome)
}

func main() {
	mt.SayHello()
	mt.Say("Hello")

	//variaveis
	var name string = "Golang"
	var testeString = "ola"
	var testeNumero = 123
	var testeString2 string
	var testeNumero2 int
	var boolTeste bool
	texto := "teste"
	fmt.Println(name)
	fmt.Println(testeString)
	fmt.Println(testeNumero)
	fmt.Println(testeString2)
	fmt.Println(testeNumero2)
	fmt.Println(boolTeste)
	fmt.Println(texto)

	//constantes
	fmt.Println(constante)

	//tipos básicos
	var idade int = 30
	var floatNumber float32 = 1.1
	var pi float64 = 3.14
	var raio float64 = 2.5
	var area = pi * raio * raio
	var maior bool = true
	var testeBool bool
	fmt.Println("Idade:", idade)
	fmt.Println("floatNumber", floatNumber)
	fmt.Println("area", area)
	fmt.Println("bool", maior)
	fmt.Println("bool 2 ", testeBool)

	//tipos compostos
	var gavetasArray [2]string
	gavetasArray[0] = "copos"
	gavetasArray[1] = "panos"
	fmt.Println(gavetasArray)

	var gavetasSlice []string
	gavetasSlice = append(gavetasSlice, "copos", "panos")
	fmt.Println(len(gavetasSlice))
	fmt.Println(gavetasSlice)

	var pessoas = map[string]int{}
	pessoas["Lais"] = 26
	pessoas["Leo"] = 32
	fmt.Println(pessoas)
	fmt.Println(pessoas["lais"])

	if idade, ok := pessoas["fabio"]; ok {
		fmt.Println("Pessoa existe no map", idade, ok)
	} else {
		fmt.Println(("Pessoa não existe no map"))
	}

	delete(pessoas, "lais")
	fmt.Println(pessoas)

	//Fluxo de controle
	// if e else
	nota := 75

	if nota >= 90 {
		fmt.Println("Aluno aprovado com distinção!")
	} else if nota >= 70 {
		fmt.Println("Aprovado")
	} else {
		fmt.Println("Aluno reprovado")
	}

	if err := thisIsAnError(); err != nil {
		fmt.Println(err.Error())
	}

	players := map[string]int{
		"lais": 26,
		"leo":  30,
	}

	if value, ok := players["lais"]; ok {
		fmt.Println("pontos:", value, ok)
	}

	//switch case
	fmt.Println("Quando que é sabado?")
	today := time.Now().Weekday()

	switch time.Saturday {
	case today + 0:
		fmt.Println("É hoje")
	case today + 1:
		fmt.Println("É amanhã")
	case today + 2:
		fmt.Println("É depois de amanhã")
	default:
		fmt.Println("Tá longe ainda...")
	}

	//for
	sum := 0
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		sum += i
	}
	fmt.Println(sum)

	//range
	nums := []string{"lais", "leo", "nath"}
	for key, value := range nums {
		fmt.Println(key, value)
	}

	//estruturando seu programa
	//struct
	cliente1 := Cliente{
		Nome:  "Lais",
		Idade: 26,
	}

	fmt.Println(cliente1)

	//funções
	resultado := Soma(2, 2)
	fmt.Println(resultado)

	multi := func(x int) int {
		return x * 2
	}

	fmt.Println(multi(3))

	//métodos
	cliente1 = Cliente{
		Nome: "Lais",
	}
	cliente1.Apresentar()

	//Avançado
	//ponteiros
	//var cliente2 *Cliente = &cliente1
	fmt.Println(&cliente1.Nome)

	//Goroutines
	go exibirMsg()
	time.Sleep(1 * time.Second)
	fmt.Println("Olá main function!")

	//channels
	ch := make(chan int, 3)

	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch)
		fmt.Println("escrita finalizada")
	}()

	time.Sleep(1 * time.Second)
	//valor := <-ch
	//fmt.Println("valor do canal:", valor)
	for valor := range ch {
		fmt.Println(valor)
	}
}

func exibirMsg() {
	fmt.Println("Olá de um goroutine!")
}

func thisIsAnError() error {
	return errors.New("Isto é um erro!")
}

func Soma(a, b int) int {
	return a + b
}
