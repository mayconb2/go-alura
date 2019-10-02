package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoramentos = 3
const esperaMonitoramento = 5

func main() {

	exibeInicio()
	lerItensDoArquivo()

	for {
		exibeMenu()

		fmt.Print("Escolha o comando: ")
		comando := leComandoInt()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			imprimeLogs()
		case 0:
			fmt.Println("Saindo do Programa...")
			os.Exit(0)
		default:
			fmt.Println("Comando não reconhecido")
			os.Exit(-1)
		}
	}
}

func exibeInicio() {
	fmt.Println("")
	fmt.Print("Digite seu nome: ")
	nome := leComandoString()
	fmt.Println("Seja bem vindo(a) Sr(a).", nome)
	versao := 1.1
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
	fmt.Println("2 - Exibir Logs dos Sites")
	fmt.Println("3 - Exibir Logs de Login")
	fmt.Println("0 - Sair do Programa")
	fmt.Println("========================")
}

func iniciarMonitoramento() {
	fmt.Println("")
	fmt.Println("Iniciando Monitoramento...")
	sites := lerItensDoArquivo()
	for index := 0; index < monitoramentos; index++ {
		for i, site := range sites {
			fmt.Println("Testando site", i+1, ":")
			testaSite(site)
		}
		time.Sleep(esperaMonitoramento * time.Second)
	}
}

func testaSite(site string) {
	resp, err := http.Get(site)
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}
	if resp.StatusCode == 200 {
		fmt.Println("Site", site, "foi carregado com sucesso!")
		registraLogStatus(site, true)
	} else {
		fmt.Println("Site", site, "está com problema e Status Code:", resp.StatusCode)
		registraLogStatus(site, false)
	}
}

func lerItensDoArquivo() []string {
	var sites []string
	arquivo, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("Erro na abertura do aqruivo:", err)
	}
	leitor := bufio.NewReader(arquivo)
	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)
		if err == io.EOF {
			break
		}
	}
	arquivo.Close()
	return sites
}

func registraLogStatus(site string, status bool) {
	arquivo, err := os.OpenFile("log-status.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Falha na abertura/leitura/ecrita do arquivo:", err)
	} else {
		arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + "-" + site + "online: " + strconv.FormatBool(status) + "\n")
	}
	arquivo.Close()
}

func imprimeLogs() {
	fmt.Println("Exibindo logs...")
	arquivo, err := ioutil.ReadFile("log-status.txt")
	if err != nil {
		fmt.Println("Falha na abertura/leitura do arquivo:", err)
	} else {
		fmt.Println(string(arquivo))
	}
}

func registraLogLogin(nome string) {
	arquivo, err := os.OpenFile("log-login.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Falha na abertura/leitura/ecrita do arquivo:", err)
	} else {
		arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + "- Usuário: " + nome + "\n")
	}
	arquivo.Close()
}
