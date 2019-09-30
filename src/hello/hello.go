package main

import "fmt"

func main() {

	exibeIntroducao()
	exibeMenu()

	comando := leComando()

	switch comando {
	case 1:
		fmt.Println("Iniciando Monitoramento...")
	case 2:
		fmt.Println("Exibindo Logs...")
	case 0:
		fmt.Println("Saindo do Programa...")
	default:
		fmt.Println("Comando não reconhecido")
	}

}

func exibeIntroducao() {

	nome := "Maycon"
	versao := 1.1

	fmt.Println("Seja bem vindo(a) Sr(a).", nome)
	fmt.Println("Este programa está na versão:", versao)

}

func leComando() int {

	comandoLido := 0
	fmt.Scan(&comandoLido)
	return comandoLido
}

func exibeMenu() {

	fmt.Println("1 - Inicar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")

}
