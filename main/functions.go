package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var mutex = &sync.RWMutex{}

func nilInitializer() interface{} {
	return nil
}

/**
 * Recupera inteiro do STDIN se for maior que 0
 */
func recuperarDado() int {
	// enquanto err != nil, tem que pegar o raio novamente
	//stdin := bufio.NewReader(os.Stdin)

	var erroDadoInvalido = nilInitializer()
	var dado = -1
	for {
		_, erroDadoInvalido = fmt.Scanf("%d", &dado)
		//stdin.ReadString('\n')

		if erroDadoInvalido == nil && dado > 0 {
			return dado
		}

		fmt.Println(NEGATIVO)
	}
}

func RecuperarInformacoesCorrida() Corrida {
	fmt.Println("1. Digite quantos metros a corrida tem: ")
	metros := recuperarDado()

	fmt.Println("2. Digite quantos corredores por raia há: ")
	qtdCorredores := recuperarDado()

	fmt.Println("3. Digite quantas raias tem: ")
	qtdRaias := recuperarDado()

	return Corrida{
		tamanho:              metros,
		quantidadeCorredores: qtdCorredores,
		quantidadeRaias:      qtdRaias,
	}
}

func GerarRaias(corrida Corrida) []Raia {
	var raias []Raia
	qtdCorredoresPorRaia := corrida.quantidadeCorredores

	for i := 0; i < corrida.quantidadeRaias; i++ {
		var corredores []int
		for j := 0; j < qtdCorredoresPorRaia; j++ {
			corredores = append(corredores, j)
		}

		raias = append(raias, Raia{
			id:             i,
			quemDeveCorrer: 0,
			corredores:     corredores,
		})
	}

	return raias
}

func raiaConcluiu(raia *Raia) bool {
	return raia.quemDeveCorrer == len(raia.corredores)
}

func Correr(dto CorrerDTO) {
	// Caso eu não deva correr, durmo
	for {
		if dto.raia.quemDeveCorrer == dto.corredor {
			break
		}
		time.Sleep(5)
	}

	// Há uma chance inicial de 20% do corredor chegar em cada iteração
	var chanceChegar = 20

	// Enquanto não cheguei, corro
	if HABILITAR_LOG_CORREDOR {
		fmt.Printf("* Corredor %d da raia %d começou a correr\n", dto.corredor, dto.raia.id)
	}

	for {
		numeroAleatorio := rand.Intn(100)
		// Verificar se o corredor chegou
		chegou := numeroAleatorio <= chanceChegar

		if chegou {
			// Passando o bastão para o próximo corredor
			dto.raia.quemDeveCorrer++
			if HABILITAR_LOG_CORREDOR {
				fmt.Printf("* Corredor %d da raia %d parou de correr\n", dto.corredor, dto.raia.id)
			}

			// Verificando se todos já chegaram
			if raiaConcluiu(dto.raia) {
				avisarConcluiu(dto.raia, dto.controleLugarPodio)
			}
			dto.waitGroup.Done()
			break
		} else {
			// Aumenta em 5% a chance de chegar na próxima rodada
			chanceChegar += 5
		}
	}
}

func avisarConcluiu(raia *Raia, podio *ControleLugarPodio) {
	mutex.Lock()
	fmt.Printf("\t*** raia %d chegou em %dº Lugar\n", raia.id, podio.lugar)
	podio.lugar++
	mutex.Unlock()
}
