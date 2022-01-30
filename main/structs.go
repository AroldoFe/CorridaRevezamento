package main

import "sync"

/**
 * Struct definindo dados da corrida
 */
type Corrida struct {
	tamanho              int
	quantidadeCorredores int
	quantidadeRaias      int
}

/**
 * Struct definindo dados da raia
 */
type Raia struct {
	// ID da Raia
	id int
	// Quem deve estar correndo no momento
	quemDeveCorrer int
	corredores     []int
}

type ControleLugarPodio struct {
	lugar int
}

/**
 * Struct juntando parâmetros da função Correr
 */
type CorrerDTO struct {
	raia               *Raia
	corredor           int
	controleLugarPodio *ControleLugarPodio
	waitGroup          *sync.WaitGroup
}
