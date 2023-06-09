package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/briandowns/spinner"
)

func main() {
	cliHandler(os.Args)
}

// Adicionar Aqui as Tags para Passar na URL
func cliHandler(args []string) {
	if len(args) > 1 {
		switch args[1] {
		case "-h":
			help()
			os.Exit(0)
		case "--help":
			helpExtend()
			os.Exit(0)
		case "-w":
			if len(args) < 4 && len(args[2]) > 0 {
				usage()
				os.Exit(0)
			} else {
				wordlist := args[2]
				url := args[3]
				processWordList(url, wordlist)
			}
		default:
			fmt.Println("Opção inválida.")
		}
	}
}

func usage() {
	println("Modo de Usar: gosearch -w wordlist url")
}

func processWordList(url, wordlist string) {
	// Fazer o GET na URL e verificar se retorna 200
	resp, err := http.Get(url)
	if err != nil {
		println("Erro ao validar URL:")
		usage()
		return
	}

	if resp.StatusCode == http.StatusOK {
		// Fazer o loop pela URL/WordList
		println("URL                                        ", url)
		println("WordList                                   ", wordlist)

		// Abrindo WordList
		wl, err := os.Open(wordlist)
		if err != nil {
			println("Erro: ", err.Error())
			return
		}
		defer wl.Close()

		scanner := bufio.NewScanner(wl)
		println("")
		for scanner.Scan() {
			word := scanner.Text()
			path := url + "/" + word
			// Cria um novo spinner
			s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
			// Inicia o spinner
			s.Start()

			res, err := http.Get(path)
			if err != nil {
				println("Erro: ", err.Error())
				s.Stop() // Para o spinner em caso de erro
				continue
			}
			if res.StatusCode != http.StatusOK {

			} else {
				println(path, "->", res.StatusCode)
			}

			s.Stop() // Para o spinner após a conclusão de uma iteração
		}

		if err := scanner.Err(); err != nil {
			println("Erro ao ler o arquivo: ", err.Error())
		}
	} else {
		fmt.Printf("%s -> %d\n", url, resp.StatusCode)
	}
}

func helpExtend() {
	println("HelpExtend Execution")
}

func help() {
	println("Executando a Funcao Help")
}
