package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/leandrobraga/goexpert-desafio-multithreading/apicep"
	"github.com/leandrobraga/goexpert-desafio-multithreading/viacep"
)

func main() {
	var cep string
	flag.StringVar(&cep, "cep", "", "cep para realizar a busca")
	flag.Parse()
	if cep == "" {
		log.Fatalln("Argumento cep obrigatório")
	}

	channelViaCep := make(chan viacep.ViaCep)
	channelAPICep := make(chan apicep.APICep)

	go apicep.Get(cep, channelAPICep)
	go viacep.Get(cep, channelViaCep)
	select {
	case viaCep := <-channelViaCep:
		fmt.Println("VIA CEP")
		fmt.Println(viaCep)
	case apiCep := <-channelAPICep:
		fmt.Println("API CEP")
		fmt.Println(apiCep)
	case <-time.After(1 * time.Second):
		log.Fatalln("requests timeout")
	}

}
