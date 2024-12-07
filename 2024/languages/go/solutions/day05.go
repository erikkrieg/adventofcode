package solutions

import (
	"fmt"
	"strings"

	"github.com/erikkrieg/adventofcode/2024/pkg/input"
	"github.com/erikkrieg/adventofcode/2024/pkg/lib"
)

type Day5 struct {
	data []string
}

type PageOrderingRule struct {
	// The number of the page that the rule applies to.
	PageNumber int
	// Ordering is only valid if this page can only come after these pages
	InvalidBefore map[int]bool
	// Ordering is only valid if this page can only come before these pages
	InvalidAfter map[int]bool
}

type PageOrderingRules map[int]*PageOrderingRule

func (p *PageOrderingRules) isValidUpdate(pages []int) bool {
	a, b := p.findFirstInvalid(pages)
	return a == -1 && b == -1
}

func (p *PageOrderingRules) findFirstInvalid(pages []int) (int, int) {
	rules := *p
	updates := make(map[int]int, len(pages))
	for i, p := range pages {
		updates[p] = i
	}
	for k, v := range updates {
		pageRules, ok := rules[k]
		if !ok {
			fmt.Errorf("Page %d has no page rules", k)
		}
		for checkPage, checkIndex := range updates {
			if _, found := pageRules.InvalidBefore[checkPage]; v < checkIndex && found {
				return v, checkIndex
			}
			if _, found := pageRules.InvalidAfter[checkPage]; v > checkIndex && found {
				return v, checkIndex
			}
		}
	}
	return -1, -1
}

func (p *PageOrderingRules) GetPageRule(pageNumber int) *PageOrderingRule {
	pageRule, ok := (*p)[pageNumber]
	if !ok {
		pageRule = &PageOrderingRule{
			PageNumber:    pageNumber,
			InvalidBefore: make(map[int]bool),
			InvalidAfter:  make(map[int]bool),
		}
		(*p)[pageNumber] = pageRule
	}
	return pageRule
}

func (d *Day5) Setup() {
	data := input.Lines("day-5")
	if useTestInput {
		data = input.Lines("test-5")
	}
	d.data = data
}

func (d *Day5) Solve() {
	fmt.Println("Day 5")
	d.Setup()
	Solution{
		Part1: d.Part1(),
		Part2: d.Part2(),
	}.Print()
}

func (d *Day5) Part1() int {
	rules := make(PageOrderingRules)
	doneParsingRules := false
	midPageNumSum := 0
	for _, line := range d.data {
		if len(line) == 0 {
			doneParsingRules = true
			continue
		}
		if !doneParsingRules {
			pages := lib.Ints(strings.Split(line, "|"))
			rules.GetPageRule(pages[0]).InvalidAfter[pages[1]] = true
			rules.GetPageRule(pages[1]).InvalidBefore[pages[0]] = true
		} else {
			pages := lib.Ints(strings.Split(line, ","))
			if rules.isValidUpdate(pages) {
				midPageNumSum += pages[len(pages)/2]
			}
		}
	}
	return midPageNumSum
}

func (d *Day5) Part2() int {
	rules := make(PageOrderingRules)
	doneParsingRules := false
	midPageNumSum := 0
	for _, line := range d.data {
		if len(line) == 0 {
			doneParsingRules = true
			continue
		}
		if !doneParsingRules {
			pages := lib.Ints(strings.Split(line, "|"))
			rules.GetPageRule(pages[0]).InvalidAfter[pages[1]] = true
			rules.GetPageRule(pages[1]).InvalidBefore[pages[0]] = true
		} else {
			pages := lib.Ints(strings.Split(line, ","))
			invalidA, invalidB := rules.findFirstInvalid(pages)
			if invalidA == -1 {
				continue
			}
			for invalidA > -1 {
				a := pages[invalidA]
				b := pages[invalidB]
				pages[invalidA] = b
				pages[invalidB] = a
				invalidA, invalidB = rules.findFirstInvalid(pages)
			}
			midPageNumSum += pages[len(pages)/2]
		}
	}
	return midPageNumSum
}

func init() {
	puzzleSolutions[4] = (&Day5{}).Solve
}
