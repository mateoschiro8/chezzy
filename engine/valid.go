package engine

func ValidateMove(board *Board, player bool, from, to uint8, piece string) (Move, bool) {

	if board.colorPcs[player].BitSet(to) {

	}

	return NewMove(from, to, 0, 0), true

}
