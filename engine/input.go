package engine

import (
	"fmt"
	"strings"
)

const (
	N  int8 = 8
	S  int8 = -8
	E  int8 = 1
	W  int8 = -1
	NE int8 = 9
	SE int8 = -7
	SW int8 = -9
	NW int8 = 7
)

func DecodeMove(board *Board, move string, player bool) {

	if len(move) < 3 {
		move = "P" + move
	}

	move = strings.ToUpper(move)
	// fmt.Println(move)
	piece, toCol, toRow := string(move[0]), uint8(move[1]-'A'+1), uint8(move[2]-'0')

	// fmt.Printf("Pieza %s, fila %d, columna %d \n", piece, row, col)

	to := (toRow-1)*8 + (toCol - 1)

	from, ok := reachable(board, piece, to, toCol, toRow, player)

	if !ok {
		fmt.Println("Movimiento erróneo")
		return
	}

	// fmt.Println(from, to)

	ValidateMove(board, player, from, to, piece)

}

/*
	Peon: el que esté en la misma columna. Si el movimiento es 2 mas de donde está ahora, el peón tiene que estar
	en la fila inicial

	Caballo: Hacer un and de los bits de los caballos con todos los posibles lugares donde podría salir un caballo
	para llegar a esa posición, y agarrar el que esté más a la izquierda

	Alfil: mismo que con caballo

	Torre: mismo que con caballo

	Reina: mismo que con caballo. Se puede hacer mezclando torre + alfil

	Rey: mismo que con caballo


*/

func reachable(board *Board, piece string, toPos, toCol, toRow uint8, player bool) (uint8, bool) {

	// Si hay una pieza del mismo jugador en esa posición
	if board.colorPcs[player].BitSet(toPos) {
		return 0, false
	}

	switch piece {
	case Pawn:
		return isReachableByPawn(board, player, toPos, toCol, toRow)

	case Rook:
		return isReachableByRook(board, toPos, toCol, toRow)

	case Bishop:
		for i, v := range SWNEDiags {
			fmt.Println(i)
			PrintBoardBits(v)
			fmt.Println()
		}

	}

	return 2, true
}

func isReachableByPawn(board *Board, player bool, toPos, toCol, toRow uint8) (uint8, bool) {

	// Hay una pieza del otro jugador en la posición destino?
	isCaptureMove := board.colorPcs[!player].BitSet(toPos)

	var fromPos uint8
	if !isCaptureMove {

		colBits := uint64(*board.pcs[Pawn]) & cols[toCol]
		fromPos = Bitboard(colBits).Msb()
		fromRow := fromPos/8 + 1

		if !(toRow == 4 && fromRow == 2 || toRow == fromRow+1) {
			return 0, false
		}

	} else {

		colBits := uint64(*board.pcs[Pawn]) & (cols[toCol-1] | cols[toCol+1])
		finalBits := colBits & rows[toRow-1]

		if finalBits == 0 {
			return 0, false
		}
		fromPos = Bitboard(finalBits).Msb()

	}
	return fromPos, true
}

func isReachableByRook(board *Board, toPos, toCol, toRow uint8) (uint8, bool) {

	attackBits := cols[toCol] | rows[toRow]
	rooks := uint64(*board.pcs[Rook]) & attackBits

	if rooks == 0 {
		return 0, false
	}

	occupied := uint64(*board.colorPcs[White] | *board.colorPcs[Black])

	for rooks > 0 {

		from := Bitboard(rooks).Msb()
		rooks &= rooks - 1

		ray := generateRay(from, toPos)
		if ray&occupied == 0 {
			return from, true
		}

	}

	return 0, false
}

// Dadas dos posiciones (que están en alguna dirección), genera una mascara con todas
// las posiciones que tienen en el medio (sin incluirlas)
func generateRay(from, to uint8) uint64 {

	var dir int8

	colFrom, rowFrom := PosToSquare(from)
	colTo, rowTo := PosToSquare(to)

	switch {
	case colFrom == colTo:
		if rowFrom < rowTo {
			dir = N
		} else {
			dir = S
		}

	case rowFrom == rowTo:
		if colFrom < colTo {
			dir = E
		} else {
			dir = S
		}

	case colFrom < colTo:
		if rowFrom < rowTo {
			dir = NE
		} else {
			dir = SE
		}

	case colFrom > colTo:
		if rowFrom < rowTo {
			dir = NW
		} else {
			dir = SW
		}
	}

	ray := uint64(0)
	sq := int8(from) + dir
	for sq != int8(to) {
		ray |= 1 << (63 - sq)
		sq += dir
	}

	return ray
}
