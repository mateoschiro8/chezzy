package engine

import (
	"fmt"
	"strings"
)

// Funciones para decodificar la entrada y obtener el movimiento

func DecodeMove(board *Board, move string) {

	if len(move) < 3 {
		move = "P" + move
	}

	move = strings.ToUpper(move)
	fmt.Println(move)
	piece, col, row := move[0], int(move[1]-'A'+1), int(move[2]-'0')

	fmt.Printf("Pieza %s, fila %d, columna %d \n", string(piece), row, col)

	pos := (row-1)*8 + (col - 1)

	// fmt.Println(pos)

	if board.wPieces.BitSet(uint8(pos)) {
		fmt.Println("Movimiento inválido")

	}

	board.ShowBoard()
}
