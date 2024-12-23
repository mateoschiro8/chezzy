package engine

import (
	"fmt"
)

type Board struct {
	// Pawn, Rook, Knight, Bishop, Queen, King
	wP, wR, wK, wB, wQ, wKn Bitboard
	bP, bR, bK, bB, bQ, bKn Bitboard
}

func (board *Board) ShowBoard() {
	fmt.Println(board.wP.String())
}

func (board *Board) Init() {
	board.wP = 0x001001020
}
