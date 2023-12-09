package main

import (
	"fmt"
	"strings"

	"github.com/floire26/logical-test-sprint/logic"
)

func LogicCaller(times int) {
	for i := 1; i <= times; i++ {
		var grade int
		fmt.Scan(&grade)
		logic.AddGrade(grade)
	}
	fmt.Println(strings.Repeat("-", 20))
	for _, grade := range logic.Grades {
		fmt.Println(grade)
	}
}
