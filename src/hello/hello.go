package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const monitoramentos = 3

func main() {
	exibeIntroducao()

	for {
		exibeMenu()

		fmt.Print("Escolha o comando: ")
		comando := leComandoInt()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs...")
		case 0:
			fmt.Println("Saindo do Programa...")
			os.Exit(0)
		default:
			fmt.Println("Comando não reconhecido")
			os.Exit(-1)
		}

	}

}

func exibeIntroducao() {
	fmt.Println("")
	fmt.Print("Digite seu nome: ")
	nome := leComandoString()
	versao := 1.1
	fmt.Println("Seja bem vindo(a) Sr(a).", nome)
	fmt.Println("Este programa está na versão:", versao)
}

func leComandoInt() int {
	comandoLido := 0
	fmt.Scan(&comandoLido)
	return comandoLido
}

func leComandoString() string {
	var comandoLido string
	fmt.Scan(&comandoLido)
	return comandoLido
}

func exibeMenu() {
	fmt.Println("")
	fmt.Println("======== |Menu| ========")
	fmt.Println("1 - Inicar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
	fmt.Println("========================")
	fmt.Println("")
}

func iniciarMonitoramento() {
	fmt.Println("")
	fmt.Println("Iniciando Monitoramento...")
	sites := []string{"http://www.cebrusa.com.br/palestras/florianopolis-sc/",
		"http://www.cebrusa.com.br/palestras/brasilia-df/",
		"http://www.cebrusa.com.br/palestras/sao-paulo-sp/"}

	for i := 0; i < monitoramentos; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i+1, ":")
			testaSite(site)
		}
	}

}

func testaSite(site string) {
	resp, err := http.Get(site)
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}
	if resp.StatusCode == 200 {
		fmt.Println("Site", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site", site, "está com problema e Status Code:", resp.StatusCode)
	}
}

func lerItensDoArquivo() []string {
	var sites []string
	arquivo, err := ioutil.ReadFile("sites.txt")
	fmt.Println(string(arquivo))
	if err != nil {
		fmt.Println("Erro na abertura do aqruivo:", err)
	}
	fmt.Println(arquivo)
	return sites
}
