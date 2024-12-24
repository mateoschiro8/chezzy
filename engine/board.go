package engine

import (
	"fmt"
)

type Board struct {
	// Pawn, Rook, Knight, Bishop, Queen, King
	wP, wR, wK, wB, wQ, wKn Bitboard
	bP, bR, bK, bB, bQ, bKn Bitboard
}

func (board *Board) LoadAndShow() {
	fmt.Println("ESTA HAY QUE HACERLA PARA QUE CARGUE LA PARTIDA Y LA MUESTRE")
}

func (board *Board) ShowBoard() {

	boardAsString := "\n"
	for i := 8; i > 0; i-- {
		boardAsString += fmt.Sprintf("%v | ", i)
		for j := 0; j < 8; j++ {
			boardAsString += board.pieceAt(8*(i-1)+j) + " "
		}
		boardAsString += "\n"
	}
	boardAsString += "   ----------------"
	boardAsString += "\n    a b c d e f g h"
	boardAsString += "\n \n"
	fmt.Print(boardAsString)
}

func (board *Board) pieceAt(pos int) string {

	// White
	if board.wP.BitSet(uint8(pos)) {
		return "P"
	}
	if board.wR.BitSet(uint8(pos)) {
		return "R"
	}
	if board.wK.BitSet(uint8(pos)) {
		return "K"
	}
	if board.wB.BitSet(uint8(pos)) {
		return "B"
	}
	if board.wQ.BitSet(uint8(pos)) {
		return "Q"
	}
	if board.wKn.BitSet(uint8(pos)) {
		return "X"
	}

	// Black
	if board.bP.BitSet(uint8(pos)) {
		return "p"
	}
	if board.bR.BitSet(uint8(pos)) {
		return "r"
	}
	if board.bK.BitSet(uint8(pos)) {
		return "k"
	}
	if board.bB.BitSet(uint8(pos)) {
		return "b"
	}
	if board.bQ.BitSet(uint8(pos)) {
		return "q"
	}
	if board.bKn.BitSet(uint8(pos)) {
		return "x"
	}

	return "."
}

func (board *Board) Init() {
	board.wP = 0b0000000011111111000000000000000000000000000000000000000000000000
	board.wR = 0b1000000100000000000000000000000000000000000000000000000000000000
	board.wK = 0b0100001000000000000000000000000000000000000000000000000000000000
	board.wB = 0b0010010000000000000000000000000000000000000000000000000000000000
	board.wQ = 0b0001000000000000000000000000000000000000000000000000000000000000
	board.wKn = 0b0000100000000000000000000000000000000000000000000000000000000000

	board.bP = 0b0000000000000000000000000000000000000000000000001111111100000000
	board.bR = 0b0000000000000000000000000000000000000000000000000000000010000001
	board.bK = 0b0000000000000000000000000000000000000000000000000000000001000010
	board.bB = 0b0000000000000000000000000000000000000000000000000000000000100100
	board.bQ = 0b0000000000000000000000000000000000000000000000000000000000010000
	board.bKn = 0b0000000000000000000000000000000000000000000000000000000000001000

	InitBitboards()
}
