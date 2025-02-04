package cmd

import (
	"chezzy/engine"
	"fmt"
	"os"
	"strings"
)

func HandleCMD() {

	engine.ClearScreen()

	if len(os.Args) < 2 {
		fmt.Println(" Bienvenido a chezzy!")
		printHelp()
		return
	}

	// fmt.Println(os.Args[1:])

	command, args := os.Args[1], os.Args[2:]

	board := engine.Board{}
	board.Init()
	engine.LoadGame(&board)

	switch command {
	case "help":
		printHelp()
	case "m":
		move, piece, ok := engine.DecodeMove(&board, args[0], engine.White)
		if ok {
			board.MakeMove(move, piece, engine.White)
		} else {
			fmt.Println("a")
		}
		board.ShowBoard()
	case "n":
		engine.SaveGame(&board, "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
		engine.LoadGame(&board)
		board.ShowBoard()
	case "s":
		board.ShowBoard()
	case "l":
		engine.SaveGame(&board, strings.Join(args, " "))
		engine.LoadGame(&board)
		board.ShowBoard()
	default:
		fmt.Printf(" Comando desconocido: %s\n", command)
		printHelp()
	}

	engine.SaveGame(&board, "")
}

// Muestra la ayuda
func printHelp() {
	fmt.Println(" Comandos disponibles:")
	fmt.Println("   m <movimiento>         - Ejecuta el movimiento indicado, en notación estándar")
	fmt.Println("   n                      - Reinicia la partida")
	fmt.Println("   s                      - Consulta el estado de la partida")
	fmt.Println("   l <game>               - Carga una partida dada en notación FEN")
	fmt.Println("   help                   - Muestra esta ayuda.")
	fmt.Println("")
}
