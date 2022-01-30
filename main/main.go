package main

import (
	"fmt"
	"sync"
)

func main() {
	var waitGroup sync.WaitGroup
	// Recuperando informações da corrida
	corrida := RecuperarInformacoesCorrida()

	// Gerando as raias
	raias := GerarRaias(corrida)

	// Gerando pódio
	controleLugarPodio := ControleLugarPodio{1}

	// O i-ésimo corredor só poderar correr se Raia.quemDeveCorrer == i
	// Caso contrário, durma

	for i := 0; i < len(raias); i++ {
		raia := raias[i]
		for corredor := 0; corredor < len(raia.corredores); corredor++ {
			waitGroup.Add(1)
			go Correr(CorrerDTO{
				raia:               &raia,
				corredor:           corredor,
				controleLugarPodio: &controleLugarPodio,
				waitGroup:          &waitGroup,
			})
		}
	}

	// Esperando todas as threads finalizarem
	waitGroup.Wait()
	fmt.Println("Corrida finalizada.")
}
