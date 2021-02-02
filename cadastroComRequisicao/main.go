package main

import (
	"os"
	"fmt"
    "os/exec"
	"cadastroComRequisicao/Cadastro"
	"time"
)

func main(){
	menu();
}
 
var listPessoas []cadastro.Pessoa

func menu(){
	var escolhaMenu int
	fmt.Printf("(1) Cadastro de usuario\n(2) Listar usuario\n(3) Sair\n");
	fmt.Scan(&escolhaMenu);
		
	switch (escolhaMenu){
		case 1:
		listPessoas = append(listPessoas, cadastro.Novo())	
		clear();
		case 2:
		carregar();
		for _, p := range listPessoas{
			p.ToString();
			fmt.Println("-----------------------");
		}
		case 3:
		os.Exit(-1);
		
		default:
		fmt.Println("aa");
	}
	menu();
}

func clear(){
	 cmd := exec.Command("clear") //Linux example, its tested
     cmd.Stdout = os.Stdout
     cmd.Run()
}

func carregar(){
	fmt.Print("carregando");
	for i := 0; i<3; i++{
		fmt.Print(".");
		time.Sleep(1000 * time.Millisecond)
	}
	clear();
}
