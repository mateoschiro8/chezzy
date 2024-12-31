package engine

// Un movimiento es un número de 32 bits

type Move uint32

// Bits [0, 5]   => Casilla origen (0-63)
// Bits [6, 11]  => Casilla destino (0-63)
// Bits [12, 13] => Tipo de movimiento
// Bits [14, 15] => Información del movimiento
// Bits [16, 31] => Puntaje del movimiento (para búsqueda)

const (
	// Tipos de movimiento
	Quiet     uint8 = 0
	Attack    uint8 = 1
	Castle    uint8 = 2
	Promotion uint8 = 3
)

func NewMove(from, to, moveType, flag uint8) Move {
	return Move(uint32(from)<<26 | uint32(to)<<20 | uint32(moveType)<<18 | uint32(flag)<<16)
}

func (move Move) From() uint8 {
	return uint8((move & 0xfc000000) >> 26)
}

func (move Move) To() uint8 {
	return uint8((move & 0x3f00000) >> 20)
}

func (move Move) Type() uint8 {
	return uint8((move & 0xc0000) >> 18)
}

func (move Move) Flag() uint8 {
	return uint8((move & 0x30000) >> 16)
}

func (move Move) Score() uint16 {
	return uint16(move & 0xffff)
}

func (move Move) MoveString() string {
	return "De casilla " + string(move.From()) + " a casilla " + string(move.To())
}
