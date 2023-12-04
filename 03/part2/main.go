package main

import (
	"os"
	"bufio"
	"fmt"
	"strconv"
	"strings"
	// "math"
	// "bytes"
)

func convert(str string) (integer int) {
	conv, _ := strconv.Atoi(str)
	integer = conv
	return
}

func isValid(in, max int) bool {
	return in >= 0 && in <= max
}

func isNumeric(in byte) bool {
	return in >= 48 && in <= 57
}

func isSpecial(in byte) bool {
	return in == 42
}

func isGear(wth, lth, i, j int, mx []string) (gear bool) {
	gear = false
	var a byte
	var b byte
	var c byte
	var d byte
	var e byte
	var f byte
	var g byte
	var h byte


	if (i == 0) && (j == 0) {
		a = 46
		b = 46
    c = 46
		d = 46
		e = mx[i][j+1]
		f = 46
		g = mx[i+1][j]
		h = mx[i+1][j+1]
	} else if (i == lth) && (j == 0) {
		a = 46
		b = mx[i-1][j]
    c = mx[i-1][j+1]
		d = 46
		e = mx[i][j+1]
		f = 46
		g = 46
		h = 46
	} else if (i == 0) && (j == wth) {
		a = 46
		b = 46
    c = 46
		d = mx[i][j-1]
		e = 46
		f = mx[i+1][j-1]
		g = mx[i+1][j]
		h = 46
	} else if (i == lth) && (j == wth) {
		a = mx[i-1][j-1]
		b = mx[i-1][j]
    c = 46
		d = mx[i][j-1]
		e = 46
		f = 46
		g = 46
		h = 46
	} else if (i == 0) {
		a = 46
		b = 46
    c = 46
		d = mx[i][j-1]
		e = mx[i][j+1]
		f = mx[i+1][j-1]
		g = mx[i+1][j]
		h = mx[i+1][j+1]
	} else if (i == lth) {
		a = mx[i-1][j-1]
		b = mx[i-1][j]
    c = mx[i-1][j+1]
		d = mx[i][j-1]
		e = mx[i][j+1]
		f = 46
		g = 46
		h = 46
	} else if (j == 0) {
		a = 46
		b = mx[i-1][j]
    c = mx[i-1][j+1]
		d = 46
		e = mx[i][j+1]
		f = 46
		g = mx[i+1][j]
		h = mx[i+1][j+1]
	} else if (j == wth) {
		a = mx[i-1][j-1]
		b = mx[i-1][j]
    c = 46
		d = mx[i][j-1]
		e = 46
		f = mx[i+1][j-1]
		g = mx[i+1][j]
		h = 46
	} else {
		a = mx[i-1][j-1]
		b = mx[i-1][j]
    c = mx[i-1][j+1]
		d = mx[i][j-1]
		e = mx[i][j+1]
		f = mx[i+1][j-1]
		g = mx[i+1][j]
		h = mx[i+1][j+1]
	}

  if isNumeric(a) {
		if isNumeric(b) {

		} else {
			
		}
	} 

	return gear
}

func main() {
	data, _ := os.Open("../test.dat")

	defer data.Close()

	scanner := bufio.NewScanner(data)
	var mx []string

	for scanner.Scan() {
    mx = append(mx, scanner.Text())
	}

  // sum := 0
	lth := len(mx)
	wth := len(mx[0])

	for i := 0; i < lth; i++ {
		fmt.Println("---------------------")
		arr := strings.Split(mx[i], "")
		// var value string
		j := 0
		
		for j < wth {
			fmt.Println("i:", i)
			fmt.Println("j:", j)

			if isSpecial(mx[i][j]) {
				fmt.Println("star:", mx[i][j])
				j = j + 1
			} else {
				fmt.Println("no star:", mx[i][j])
				j = j + 1
			}
		}
	}
}
