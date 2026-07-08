package main

import "fmt"

func determinanteOtimizado(mat [][]int) int {
	var ordem int
	var ehQuadrada bool
	var det int

	ehQuadrada, ordem = verificaQuadradaOrdem(mat)
	det = 0
	if ehQuadrada {
		// imprimeMatriz(mat)
		switch ordem {
		case 1:
			// fmt.Println("Ordem 1")
			det = detOrdem1(mat)
		case 2:
			//fmt.Println("Ordem 2")
			det = detOrdem2(mat)
		default:
			//fmt.Println("Ordem ", ordem)
			det = detOrdemNOtimizado(mat)

		}
		//imprimeMatriz(mat)
		//fmt.Println("Det ", det)

	} else {
		fmt.Println("Matriz nao eh quadrada!! retornando 0")
	}
	return det
}

func melhorLinhaColuna(matriz [][]int) (string, int) {

	var linha, coluna, tamanho int
	var zerosLinha, zerosColuna, maisZerosLinha, maisZerosColuna, melhorLinha, melhorColuna, resultado int
	var tipo string

	maisZerosLinha = 0
	maisZerosColuna = 0
	melhorLinha = 0
	melhorColuna = 0

	tamanho = len(matriz)

	for linha = 0; linha < tamanho; linha++ {

		zerosLinha = 0

		for coluna = 0; coluna < tamanho; coluna++ {

			if matriz[linha][coluna] == 0 {

				zerosLinha++

			}

		}

		if zerosLinha > maisZerosLinha {

			maisZerosLinha = zerosLinha
			melhorLinha = linha

		}

	}

	for coluna = 0; coluna < tamanho; coluna++ {

		zerosColuna = 0

		for linha = 0; linha < tamanho; linha++ {

			if matriz[linha][coluna] == 0 {

				zerosColuna++

			}

		}

		if zerosColuna > maisZerosColuna {

			maisZerosColuna = zerosColuna
			melhorColuna = coluna

		}

	}

	resultado = melhorLinha
	tipo = "linha"

	if maisZerosColuna > maisZerosLinha {

		resultado = melhorColuna
		tipo = "coluna"

	}

	return tipo, resultado

}

func detOrdemNOtimizado(mat [][]int) int {

	var sinal, cofator, detTemp, resposta, contL, contC, numL, numC int
	var indiceLinhaOuColuna int
	var tipo string
	var matMenor [][]int

	numL = len(mat)
	numC = len(mat[0])

	resposta = 0

	tipo, indiceLinhaOuColuna = melhorLinhaColuna(mat)

	if tipo == "linha" {

		contL = indiceLinhaOuColuna

		for contC = 0; contC < numC; contC++ {

			cofator = mat[contL][contC]

			if cofator != 0 {

				sinal = calculaSinal(contL, contC)
				matMenor = criaMatriz(numL-1, numC-1)
				copiaMatrizMaiorParaMenor(mat, matMenor, contL, contC)
				detTemp = determinanteOtimizado(matMenor)
				resposta = resposta + (cofator * sinal * detTemp)

			}

		}

	} else {

		contC = indiceLinhaOuColuna

		for contL = 0; contL < numL; contL++ {

			cofator = mat[contL][contC]

			if cofator != 0 {

				sinal = calculaSinal(contL, contC)
				matMenor = criaMatriz(numL-1, numC-1)
				copiaMatrizMaiorParaMenor(mat, matMenor, contL, contC)
				detTemp = determinanteOtimizado(matMenor)
				resposta = resposta + (cofator * sinal * detTemp)

			}

		}

	}

	return resposta
}
