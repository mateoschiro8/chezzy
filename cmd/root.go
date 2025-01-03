package cmd

import (
	"chezzy/engine"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

// ClearScreen limpia la terminal
func ClearScreen() {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default: // Linux o MacOS
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func HandleCMD() {

	ClearScreen()

	if len(os.Args) < 2 {
		fmt.Println(" Bienvenido a chezzy!")
		printHelp()
		return
	}

	// fmt.Println(os.Args[1:])

	command, args := os.Args[1], os.Args[2:]

	board := engine.Board{}
	board.Init()

	switch command {
	case "help":
		printHelp()
	case "m":
		engine.DecodeMove(&board, args[0])
	case "n":
		engine.SaveGame(&board, "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
		board.ShowBoard()
	case "s":
		board.ShowBoard()
	case "l":
		engine.SaveGame(&board, strings.Join(args, " "))
		board.ShowBoard()
	default:
		fmt.Printf(" Comando desconocido: %s\n", command)
		printHelp()
	}
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
