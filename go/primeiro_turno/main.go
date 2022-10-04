package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var url = "https://resultados.tse.jus.br/oficial/ele2022/544/dados/br/br-c0001-e000544-v.json"

type Result struct {
	Abr []Abr `json:"abr"`
}

type Abr struct {
	Data          string       `json:"dt"`
	Hora          string       `json:"ht"`
	UrnasApuradas string       `json:"pst"`
	Candidatos    []Candidatos `json:"cand"`
}

type Candidatos struct {
	Seq                      string `json:"seq"`
	Numero                   string `json:"n"`
	VotosApurados            string `json:"vap"`
	VotosApuradosPorcentagem string `json:"pvap"`
	Eleito                   string `json:"e"`
}

func main() {
	for true {
		apuracao()
		time.Sleep(60 * time.Second)
	}
}

func apuracao() {
	tseClient := http.Client{
		Timeout: time.Second * 2,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, err := tseClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var result Result
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Fatal(err)
	}

	for _, abr := range result.Abr {
		fmt.Println("")
		fmt.Println("Última atualização:", abr.Data, abr.Hora)
		fmt.Println(abr.UrnasApuradas, "% das seções totalizadas")
		fmt.Println("")
		// candidatos
		for _, candidatos := range abr.Candidatos {
			fmt.Println(candidatos.Numero, candidatos.VotosApuradosPorcentagem, candidatos.VotosApurados)
		}
		fmt.Println("")
		fmt.Println("################################################")
	}
}
