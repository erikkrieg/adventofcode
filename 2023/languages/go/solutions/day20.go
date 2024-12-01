package solutions

import (
	"fmt"
	"math"
	"strings"

	"github.com/erikkrieg/adventofcode/2023/pkg/input"
)

func init() {
	puzzleSolutions[19] = day20Solution
}

func setupDay20() []string {
	data := input.Lines("day-20")
	if useTestInput {
		data = input.Lines("test-20")
	}
	return data
}

func day20Solution() {
	fmt.Println("Day 20")
	data := setupDay20()

	modules := make(Modules)
	for _, d := range data {
		splitA := strings.Split(d, " -> ")
		kind := splitA[0][:1]
		id := splitA[0][1:]
		if kind == "b" {
			id = "broadcaster"
		}
		modules[id] = &Module{
			id:           id,
			kind:         kind,
			destinations: strings.Split(splitA[1], ", "),
			inputs:       []string{},
			inputPulses:  make(map[string]Pulse),
		}
	}

	for id, module := range modules {
		for _, destId := range module.destinations {
			if dest, ok := modules[destId]; ok {
				dest.inputs = append(dest.inputs, id)
			}
		}
	}

	high := 0
	low := 0
	activeMachineAt := 0

Buttons:
	for b := 0; b < math.MaxInt; b++ {
		pulses := []PendingPulse{{signal: 0, input: "button", destination: "broadcaster"}}
		for i := 0; i < len(pulses); i++ {
			if pulses[i].destination == "rx" {
				fmt.Printf("rx sig: %d, button press: %d\n", pulses[i].signal, b+1)
				if pulses[i].signal == 0 {
					return
				}
			}
			if pulses[i].signal == 0 {
				if b < 1000 {
					low++
				}
				if pulses[i].destination == "rx" && activeMachineAt == 0 {
					activeMachineAt = b + 1
					fmt.Println(b + 1)
					break Buttons
				}
			} else if b < 1000 {
				high++
			}
			if dest, ok := modules[pulses[i].destination]; ok {
				next := dest.Pulse(pulses[i])
				if len(next) == 0 {
					continue
				}
				pulses = append(pulses, next...)
			}
		}
	}

	Solution{
		Part1: low * high,
		Part2: activeMachineAt,
	}.Print()
}

type Pulse = int
type Modules = map[string]*Module
type Module struct {
	id           string
	kind         string
	destinations []string
	inputs       []string
	inputPulses  map[string]Pulse
	flip         bool
}
type PendingPulse struct {
	signal      Pulse
	input       string
	destination string
}

func (m *Module) Pulse(p PendingPulse) []PendingPulse {
	var pending []PendingPulse
	nextSignal := p.signal
	switch m.kind {
	case "%":
		if p.signal == 1 {
			nextSignal = -1
		} else if p.signal == 0 {
			m.flip = !m.flip
			if m.flip {
				nextSignal = 1
			} else {
				nextSignal = 0
			}
		}
	case "&":
		m.inputPulses[p.input] = p.signal
		pulses := 0
		for _, p := range m.inputPulses {
			pulses += p
		}
		if pulses == len(m.inputs) {
			nextSignal = 0
		} else {
			nextSignal = 1
		}
	}

	if nextSignal < 0 {
		return nil
	}

	for _, d := range m.destinations {
		pending = append(pending, PendingPulse{
			signal:      nextSignal,
			input:       m.id,
			destination: d,
		})
	}

	return pending
}
