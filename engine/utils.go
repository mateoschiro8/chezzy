package engine

import "fmt"

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
