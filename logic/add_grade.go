package logic

import (
	"errors"
)

var (
	lowRound    = 2
	highRound   = 7
	lowMult     = 5
	highMult    = 10
	gradeMod    = 10
	gradeCutoff = 37
	minGrade    = 0
	maxGrade    = 100
	Grades      = []int{}
)

func AddGrade(grade int) error {
	if grade < minGrade || grade > maxGrade {
		return errors.New("")
	}

	if grade < gradeCutoff {
		Grades = append(Grades, grade)
		return nil
	}

	mod := grade % gradeMod

	if mod > lowRound && mod < lowMult {
		grade += lowMult - mod
	}

	if mod > highRound {
		grade += highMult - mod
	}

	Grades = append(Grades, grade)
	return nil
}
