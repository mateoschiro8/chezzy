package engine

import (
	"fmt"
	"strings"
)

func DecodeMove(board *Board, move string, player bool) {

	if len(move) < 3 {
		move = "P" + move
	}

	move = strings.ToUpper(move)
	fmt.Println(move)
	piece, col, row := move[0], int(move[1]-'A'+1), int(move[2]-'0')

	fmt.Printf("Pieza %s, fila %d, columna %d \n", string(piece), row, col)

	var from uint8 = 2
	to := uint8((row-1)*8 + (col - 1))

	// fmt.Println(from, to)

	ValidateMove(board, player, from, to, string(piece))

}
