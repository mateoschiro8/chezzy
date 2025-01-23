package engine

func ValidateMove(board *Board, player bool, from, to uint8, piece string) {

	if board.colorPcs[player].BitSet(to) {

	}

	// NewMove(from, to, 0, 0), true

}
