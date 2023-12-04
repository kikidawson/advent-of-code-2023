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

func isNumeric(s string) bool {
  _, err := strconv.ParseInt(s, 10, 64)
  return err == nil
}

func isSpecial(in byte) bool {
	return ((in >= 33 && in <= 47) || in == 61 || in == 64) && !(in == 46)
}

func isPart(wth, lth, i, j int, mx []string) (part bool) {
	part = false
// 10, 10, 9, 1, matrix
	if (i == 0) {
		if (j == 0) {
			if isSpecial(mx[i][j+1]) || isSpecial(mx[i+1][j]) || isSpecial(mx[i+1][j+1]) {
				part = true
			}
		} else if (j == (wth - 1)) {
			if isSpecial(mx[i][j-1]) || isSpecial(mx[i+1][j-1]) || isSpecial(mx[i+1][j]) {
				part = true
			}
		} else {
			if isSpecial(mx[i][j-1]) || isSpecial(mx[i][j+1]) || isSpecial(mx[i+1][j-1]) || isSpecial(mx[i+1][j]) || isSpecial(mx[i+1][j+1]) {
				part = true
			}
		}
	} else if (i == (lth - 1)) {
		if (j == 0) {
			if isSpecial(mx[i-1][j]) || isSpecial(mx[i-1][j+1]) || isSpecial(mx[i][j+1]) {
				part = true
			}
		} else if (j == (wth - 1)) {
			if isSpecial(mx[i-1][j-1]) || isSpecial(mx[i-1][j]) || isSpecial(mx[i][j-1]) {
				part = true
			}
		} else {
			if isSpecial(mx[i-1][j-1]) || isSpecial(mx[i-1][j]) || isSpecial(mx[i-1][j+1]) || isSpecial(mx[i][j-1]) || isSpecial(mx[i][j+1]) {
				part = true
			}
		}
	} else if (j == 0) {
		if isSpecial(mx[i-1][j]) || isSpecial(mx[i-1][j+1]) || isSpecial(mx[i][j+1]) || isSpecial(mx[i+1][j]) || isSpecial(mx[i+1][j+1]) {
			part = true
		}
	} else if (j == (wth - 1)) {
		if isSpecial(mx[i-1][j-1]) || isSpecial(mx[i-1][j]) || isSpecial(mx[i][j-1]) || isSpecial(mx[i+1][j-1]) || isSpecial(mx[i+1][j]) {
			part = true
		}
	} else {
		if isSpecial(mx[i-1][j-1]) || isSpecial(mx[i-1][j]) || isSpecial(mx[i-1][j+1]) || isSpecial(mx[i][j-1]) || isSpecial(mx[i][j+1]) || isSpecial(mx[i+1][j-1]) || isSpecial(mx[i+1][j]) || isSpecial(mx[i+1][j+1]) {
			part = true
		}
	}

	return
}

func main() {
	data, _ := os.Open("../actual.dat")

	defer data.Close()

	scanner := bufio.NewScanner(data)
	var mx []string

	for scanner.Scan() {
    mx = append(mx, scanner.Text())
	}

  sum := 0
	lth := len(mx)
	wth := len(mx[0])

	for i := 0; i < lth; i++ {
		fmt.Println("---------------------")
		arr := strings.Split(mx[i], "")
		var value string
		j := 0
		
		for j < wth {
			fmt.Println("i:", i)
			fmt.Println("j:", j)
			// i=81
			// j=137
			fmt.Println(arr[j])
			if isNumeric(arr[j]) {
				value = value + arr[j]
				// 2
				// wth-1=139
				if (j == (wth - 1)) {
					if isPart(wth, lth, i, j, mx) {
						sum = sum + convert(value)
						j = j + len(value)
						value = ""
					} else {
						j = j + 1
						value = ""
					}
				}
				fmt.Println("wth:", wth)
				out:
				for k := (j + 1); k < wth ; k++ {
					if isNumeric(arr[k]) {
						value = value + arr[k]
						fmt.Println("k:", k)
						if (k == wth - 1) {
							for x := j; x < (j + len(value)); x++ {
								if isPart(wth, lth, i, x, mx) {
									sum = sum + convert(value)
									j = j + len(value)
									value = ""
									break out
								} else if (x == (j + len(value) - 1)) {
									j = j + len(value)
									value = ""
									break out
								}
							}
						}
					} else {
						for x := j; x < (j + len(value)); x++ {
							if isPart(wth, lth, i, x, mx) {
								sum = sum + convert(value)
								j = j + len(value)
								value = ""
								break out
							} else if (x == (j + len(value) - 1)) {
								j = j + len(value)
								value = ""
								break out
							}
						}
					}
				}

				// if part {
				// 	for k := (j + 1); k < (wth - j); k++ {
				// 		if isNumeric(arr[j + k]) {
				// 			value = value + arr[j + k]
				// 		} else {
				// 			break
				// 		}
				// 	}
				// }

				// fmt.Println(part)

				// mx[i-1][j-1]
				// mx[i-1][j]
				// mx[i-1][j+1] 
				// mx[i][j-1]
				// mx[i][j+1] 
				// mx[i+1][j-1]
				// mx[i+1][j]
				// mx[i+1][j+1]

				// value = append(value, arr[j])
				// for k := (j + 1); k < (wth - j); k++ {
				// 	if isNumeric(arr[k]) {
				// 	  value = append(value, arr[k])
				// 	} else {
				// 		if mx[i-1][j-1] >= 33 && mx[i-1][j-1] <= 47 || mx[i-1][j-1] == 61 || mx[i-1][j-1] == 64 {
				// 			fmt.Println("squash value arr into int and add to part number")
				// 		}
				// 	}
				// 	fmt.Println(value)
				// }
			} else {
				j = j + 1
			}
		}
		fmt.Println("sum:", sum)
	}
}

// if number - store temp
// check next, if number append to temp store etc etc until no number next
// store coords of number
// check surrounding coords for char
// 0-9 = 48-57
//  . = 46


// 467..114..
// ...*......
// ..35..633.
// (i-1)(j-1)
// ......#...
// 617*......
// .....+.58.
// ..592.....
// ......755.
// ...$.*....
// .664.598..