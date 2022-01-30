# Implementação de uma Corrida de Revezamento concorrente em Go

## Autor

* Aroldo Felix Pereira Junior (junioraroldo37@gmail.com)

## Objetivos

* O objetivo deste trabalho é estimular o projeto, implementação e avaliação de soluções para problemas
por meio de programação concorrente, em especial colocando em prática os conceitos e mecanismos
de sincronização de threads.

## Problemática
Simulação de uma corrida de revezamento em que uma equipe possui x corredores. 
* O segundo, terceiro, quarto, ..., e x-ésimo corredores não podem começar a correr até que recebam o bastão entregue pelo corredor que o antecedeu.
* O bastão deve ser passado para o próximo corredor até que o último corredor conclua o trecho a ser percorrido, momento em que o programa deverá ser encerrado.

## Metodologia

* Foram criadas structs para guardar certos dados
* Cada corredor irá correr em uma thread
* O corredor irá passar o bastão em uma decisão randômica (Inicialmente tem 25% de chances de passar o bastão)
  * A cada iteração que ele não entregar o bastão, a chance aumenta em 5%
* Quando o último corredor chegar, será impresso a posição da raia no ranking

## Execução

Foi utilizado o Go 1.17. O programa possui um único método Main.
Este MAIN requisitará 3 informações:
  * Quantos metros tem a corrida (Somente a título de informação)
  * Quantos corredores há por raia
  * Quantas raias há

### Observação
* A saída no console de cada operação de corredor é definida no arquivo `/main/constants.go`
