package engine

// El tablero, junta toda la información de la partida y se interactúa con él

import (
	"fmt"
)

type Board struct {
	// Pawn, Rook, kNight, Bishop, Queen, King
	wP, wR, wN, wB, wQ, wK Bitboard
	bP, bR, bN, bB, bQ, bK Bitboard

	wPieces, bPieces Bitboard
}

func (board *Board) ShowBoard() {

	LoadGame(board)

	boardAsString := "\n"
	for i := 8; i > 0; i-- {
		boardAsString += fmt.Sprintf(" %v | ", i)
		for j := 0; j < 8; j++ {
			boardAsString += board.pieceAt(uint8(8*(i-1)+j)) + " "
		}
		boardAsString += "\n"
	}
	boardAsString += "    ----------------"
	boardAsString += "\n     a b c d e f g h"
	boardAsString += "\n \n"
	fmt.Print(boardAsString)
}

func (board *Board) pieceAt(pos uint8) string {

	// White
	if board.wP.BitSet(pos) {
		return "P"
	}
	if board.wR.BitSet(pos) {
		return "R"
	}
	if board.wN.BitSet(pos) {
		return "N"
	}
	if board.wB.BitSet(pos) {
		return "B"
	}
	if board.wQ.BitSet(pos) {
		return "Q"
	}
	if board.wK.BitSet(pos) {
		return "K"
	}

	// Black
	if board.bP.BitSet(pos) {
		return "p"
	}
	if board.bR.BitSet(pos) {
		return "r"
	}
	if board.bN.BitSet(pos) {
		return "n"
	}
	if board.bB.BitSet(pos) {
		return "b"
	}
	if board.bQ.BitSet(pos) {
		return "q"
	}
	if board.bK.BitSet(pos) {
		return "k"
	}

	return "."
}

func (board *Board) Init() {

	LoadGame(board)
	// saveGame(board)
}
