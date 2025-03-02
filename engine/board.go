package engine

// El tablero, junta toda la información de la partida y se interactúa con él

import (
	"fmt"
)

const (
	White bool = true
	Black bool = false

	Pawn   string = "P"
	Rook   string = "R"
	Knight string = "N"
	Bishop string = "B"
	Queen  string = "Q"
	King   string = "K"
)

var pieces = [13]string{"P", "R", "N", "B", "Q", "K",
	"p", "r", "n", "b", "q", "k"}

// var whitePieces = [13]string{"P", "R", "N", "B", "Q", "K"}
// var blackPieces = [13]string{"p", "r", "n", "b", "q", "k"}

type Board struct {
	// Pawn, Rook, kNight, Bishop, Queen, King
	pcs      map[string]*Bitboard
	colorPcs map[bool]*Bitboard
}

func (board *Board) MakeMove(move Move, piece string, player bool) {

	from, to := move.From(), move.To()

	board.colorPcs[player].ClearBit(from)
	board.colorPcs[player].SetBit(to)

	board.pcs[piece].ClearBit(from)
	board.pcs[piece].SetBit(to)

}

func (board *Board) ShowBoard() {

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

	for k, v := range board.pcs {
		if v.BitSet(pos) {
			return k
		}
	}
	return "."
}

func (board *Board) Init() {

	board.pcs = make(map[string]*Bitboard)
	board.colorPcs = make(map[bool]*Bitboard)

	for _, v := range pieces {
		board.pcs[v] = new(Bitboard)

	}

	board.colorPcs[White] = new(Bitboard)
	board.colorPcs[Black] = new(Bitboard)

	board.clearPieces()

	LoadGame(board)
}

func (board *Board) clearPieces() {

	for _, v := range pieces {
		*board.pcs[v] = 0
	}
	*board.colorPcs[White] = 0
	*board.colorPcs[Black] = 0
}
