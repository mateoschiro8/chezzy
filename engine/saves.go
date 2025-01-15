package engine

import (
	"fmt"
	"os"
	"strings"
)

// Funciones para guardar y cargar partidas

func SaveGame(board *Board, fen string) {
	fenString := ""
	if len(fen) > 0 {
		fenString = fen
	} else {
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
	}
	// fmt.Println("FEN: ", fen)

	err := os.WriteFile("save.fen", []byte(fenString), 0755)
	if err != nil {
		fmt.Printf("unable to write file: %w", err)
	}

}

func LoadGame(board *Board) {
	board.clearPieces()
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
			if piece != '.' {
				board.pcs[string(piece)].SetBit(k)
			}
		}

	}

	*board.colorPcs[White] = *board.pcs["P"] | *board.pcs["R"] | *board.pcs["N"] |
		*board.pcs["B"] | *board.pcs["Q"] | *board.pcs["K"]

	*board.colorPcs[Black] = *board.pcs["p"] | *board.pcs["r"] | *board.pcs["n"] |
		*board.pcs["b"] | *board.pcs["q"] | *board.pcs["k"]
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
