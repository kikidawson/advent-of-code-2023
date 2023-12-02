package main

import (
	"os"
	"log"
	"bufio"
	"fmt"
	"strings"
	"regexp"
	"strconv"
)

func convert(str string) (integer int) {
	conv, err := strconv.Atoi(str)

	if err != nil {
		fmt.Println(err)
	}
  
	integer = conv
	return
}

func calculate(data string) (power int) {
	data_slice := strings.Split(data, ":")
	re := regexp.MustCompile("[0-9]+")
	games := strings.Replace(data_slice[1], ";", ",", -1)
	games_split := strings.Split(games, ",")
	red_min := 0
	green_min := 0
	blue_min := 0

	for i := 0; i < len(games_split); i++ {
		if strings.Contains(games_split[i], "red") {
			red_string := re.FindAllString(games_split[i], -1)
			red := convert(red_string[0])
			if red > red_min {
				red_min = red
			}
		} else if strings.Contains(games_split[i], "green") {
			green_string := re.FindAllString(games_split[i], -1)
			green := convert(green_string[0])
			if green > green_min {
				green_min = green
			}
		} else if strings.Contains(games_split[i], "blue") {
			blue_string := re.FindAllString(games_split[i], -1)
			blue := convert(blue_string[0])
			if blue > blue_min {
				blue_min = blue
			}
		}
	}

	power = red_min * green_min * blue_min
	return
}

func sum(values []int) (sum int) {
	sum = 0
	values_length := len(values)

	for i := 0; i < values_length; i++ {
		sum += values[i]
	}

	return
}

func main() {
	data, err := os.Open("../actual.dat")

	if err != nil {
		log.Fatal(err)
	}

	defer data.Close()

	scanner := bufio.NewScanner(data)
	var powers []int

	for scanner.Scan() {
		power := calculate(scanner.Text())
		powers = append(powers, power)
	}

	fmt.Println("sum: ", sum(powers))
}
