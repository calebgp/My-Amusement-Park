package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type jogador struct {
	saldo string
}
type Brinquedos struct {
	Brinquedos []Brinquedo `json:"brinquedos"`
}
type Brinquedo struct {
	Nome          string  `json:"nome"`
	Custo         int     `json:"custo"`
	Popularidade  float64 `json:"popularidade"`
	Ingresso      int     `json:"ingresso"`
	Espaco        int     `json:"espaco"`
	QuantidadeMax int     `json:"quantidademax"`
	Code          string  `json:"code"`
}
type Quantidade struct {
	Montanha  int
	bate_bate int
	carrosel  int
	kami_kaze int
	roda      int
}

var fileScanner *bufio.Scanner
var ToysList []Brinquedo

func main() {
	rand.Seed(time.Now().UnixNano())
	jsonFile, err := os.Open("brinquedos.json")
	if err != nil {
		log.Fatalf("Can't %s", err)
	}
	textoFile, _ := ioutil.ReadAll(jsonFile)
	var brinquedos Brinquedos
	json.Unmarshal(textoFile, &brinquedos)
	for e := 1; e < len(brinquedos.Brinquedos); e++ {
		brinquedo := Brinquedo{
			Nome:          brinquedos.Brinquedos[e].Nome,
			Custo:         brinquedos.Brinquedos[e].Custo,
			Popularidade:  brinquedos.Brinquedos[e].Popularidade,
			Ingresso:      brinquedos.Brinquedos[e].Ingresso,
			Espaco:        brinquedos.Brinquedos[e].Espaco,
			QuantidadeMax: brinquedos.Brinquedos[e].QuantidadeMax,
			Code:          brinquedos.Brinquedos[e].Code,
		}
		ToysList = append(ToysList, brinquedo)
	}
	player := jogador{
		saldo: "1000000",
	}
	var (
		Introdução        string
		comando           string
		comecou           bool
		vezes             int
		Opcoes            string
		MenuDecompras     string
		Espaco_disponivel int
	)
	Opcoes = "1 - Menu de Compras \n2 - Vender um brinquedo\n3 - Passar um dia\n4- Passar uma semana"
	MenuDecompras = "O que deseja comprar?\n1- Lotes\n2-Brinquedos\n"
	Introdução = "Bem Vinda(o) ao MyAmusementPark\n Neste jogo você pode criar o seu propio parque de diversões! \n Deseja Começar?\n 1-Sim\n 2-Não"
	fmt.Println(Introdução)
	fmt.Scanf("%s", &comando)

	if comando == "s" || comando == "sim" || comando == "ss" || comando == "Sim" || comando == "1" {
		comecou = true
	} else {
		fmt.Printf("Ok, então volte sempre que quiser\n")
	}
	if comecou {
		Espaco_disponivel = 1000
		saldo_jogador, err := strconv.Atoi(player.saldo)
		if err != nil {
			log.Fatal(err)
		}
		index := 1
		i := 0
		z := 0
		vezes = 0
		myToys := []string{}
		valores := []int{}
		renda_jogador := []int{}
		brinquedo2 := ToysList[0]
		renda_jogador = append(renda_jogador, renda(brinquedo2.Popularidade, brinquedo2.Ingresso))
		myToys = append(myToys, brinquedo2.Nome)
		for comando != "Q" {
			z = 0
			fmt.Printf("\n\nSaldo: %d\n", saldo_jogador)
			fmt.Printf("Brinquedos Atuais: \n")
			for z != len(myToys) {
				fmt.Printf("%s\n", myToys[z])
				z++
			}
			if vezes > 1 {
				fmt.Printf("Oque deseja fazer agora?\n")
			} else {
				fmt.Printf("Oque deseja fazer logo de inicio?\n")
			}
			fmt.Println(Opcoes)
			fmt.Scanf("%s", &comando)
			if comando == "1" || comando == "Menu de compras" {
				fmt.Println(MenuDecompras)
				fmt.Scanf("%s", &comando)
				if comando == "1" {
				}
				if comando == "2" {
				}
				listaBrinquedos()
				fmt.Scanf("%d\n", &index)
				brinquedoComprado := ToysList[index-1]
				if brinquedoComprado.Custo > saldo_jogador {
					fmt.Printf("Você não tem dinheiro o suficiente\n")
				} else {
					if brinquedoComprado.Espaco > Espaco_disponivel {
						fmt.Printf("Você não tem espaço suficiente\n")
					}
					saldo_jogador -= brinquedoComprado.Custo
					fmt.Printf("Parabéns voce comprou o(a) %s", brinquedoComprado.Nome)
					myToys = append(myToys, brinquedoComprado.Nome)
				}
			}
			if comando == "2" || comando == "Vender um brinquedo" {
				for i != len(myToys) {
					fmt.Printf("%d - %s\n Valor: %d", i+1, myToys[i], valores[i])
					i++
				}
				fmt.Scanf("%d", &index)
				myToys = remove(myToys, index-1)
				saldo_jogador += valores[index-1]

			}
			if comando == "3" || comando == "Passar um dia" {
				timer := time.NewTimer(3 * time.Second)
				fmt.Printf("Passando o dia (Aguarde 3 segundos)\n")
				<-timer.C
				if len(renda_jogador) == 1 {

				}
			}
			if comando == "4" || comando == "Passar uma semana" {
				timer2 := time.NewTimer(7 * time.Second)
				fmt.Printf("Passando a semana(Aguarde 7 segundos)\n")
				<-timer2.C
			}
			vezes++
		}
	}
}
func listaBrinquedos() {
	i := 0
	j := 1
	for i != len(ToysList) {
		brinquedo := ToysList[i]
		fmt.Printf("%d- %s\nCusto: %d\nPopularidade: %0.f\n", j, brinquedo.Nome, brinquedo.Custo, math.Round(brinquedo.Popularidade))
		i++
		j++
	}
}

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}
func renda(s float64, r int) int {
	populacao := 0
	s = math.Round(s)
	var y int = int(s)
	populacao = y*rand.Intn(100) + 100
	w := r * populacao
	return w
}
