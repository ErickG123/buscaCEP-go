package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	// Fazendo uma requição HTTP através de um CEP digitada
	for _, cep := range os.Args[1:] {
		request, err := http.Get("https://viacep.com.br/ws/" + cep + "/json")
		if err != nil {
			fmt.Fprint(os.Stderr, "Erro ao fazer a requisição: %v \n", err)
		}

		// Fechando a requisição
		defer request.Body.Close()

		// Lendo a resposta da Requisição
		// O Response vem no formato de JSON
		response, err := io.ReadAll(request.Body)
		if err != nil {
			fmt.Fprint(os.Stderr, "Erro ao ler a resposta: %v \n", err)
		}

		// Transformando o Response vinda em JSON em uma Struct
		var data ViaCEP
		err = json.Unmarshal(response, &data)
		if err != nil {
			fmt.Fprint(os.Stderr, "Erro ao fazer o parse da resposta: %v \n", err)
		}

		// Gravando os dados em um Arquivo
		file, err := os.Create("cidade.txt")
		if err != nil {
			fmt.Fprint(os.Stderr, "Erro ao criar o arquivo: %v \n", err)
		}

		// Fechando o arquivo
		defer file.Close()

		_, err = file.WriteString(fmt.Sprintf("CEP: %s, Localidade: %s, UF: %s", data.Cep, data.Localidade, data.Uf))
		fmt.Println("Arquivo criado com sucesso")
		fmt.Println("Cidade:", data.Localidade)
	}
}
