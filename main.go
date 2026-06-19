package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/MJ-NMR/gol/core"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	pro := tea.NewProgram(model{})
	if _, err := pro.Run(); err != nil {
		fmt.Printf("there an error : %v", err)
		os.Exit(1)
	}
}

type model struct {
	state   core.State
	courser core.Point
	playing bool
	stop    bool
}

type next struct{}

func (m model) sendNext() tea.Msg {
	if m.stop {
		return nil
	}

	time.Sleep(100 * time.Millisecond)
	return next{}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.state = core.CreateState(msg.Height-3, msg.Width/2)
		m.state.Seed()
		return m, nil

	case next:
		if !m.playing {
			break
		}
		m.state.NextRound()
		return m, m.sendNext

	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k":
			if m.playing {
				return m, nil
			}
			if m.courser.Y > 0 {
				m.courser.Y -= 1
			}

		case "down", "j":
			if m.playing {
				return m, nil
			}
			if m.courser.Y < m.state.Rows-1 {
				m.courser.Y += 1
			}

		case "right", "l":
			if m.playing {
				return m, nil
			}
			if m.courser.X < m.state.Cols-1 {
				m.courser.X += 1
			}

		case "left", "h":
			if m.playing {
				return m, nil
			}
			if m.courser.X > 0 {
				m.courser.X -= 1
			}

		case " ":
			if m.playing {
				return m, nil
			}
			m.state.Toggle(m.courser)

		case "s":
			if m.playing {
				return m, nil
			}
			m.state.NextRound()

		case "n":
			m.state.Seed()

		case "enter":
			if m.playing {
				m.playing = false
				m.stop = true
				return m, nil
			}
			m.playing = true
			m.stop = false
			return m, m.sendNext
		}
	}
	return m, nil
}

func (m model) View() string {
	var b strings.Builder

	for y := range m.state.Rows {
		for x := range m.state.Cols {
			p := core.Point{X: x, Y: y}
			if p == m.courser && !m.playing {
				b.WriteString("\033[7m>\033[0m")
			} else {
				b.WriteByte(' ')
			}

			if m.state.Get(p) {
				b.WriteString("\033[32m󰝤\033[0m")
			} else {
				// b.WriteString("\033[90m \033[0m")
				b.WriteString("\033[2;30m󰝤\033[0m")
			}
		}
		b.WriteByte('\n')
	}

	b.WriteString("\nPress \033[32mq\033[0m: quit, \033[32ms\033[0m: one step, \033[32mSpace\033[0m: toggle cell, \033[32mEnter\033[0m: play/stop, \033[32mn\033[0m: new\n")

	return b.String()
}
