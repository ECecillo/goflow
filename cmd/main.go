package main

import (
	"fmt"
	"goflow/tui"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(tui.InitialModel())
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Erreur: %v", err)
		os.Exit(1)
	}
}
