package apicep

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/leandrobraga/goexpert-desafio-multithreading/validators"
)

type APICep struct {
	Code       string `json:"code"`
	State      string `json:"state"`
	City       string `json:"city"`
	District   string `json:"district"`
	Address    string `json:"address"`
	Status     int    `json:"status"`
	Ok         bool   `json:"ok"`
	StatusText string `json:"statusText"`
}

func getValidatedCepToAPICep(cep string) (string, error) {
	if !validators.IsValidateSizeCep(cep) {
		return "", errors.New("Tamanho inválido. Cep deve ser n formato 99999999 ou 99.999-999")
	}
	cep = strings.ReplaceAll(cep, ".", "")
	return cep, nil
}

func Get(cep string, ch chan<- APICep) {
	var c APICep
	parsedCep, err := getValidatedCepToAPICep(cep)
	if err != nil {
		log.Fatalln(err.Error())
	}
	res, err := http.Get(fmt.Sprintf("https://cdn.apicep.com/file/apicep/%s.json", parsedCep))
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
