package main

import (
	"os"
	"bufio"
	"fmt"
	// "regexp"
	"strconv"
	"strings"
)

func convert(str string) (integer int) {
	conv, _ := strconv.Atoi(str)
	integer = conv
	return
}

func main() {
  table, _ := os.Open("../actual.dat")

  defer table.Close()

  scanner := bufio.NewScanner(table)

	var race []int

  for scanner.Scan() {
		score := strings.ReplaceAll(scanner.Text(), " ", "")
	  data := strings.Split(score, ":")
		str := convert(data[1])
		race = append(race, str)
  }

	var wins []int
	var ways []int

	time := race[0]
	dist := race[1]

	for k := 0; k <= time; k++ {
		score := (time - k) * k

		if (score > dist) {
			ways = append(ways, score)
		}
	}

	win := len(ways)
	wins = append(wins, win)
	fmt.Println(wins)
}
