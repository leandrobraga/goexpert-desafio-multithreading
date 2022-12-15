package viacep

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"

	"github.com/leandrobraga/goexpert-desafio-multithreading/validators"
)

type ViaCep struct {
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

func clearCep(cep string) (string, error) {
	if !validators.IsValidateSizeCep(cep) {
		return "", errors.New("tamanho inválido. Cep deve ser n formato 99999999 ou 99.999-999")
	}
	re, _ := regexp.Compile(`[^0-9]`)

	cep = re.ReplaceAllString(cep, "")

	return cep, nil
}

func Get(cep string, ch chan<- ViaCep) {
	var c ViaCep
	parsedCep, err := clearCep(cep)
	if err != nil {
		log.Fatalln(err.Error())
	}
	res, err := http.Get(fmt.Sprintf("http://viacep.com.br/ws/%s/json/", parsedCep))
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err.Error())
	}
	err = json.Unmarshal(body, &c)
	if err != nil {
		log.Fatalln(err.Error())
	}
	ch <- c
}
