package main

import (
	"os"
	"bufio"
	"fmt"
	"regexp"
	"strconv"
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

	var races [][]int
	var rarray [][]string
	re := regexp.MustCompile("[0-9]+")

  for scanner.Scan() {
	  score := re.FindAllString(scanner.Text(), -1)
		rarray = append(rarray, score)
  }

	no := len(rarray[0])

	for i := 0; i < no; i++ {
		var race []int
		time := convert(rarray[0][i])
		dist := convert(rarray[1][i])
		race = append(race, time)
		race = append(race, dist)
		races = append(races, race)
	}

	var wins []int

	for j := 0; j < no; j++ {
		var ways []int

		time := races[j][0]
		dist := races[j][1]

		for k := 0; k <= time; k++ {
			score := (time - k) * k

			if (score > dist) {
				ways = append(ways, score)
			}
		}

		win := len(ways)
		wins = append(wins, win)
	}

	total := 1

	for x := 0; x < no; x++ {
		total *= wins[x]
	}

	fmt.Println(total)
}
