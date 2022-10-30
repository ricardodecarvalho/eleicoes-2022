package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"text/tabwriter"
	"time"
)

var url = "https://resultados.tse.jus.br/oficial/ele2022/545/dados-simplificados/br/br-c0001-e000545-r.json"

type Result struct {
	Data          string       `json:"dg"`
	Hora          string       `json:"hg"`
	UrnasApuradas string       `json:"pst"`
	Candidatos    []Candidatos `json:"cand"`
}

type Candidatos struct {
	Seq                      string `json:"seq"`
	Numero                   string `json:"n"`
	Nome                     string `json:"nm"`
	VotosApurados            string `json:"vap"`
	VotosApuradosPorcentagem string `json:"pvap"`
	Eleito                   string `json:"e"`
}

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

var (
	Client HTTPClient
)

func init() {
	Client = &http.Client{
		Timeout: time.Second * 2,
	}
}

func main() {
	for {
		err := Apuracao()
		if err != nil {
			log.Fatal(err)
		}

		time.Sleep(60 * time.Second)

		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func Apuracao() error {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	res, err := Client.Do(req)
	if err != nil {
		return err
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var result Result
	err = json.Unmarshal(body, &result)
	if err != nil {
		return err
	}

	fmt.Println("")
	fmt.Println(" Última atualização:", result.Data, result.Hora)
	fmt.Println("", result.UrnasApuradas, "% das seções totalizadas")
	fmt.Println("")

	w := tabwriter.NewWriter(os.Stdout, 1, 1, 5, ' ', tabwriter.AlignRight)
	fmt.Fprintln(
		w,
		"Número\tCandidato\t%\tTotal\t",
	)
	// candidatos
	for _, candidatos := range result.Candidatos {
		fmt.Fprintf(
			w,
			"%s\t%s\t%s\t%s\t\n",
			candidatos.Numero,
			candidatos.Nome,
			candidatos.VotosApuradosPorcentagem,
			candidatos.VotosApurados,
		)
	}
	w.Flush()
	fmt.Println("")

	return nil
}
