package pkgconfig

import (
	"github.com/IGLOU-EU/go-wildcard"
)

type Pattern string

func (p Pattern) Match(in string) bool {
	return wildcard.Match(string(p), in)
}

type Patterns []Pattern

func (s Patterns) Match(in string) bool {
	for _, pattern := range s {
		if wildcard.Match(string(pattern), in) {
			return true
		}
	}
	return false
}
