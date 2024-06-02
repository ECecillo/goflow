package tui

import (
	"fmt"
	"goflow/flowmodor"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	choices        []string
	flowmodorTimer *flowmodor.FlowmodorTimer
	cursorIndex    int
}

func InitialModel() model {
	return model{
		choices:        []string{"Start", "Stop", "Break time"},
		flowmodorTimer: flowmodor.NewFlowmodorTimer(),
		cursorIndex:    0,
	}
}

type tickMsg time.Time

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursorIndex > 0 {
				m.cursorIndex--
			}
		case "down", "j":
			if m.cursorIndex < len(m.choices)-1 {
				m.cursorIndex++
			}
		case "enter":
			switch m.cursorIndex {
			case 0:
				m.flowmodorTimer.Start()
				return m, tea.Tick(time.Second, func(t time.Time) tea.Msg {
					return tickMsg(t)
				})
			case 1:
				m.flowmodorTimer.Stop()
				return m, nil
			case 2:
				m.flowmodorTimer.Break()
			}
		}
	case tickMsg:
		// React only when stop command has no been issued.
		if m.flowmodorTimer.StopTime.IsZero() {
			m.flowmodorTimer.ElapsedTime = time.Since(m.flowmodorTimer.StartTime).Round(time.Second)
			// Emit a new tick each second to update the timer.
			return m, tea.Tick(time.Second, func(t time.Time) tea.Msg {
				return tickMsg(t)
			})
		}
	}
	return m, nil
}

func (m model) View() string {
	var s string
	s = m.flowmodorTimer.FormatElapsedTime()
	if m.flowmodorTimer.BreakTime > 0 {
		s += m.flowmodorTimer.FormatBreakTime()
	}
	if m.flowmodorTimer.ElapsedTime == 0 {
		s += "\n  Appuyez sur Start pour démarrer le chronomètre\n\n"
	}
	for i, choice := range m.choices {
		cursor := " "

		if m.cursorIndex == i {
			s += ">"
		}

		s += fmt.Sprintf("%s %s \n", cursor, choice)
	}

	s += "\n  Appuyez sur Ctrl+C ou q pour quitter.\n"
	return s
}
