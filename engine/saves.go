package engine

import (
	"fmt"
	"os"
	"strings"
)

// Funciones para guardar y cargar partidas

func SaveGame(board *Board, fen string) {
	fmt.Println("FEN: ", fen)
	fenString := ""
	emptySpaces := 0
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			k := uint8(8*(7-i) + j)
			piece := board.pieceAt(k)
			if piece == "." {
				emptySpaces++
			} else {
				if emptySpaces != 0 {
					fenString += fmt.Sprint(emptySpaces)
				}
				fenString += piece
				emptySpaces = 0
			}
		}
		if emptySpaces != 0 {
			fenString += fmt.Sprint(emptySpaces)
			emptySpaces = 0
		}
		fenString += "/"
	}
	fenString = fenString[:len(fenString)-1]
	// fmt.Println(fenString)

	if len(fen) > 0 {
		fenString = fen
	}

	err := os.WriteFile("save.fen", []byte(fenString), 0755)
	if err != nil {
		fmt.Printf("unable to write file: %w", err)
	}

}

func LoadGame(board *Board) {
	dat, _ := os.ReadFile("save.fen")
	save := string(dat)

	pcs, info, _ := strings.Cut(save, " ")

	info += "BASATA"
	pieces := strings.Split(pcs, "/")

	for i, v := range pieces {
		file := fillSquares(v)
		for j := 0; j < len(file); j++ {
			k := uint8(8*(7-i) + j)
			piece := file[j]
			switch piece {
			case 'P':
				board.wP.SetBit(k)
			case 'R':
				board.wR.SetBit(k)
			case 'N':
				board.wN.SetBit(k)
			case 'B':
				board.wB.SetBit(k)
			case 'Q':
				board.wQ.SetBit(k)
			case 'K':
				board.wK.SetBit(k)
			case 'p':
				board.bP.SetBit(k)
			case 'r':
				board.bR.SetBit(k)
			case 'n':
				board.bN.SetBit(k)
			case 'b':
				board.bB.SetBit(k)
			case 'q':
				board.bQ.SetBit(k)
			case 'k':
				board.bK.SetBit(k)
			}
		}

	}

	board.wPieces = board.wP | board.wR | board.wN | board.wB | board.wQ | board.wK
	board.bPieces = board.bP | board.bR | board.bN | board.bB | board.bQ | board.bK
}

func fillSquares(file string) string {
	tmp := ""
	for _, v := range file {
		if v >= '1' && v <= '8' {
			for i := 0; i < int(v-'0'); i++ {
				tmp += "."
			}
		} else {
			tmp += string(v)
		}
	}
	return tmp
}
