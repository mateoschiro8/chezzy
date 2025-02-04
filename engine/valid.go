package engine

func ValidateMove(board *Board, player bool, from, to uint8, piece string) (Move, string, bool) {

	return NewMove(from, to, 0, 0), piece, true

}
