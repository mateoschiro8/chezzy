package cmd

import (
	"chezzy/engine"
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

// ClearScreen limpia la terminal
func ClearScreen() {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default: // Linux o MacOS
		fmt.Print("\033[H\033[2J")
	}
}

func HandleCMD() {

	ClearScreen()

	if len(os.Args) < 2 {
		fmt.Println("Bienvenido a chezzy!")
		printHelp()
		return
	}

	// fmt.Println(os.Args[1:])

	command := os.Args[1]
	// args := os.Args[2:]

	board := engine.Board{}

	switch command {
	case "help":
		printHelp()
	case "m":

	case "n":

	case "s":
		board.LoadAndShow()
	default:
		fmt.Printf("Comando desconocido: %s\n", command)
		printHelp()
	}
}

// Muestra la ayuda
func printHelp() {
	fmt.Println("Comandos disponibles:")
	fmt.Println("  m <movimiento>         - Ejecuta el movimiento indicado, en notación estándar")
	fmt.Println("  n {b,w}                - Reinicia la partida, indicando el color con el que se desea jugar")
	fmt.Println("  s                      - Consulta el estado de la partida")
	fmt.Println("  l <game>               - Carga una partida dada en notación FEN")
	fmt.Println("  help                   - Muestra esta ayuda.")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
}
