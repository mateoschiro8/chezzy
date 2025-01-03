package engine

// Contiene la implementación de un tipo de dato bitboard

import (
	"math/bits"
)

// Un tipo para representar un bitboard, que es un número de 64 bits.
type Bitboard uint64

// El bit más significativo (el bit 0) representa la casilla A1,
// y el menos significativo (el bit 63) la casilla H8

const FullBB Bitboard = 0xffffffffffffffff
const EmptyBB Bitboard = 0x0

func (bitboard *Bitboard) SetBit(sq uint8) {
	*bitboard |= 0x8000000000000000 >> sq
}

func (bitboard *Bitboard) ClearBit(sq uint8) {
	*bitboard &= 0x8000000000000000 >> sq
}

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
