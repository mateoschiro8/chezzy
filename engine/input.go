package engine

import (
	"fmt"
	"strings"
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

	// Hay una pieza del otro jugador en la posición destino?
	isCaptureMove := board.colorPcs[!player].BitSet(toPos)

	switch piece {
	case "P":

		if !isCaptureMove {

			colBits := uint64(*board.pcs["P"]) & cols[toCol]

			// fmt.Printf("Binario: %b\n", colBits)
			// fmt.Println(BitToSquare(Bitboard(colBits).Msb()))

			PrintBoardBits(colBits)
			fromPos := Bitboard(colBits).Msb()
			fromRow := fromPos/8 + 1
			fmt.Println("Fila origen:", fromRow)
			fmt.Println("Fila destino:", toRow)

			if !(toRow == 4 && fromRow == 2 || toRow == fromRow+1) {
				return 0, false
			}

			return fromPos, true
		} else {
			fmt.Println("Es captura")
			colBits := uint64(*board.pcs["P"]) & (cols[toCol-1] | cols[toCol+1])
			PrintBoardBits(colBits)
		}

	}

	return 2, true
}

func BitToSquare(bitIndex uint8) string {
	fmt.Println(bitIndex)
	// Validar que el índice esté en el rango correcto
	if bitIndex > 63 {
		return "Índice de bit inválido"
	}

	// Calcular la columna (A-H)
	// La columna se obtiene con el módulo 8
	column := rune('A' + (bitIndex % 8))

	// Calcular la fila (1-8)
	// La fila se obtiene dividiendo por 8 y sumando 1
	// Invertimos la numeración para que el bit 0 sea A1 y el 63 sea H8
	row := 8 - (bitIndex / 8)

	// Convertir a notación de ajedrez
	return fmt.Sprintf("%c%d", column, row)
}

func PrintBoardBits(board uint64) {
	for i := 0; i < 8; i++ {
		// Extraer los 8 bits correspondientes a cada fila
		fila := (board >> (uint(i) * 8)) & 0xFF

		// Imprimir cada bit de la fila de izquierda a derecha (columna A a la izquierda)
		for j := 7; j >= 0; j-- {
			bit := (fila >> uint(j)) & 1
			fmt.Printf("%d", bit)
		}
		fmt.Println() // Salto de línea después de cada fila
	}
}
