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

func set_max(gm string, max int, id int) (nid int) {
	re := regexp.MustCompile("[0-9]+")
	str := re.FindAllString(gm, -1)
	clr := convert(str[0])
  nid = id

	if clr > max {
		nid = 0
	}

	return
}

func calculate(data string) (id int) {
	dslice := strings.Split(data, ":")
	re := regexp.MustCompile("[0-9]+")
	id = convert(re.FindAllString(dslice[0], -1)[0])
	gm := strings.Replace(dslice[1], ";", ",", -1)
	draw := strings.Split(gm, ",")

	rmax := 12
	gmax := 13
	bmax := 14

	for i := 0; i < len(draw); i++ {
		if strings.Contains(draw[i], "red") {
			id = set_max(draw[i], rmax, id)
		} else if strings.Contains(draw[i], "green") {
			id = set_max(draw[i], gmax, id)
		} else if strings.Contains(draw[i], "blue") {
			id = set_max(draw[i], bmax, id)
		}
	}

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
	var ids []int

	for scanner.Scan() {
		id := calculate(scanner.Text())
		ids = append(ids, id)
	}

	fmt.Println("sum: ", sum(ids))
}
