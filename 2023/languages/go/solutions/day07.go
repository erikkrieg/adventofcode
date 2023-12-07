package solutions

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/erikkrieg/adventofcode/2023/pkg/input"
	"github.com/erikkrieg/adventofcode/2023/pkg/lib"
)

func init() {
	puzzleSolutions[6] = day7Solution
}

func day7Solution() {
	fmt.Println("Day 7")
	hands := day7Setup()
	Solution{
		Part1: totalWinnings(sortHands(hands, false)),
		Part2: totalWinnings(sortHands(hands, true)),
	}.Print()

}

func day7Setup() []string {
	data := input.Lines("day-7")
	if useTestInput {
		data = input.Lines("test-7")
	}
	return data
}

type Hand struct {
	cards string
	bid   int
	rank  int
}

var order = "AKQJT98765432"

func sortHands(hands []string, wildcards bool) []*Hand {
	sorted := make([]*Hand, len(hands))
	for i, hand := range hands {
		t := strings.TrimSpace(hand)
		s := strings.Split(t, " ")
		bid, err := strconv.Atoi(s[1])
		if err != nil {
			panic(err)
		}
		sorted[i] = &Hand{cards: s[0], bid: bid, rank: -1}
	}
	sort.Slice(sorted, func(i, j int) bool {
		a := sorted[i]
		b := sorted[j]
		ra := rankHand(a, wildcards)
		rb := rankHand(b, wildcards)
		if ra == rb {
			for z := 0; z < 5; z++ {
				if a.cards[z] == b.cards[z] {
					continue
				}
				if wildcards {
					if a.cards[z] == 'J' {
						return false
					}
					if b.cards[z] == 'J' {
						return true
					}
				}
				return strings.IndexRune(order, rune(a.cards[z])) < strings.IndexRune(order, rune(b.cards[z]))
			}
		}
		return ra < rb
	})
	if wildcards {
		for _, s := range sorted {
			fmt.Printf("%+v\n", s)
		}
	}
	return sorted
}

func rankHand(hand *Hand, wildcards bool) int {
	if hand.rank != -1 {
		return hand.rank
	}
	kinds := make(map[string]int)
	for _, c := range hand.cards {
		kinds[string(c)]++
	}
	ranks := [5]int{}
	for label, k := range kinds {
		if wildcards && label == "J" {
			continue
		}
		ranks[5-k]++
	}
	rank := 6
	if ranks[0] == 1 {
		rank = 0
	} else if ranks[1] == 1 {
		rank = 1
		if wildcards && kinds["J"] > 0 {
			rank = 0
		}
	} else if ranks[2] == 1 && ranks[3] == 1 {
		rank = 2
	} else if ranks[2] == 1 {
		rank = 3
		if wildcards && kinds["J"] > 0 {
			rank -= kinds["J"] + 1
		}
	} else if ranks[3] == 2 {
		rank = 4
		if wildcards && kinds["J"] > 0 {
			rank = 2
		}
	} else if ranks[3] == 1 {
		rank = 5
		if wildcards && kinds["J"] > 0 {
			rank = lib.Max(0, 5-(2*kinds["J"]))
		}
	} else if wildcards {
		switch kinds["J"] {
		case 1:
			rank = 5
		case 2:
			rank = 3
		case 3:
			rank = 1
		case 4:
			rank = 0
		case 5:
			rank = 0
		}
	}
	(*hand).rank = rank
	return rank
}

func totalWinnings(hands []*Hand) int {
	total := 0
	for i, hand := range hands {
		total += (len(hands) - i) * hand.bid
	}
	return total
}
