package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

// model represents the Bubble Tea model for the game
type model struct {
	gameState gameState
	tick      time.Time
	width     int
	height    int
}

type tickMsg struct{}

// initialModel sets up the initial state of the model
func initialModel() model {
	deck := shuffleDeck(newDeck())
	gs := gameState{
		Deck: deck,
	}
	gs.DealInitialCards()
	return model{gameState: gs}
}

// Init is called once when the program starts; it can be used to send initial commands
func (m model) Init() tea.Cmd {
	return nil
}

// Update is called when messages or commands are received; it handles the game logic
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit

		case "h":
			if !m.gameState.GameOver {
				m.gameState.PlayerHit()
				if m.gameState.GameOver {
					m.gameState.DetermineWinner()
				}
			}

		case "s":
			if !m.gameState.GameOver {
				m.gameState.DealerPlay()
				m.gameState.GameOver = true // Ensure the game ends after the dealer plays
				m.gameState.DetermineWinner()
			}

		case "r":
			// Reset the game state
			return initialModel(), nil
		}

	}

	return m, nil
}

// View renders the UI, which is just a string
func (m model) View() string {
	return renderGameState(m.gameState, m.width)
}

// main is the entry point for the application
func main() {
	p := tea.NewProgram(initialModel())

	if err := p.Start(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

func restartGame() {
	execPath, err := os.Executable()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to get executable path: %v\n", err)
		return
	}

	// Start a new instance of the program
	cmd := exec.Command(execPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Start()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to start a new instance: %v\n", err)
		return
	}

	// Exit the current instance
	os.Exit(0)
}
