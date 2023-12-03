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

func convert(str string) (in int) {
	conv, err := strconv.Atoi(str)

	if err != nil {
		fmt.Println(err)
	}
  
	in = conv
	return
}

func set_min(draw string, min int) (nmin int) {
	re := regexp.MustCompile("[0-9]+")
	str := re.FindAllString(draw, -1)
	clr := convert(str[0])
  nmin = min

	if clr > min {
		nmin = clr
	}

	return
}

func calculate(data string) (power int) {
	dslice := strings.Split(data, ":")
	gm := strings.Replace(dslice[1], ";", ",", -1)
	draw := strings.Split(gm, ",")

	rmin := 0
	gmin := 0
	bmin := 0

	for i := 0; i < len(draw); i++ {
		if strings.Contains(draw[i], "red") {
			rmin = set_min(draw[i], rmin)
		} else if strings.Contains(draw[i], "green") {
			gmin = set_min(draw[i], gmin)
		} else if strings.Contains(draw[i], "blue") {
			bmin = set_min(draw[i], bmin)
		}
	}

	power = rmin * gmin * bmin
	return
}

func sum(values []int) (sum int) {
	sum = 0
	vlen := len(values)

	for i := 0; i < vlen; i++ {
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
