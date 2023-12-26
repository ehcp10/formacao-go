package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Usuarios struct {
	Usuarios []Usuario `json:"usuarios"`
}

type Usuario struct {
	Nome   string `json:"nome"`
	Tipo   string `json:"tipo"`
	Idade  int    `json:"idade"`
	Social Social `json:"social"`
}

type Social struct {
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
}

func main() {
	//path, err := os.Getwd()
	//fullPath := path + "/json/usuarios.json"
	jsonFile, err := os.Open("./json/usuarios.json")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("arquivo aberto com sucesso")
	defer jsonFile.Close()

	var usuarios Usuarios
	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(byteValue, &usuarios)
	if err != nil {
		fmt.Println(err)
	}

	for _, user := range usuarios.Usuarios {
		fmt.Println("Tipo do usuário " + user.Tipo)
		fmt.Println("Nome do usuário " + user.Nome)
	}
}
