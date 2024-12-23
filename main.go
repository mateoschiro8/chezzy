package main

import (
	"chezzy/engine"
)

func main() {
	// cmd.HandleCMD()
	board := engine.Board{}
	board.Init()
	board.ShowBoard()
}
