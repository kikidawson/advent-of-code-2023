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

func findWins(data string, cards, wins map[int]int) {
	re := regexp.MustCompile("[0-9]+")
	card := strings.Split(data, ":")
	numb := convert((re.FindAllString(card[0], -1))[0])
	game := strings.Split(card[1], " | ")
	want := re.FindAllString(game[0], -1)
	have := re.FindAllString(game[1], -1)
	var win int

	for i := 0; i < len(have); i++ {
		if slices.Contains(want, have[i]) {
			win += 1
		}
	}

	cards[numb] = 1
	wins[numb] = win
}

func countCards(cards, wins map[int]int) {
	for j := 1; j <= len(wins); j++ {
		for k := (j + 1); k <= (j + wins[j]); k++ {
			cards[k] += cards[j]
		}
	}
}

func sum(cards, wins map[int]int) (total int) {
	total = 0

	for x := 0; x <= len(wins); x++ {
		total += cards[x]
	}

	return
}

func main() {
  table, _ := os.Open("../actual.dat")

  defer table.Close()

  scanner := bufio.NewScanner(table)

	var wins map[int]int
	wins = make(map[int]int)
	var cards map[int]int
	cards = make(map[int]int)

  for scanner.Scan() {
		findWins(scanner.Text(), cards, wins)
  }

	countCards(cards, wins)
  total := sum(cards, wins)
	fmt.Println("Total:", total)
}
