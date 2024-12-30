package engine

import (
	"fmt"
	"os"
	"strings"
)

// Funciones para guardar y cargar partidas

func saveGame(board *Board) {
	fenString := ""
	emptySpaces := 0
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			k := uint8(8*(7-i) + j)
			piece := board.pieceAt(k)
			if piece == "." {
				emptySpaces++
			} else {
				fenString += fmt.Sprint(emptySpaces) + piece
				emptySpaces = 0
			}
		}
		if emptySpaces != 0 {
			fenString += fmt.Sprint(emptySpaces)
			emptySpaces = 0
		}
		fenString += "/"
	}
	fmt.Println(fenString)
}

func loadGame(board *Board) {
	dat, _ := os.ReadFile("save.fen")
	save := string(dat)

	pcs, info, _ := strings.Cut(save, " ")

	fmt.Println(pcs)
	fmt.Println(info)

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
	saveGame(board)
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
