package main

import (
	"fmt"
	"os"
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
	core.State
	courser core.Point
	playing bool
	stop    bool
}

type next struct{}

func (m model) sendNext() tea.Msg {
	if m.stop {
		return nil
	}

	<-time.After(100 * time.Millisecond)
	return next{}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.State = core.CreateState(msg.Height-3, msg.Width/2)
		return m, nil

	case next:
		if !m.playing {
			break
		}
		m.NextRound()
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
			if m.courser.Y < len(m.State)-1 {
				m.courser.Y += 1
			}

		case "right", "l":
			if m.playing {
				return m, nil
			}
			if m.courser.X < len(m.State[0])-1 {
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
			m.State.Toggle(m.courser)

		case "s":
			if m.playing {
				return m, nil
			}
			m.NextRound()

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

func (m model) View() (s string) {

	for y, row := range m.State {
		for x := range row {
			p := core.Point{X: x, Y: y}
			if p == m.courser {
				s += "\033[7m>\033[0m"
			} else {
				s += " "
			}

			if m.State.Get(p) {
				s += "\033[32m◼\033[0m"
			} else {
				s += "\033[90m░\033[0m"
			}
		}
		s += "\n"
	}

	s += "\nPress \033[32mq\033[0m: quit, \033[32ms\033[0m: one step, \033[32mSpace\033[0m: toggele cell, \033[32mEnter\033[0m: play/stop\n"

	return s
}
