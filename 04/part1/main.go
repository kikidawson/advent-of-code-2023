package main

import (
  "fmt"
  "os"
  "bufio"
  "regexp"
  "strings"
	"slices"
	"strconv"
)

func convert(str string) (integer int) {
	conv, _ := strconv.Atoi(str)
	integer = conv
	return
}

func addScore(wins int) (score int) {
	if (wins == 0) {
		score = 0
	} else if (wins == 1) {
		score = 1
	} else {
		score = 1
		for j := 1; j < wins; j++ {
			score = score * 2
		}
	}

	return
}

func findWins(data string) (numb, score int) {
	re := regexp.MustCompile("[0-9]+")
	card := strings.Split(data, ":")
	numb = convert((re.FindAllString(card[0], -1))[0])
	game := strings.Split(card[1], " | ")
	want := re.FindAllString(game[0], -1)
	have := re.FindAllString(game[1], -1)
	var win []string

	for i := 0; i < len(have); i++ {
		if slices.Contains(want, have[i]) {
			win = append(win, have[i])
		}
	}

	score = addScore(len(win))

	return
}

func main() {
  table, _ := os.Open("../actual.dat")

  defer table.Close()

  scanner := bufio.NewScanner(table)

	var total int
	var m map[int]int
	m = make(map[int]int)

  for scanner.Scan() {
		numb, score := findWins(scanner.Text())

		m[numb] = score
		total += score
  }

	fmt.Println("total:", total)
}
