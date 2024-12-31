package engine

import "fmt"

// Funciones para decodificar la entrada y obtener el movimiento

func DecodeMove(board *Board, move string) {

	if len(move) < 3 {
		move = "P" + move
	}

	fmt.Println(move)
}
