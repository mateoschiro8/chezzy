package engine

import (
	"fmt"
	"strings"
)

func DecodeMove(board *Board, move string, player bool) (Move, bool) {

	if len(move) < 3 {
		move = "P" + move
	}

	move = strings.ToUpper(move)
	fmt.Println(move)
	piece, col, row := move[0], int(move[1]-'A'+1), int(move[2]-'0')

	fmt.Printf("Pieza %s, fila %d, columna %d \n", string(piece), row, col)

	pos := uint8((row-1)*8 + (col - 1))

	// fmt.Println(pos)

	if board.colorPcs[player].BitSet(pos) {
		return NewMove(0, 0, 0, 0), false
	}

	board.ShowBoard()

	return NewMove(0, pos, 0, 0), true
}
