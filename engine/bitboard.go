package engine

// Contiene la implementación de un tipo de dato bitboard

import (
	"fmt"
	"math/bits"
)

// Un tipo para representar un bitboard, que es un número de 64 bits.
// El bit más significativo (el bit 0) representa la casilla A1,
// y el menos significativo (el bit 63) la casilla H8
type Bitboard uint64

// A constant representing a bitboard with every square set and every square empty.
const FullBB Bitboard = 0xffffffffffffffff
const EmptyBB Bitboard = 0x0

// Set the bit at given square.
func (bitboard *Bitboard) SetBit(sq uint8) {
	*bitboard |= 0x8000000000000000 >> sq
}

// Clear the bit at given square.
func (bitboard *Bitboard) ClearBit(sq uint8) {
	*bitboard &= 0x8000000000000000 >> sq
}

// Test whether the bit of the given bitbord at the given
// position is set.
func (bb Bitboard) BitSet(sq uint8) bool {
	return (bb & (0x8000000000000000 >> sq)) != 0
}

// Get the position of the MSB of the given bitboard.
func (bitboard Bitboard) Msb() uint8 {
	return uint8(bits.LeadingZeros64(uint64(bitboard)))
}

// Get the position of the MSB of the given bitboard,
// and clear the MSB.
func (bitboard *Bitboard) PopBit() uint8 {
	sq := bitboard.Msb()
	bitboard.ClearBit(sq)
	return sq
}

// Count the bits in a given bitboard using the SWAR-popcount
// algorithm for 64-bit integers.
func (bitboard Bitboard) CountBits() int {
	return bits.OnesCount64(uint64(bitboard))
}

// Return a string representation of the given bitboard
func (bitboard Bitboard) String() (bitboardAsString string) {
	bitstring := fmt.Sprintf("%064b\n", bitboard)
	bitboardAsString += "\n"
	for rankStartPos := 56; rankStartPos >= 0; rankStartPos -= 8 {
		bitboardAsString += fmt.Sprintf("%v | ", (rankStartPos/8)+1)
		for index := rankStartPos; index < rankStartPos+8; index++ {
			squareChar := bitstring[index]
			if squareChar == '0' {
				squareChar = '.'
			}
			bitboardAsString += fmt.Sprintf("%c ", squareChar)
		}
		bitboardAsString += "\n"
	}

	bitboardAsString += "   ----------------"
	bitboardAsString += "\n    a b c d e f g h"
	bitboardAsString += "\n"
	return bitboardAsString
}
