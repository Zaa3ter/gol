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
	frame   core.State
	courser location
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

type location struct {
	x int
	y int
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.frame = core.CreateState(uint(msg.Height-3), uint(msg.Width/2))
		m.courser = location{0, 0}
		return m, nil

	case next:
		if !m.playing {
			break
		}
		m.frame = core.PlayRound(m.frame)
		return m, m.sendNext

	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k":
			if m.playing {
				return m, nil
			}
			if m.courser.y > 0 {
				m.courser.y -= 1
			}

		case "down", "j":
			if m.playing {
				return m, nil
			}
			if m.courser.y < len(m.frame)-1 {
				m.courser.y += 1
			}

		case "right", "l":
			if m.playing {
				return m, nil
			}
			if m.courser.x < len(m.frame[0])-1 {
				m.courser.x += 1
			}

		case "left", "h":
			if m.playing {
				return m, nil
			}
			if m.courser.x > 0 {
				m.courser.x -= 1
			}

		case " ":
			if m.playing {
				return m, nil
			}
			m.frame[m.courser.y][m.courser.x] = !m.frame[m.courser.y][m.courser.x]

		case "s":
			if m.playing {
				return m, nil
			}
			m.frame = core.PlayRound(m.frame)
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

	for y, row := range m.frame {
		for x := range row {
			if m.courser.y == y && m.courser.x == x {
				s += "\033[7m>\033[0m"
			} else {
				s += " "
			}

			if m.frame[y][x] {
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
