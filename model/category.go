package model

import "fmt"

type CategoryType int

const (
	Needs CategoryType = iota
	Wants
)

func (ct CategoryType) String() string {
	switch ct {
	case Needs:
		return "Needs"
	case Wants:
		return "Wants"
	default:
		return "Unknown"
	}
}

func StringToCategory(s string) (CategoryType, error) {
	switch s {
	case "Needs":
		return Needs, nil
	case "Wants":
		return Wants, nil
	default:
		return -1, fmt.Errorf("invalid category: %s", s)
	}
}
