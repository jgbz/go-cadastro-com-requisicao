package cadastro

import (
	"fmt"
	"net/http"
	"encoding/json"
	"io/ioutil"
)

type Pessoa struct {
	Nome string `json:"nome"`
	Sobrenome string
	Cpf string
	Email string
	Telefone string
	Cep string
	Rua string `json:"logradouro"`
	Numero string
	Bairro string `json:"bairro"`
	Cidade string `json:"localidade"`
}

func Novo() Pessoa{
	var a Pessoa
	fmt.Println("Informe o nome da pessoa");
	fmt.Scan(&a.Nome);
	fmt.Println("Informe o sobrenome da pessoa");
	fmt.Scan(&a.Sobrenome);
	fmt.Println("Informe o cpf da pessoa");
	fmt.Scan(&a.Cpf);
	fmt.Println("Informe o email da pessoa");
	fmt.Scan(&a.Email);
	fmt.Println("Informe o telefone da pessoa");
	fmt.Scan(&a.Telefone);
	fmt.Println("Informe o cep da pessoa");
	fmt.Scan(&a.Cep);
	
	resp, _ := http.Get("https://viacep.com.br/ws/" + a.Cep + "/json/")
	body, _ := ioutil.ReadAll(resp.Body)
	err := json.Unmarshal(body, &a);
	if err != nil{
		fmt.Println("error:", err);
	}	
	return a
}

func (p Pessoa) ToString(){
	fmt.Println("Nome: ", p.Nome);
	fmt.Println("Sobrenome: ", p.Sobrenome);
	fmt.Println("Cpf: ", p.Cpf);
	fmt.Println("Email: ", p.Email);
	fmt.Println("Telefone: ", p.Telefone);
	fmt.Println("Cep: ", p.Cep);
	fmt.Println("Endereco: ", p.Rua, " N: ", p.Numero);
	fmt.Println("Bairro: ", p.Bairro, " Cidade: ", p.Cidade);
}
