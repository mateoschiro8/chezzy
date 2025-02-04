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

func DecodeMove(board *Board, move string, player bool) (Move, string, bool) {

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
		return NewMove(0, 0, 0, 0), " ", false
	}

	// fmt.Println(from, to)

	return ValidateMove(board, player, from, to, piece)

}

func reachable(board *Board, piece string, toPos, toCol, toRow uint8, player bool) (uint8, bool) {

	// Si hay una pieza del mismo jugador en esa posición
	if board.colorPcs[player].BitSet(toPos) {
		return 0, false
	}

	switch piece {
	case Pawn:
		return isReachableByPawn(board, player, toPos, toCol, toRow)

	case Rook:
		return isReachableBy(board, Rook, toPos, cols[toCol]|rows[toRow])

	case Knight:
		return isReachableByKnight(board, player, toPos)

	case Bishop:
		return isReachableBy(board, Bishop, toPos, SENWDiags[toPos]|SWNEDiags[toPos])

	case Queen:
		return isReachableBy(board, Queen, toPos, cols[toCol]|rows[toRow]|SENWDiags[toPos]|SWNEDiags[toPos])

	case King:
		return isReachableBy(board, King, toPos, KingMoves[toPos])
	}

	return 0, false
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

func isReachableByKnight(board *Board, player bool, toPos uint8) (uint8, bool) {

	pieces := uint64(*board.pcs[Knight]) & KnightMoves[toPos]

	if pieces == 0 {
		return 0, false
	}

	from := Bitboard(pieces).Msb()
	return from, true
}

func isReachableBy(board *Board, piece string, toPos uint8, attackBits uint64) (uint8, bool) {

	pieces := uint64(*board.pcs[piece]) & attackBits

	if pieces == 0 {
		return 0, false
	}

	occupied := uint64(*board.colorPcs[White] | *board.colorPcs[Black])

	for pieces > 0 {

		from := Bitboard(pieces).Msb()
		pieces &= pieces - 1

		ray := generateRay(from, toPos)
		// PrintBoardBits(ray)
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
			dir = W
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
