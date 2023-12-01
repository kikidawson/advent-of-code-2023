package main

import (
  "fmt"
  "os"
  "log"
  "bufio"
  "regexp"
  "strconv"
)

func calibrate_first(uncalibrate_value string) (calibrated_value string) {
  re := regexp.MustCompile("one|two|three|four|five|six|seven|eight|nine|1|2|3|4|5|6|7|8|9")
  value := re.FindAllString(uncalibrate_value, 1)

	switch value[0] {
	case "one":
		calibrated_value = "1"
	case "two":
		calibrated_value = "2"
	case "three":
		calibrated_value = "3"
	case "four":
		calibrated_value = "4"
	case "five":
		calibrated_value = "5"
	case "six":
		calibrated_value = "6"
	case "seven":
		calibrated_value = "7"
	case "eight":
		calibrated_value = "8"
	case "nine":
		calibrated_value = "9"
	default:
		calibrated_value = value[0]
	}

  return
}

func Reverse(s string) (result string) {
  for _,v := range s {
    result = string(v) + result
  }
  return 
}

func calibrate_last(uncalibrate_value string) (calibrated_value string) {
  reversed_uncalibrate_value := Reverse(uncalibrate_value)

  re := regexp.MustCompile("eno|owt|eerht|ruof|evif|xis|neves|thgie|enin|1|2|3|4|5|6|7|8|9")
  value := re.FindAllString(reversed_uncalibrate_value, 1)

	switch value[0] {
	case "eno":
		calibrated_value = "1"
	case "owt":
		calibrated_value = "2"
	case "eerht":
		calibrated_value = "3"
	case "ruof":
		calibrated_value = "4"
	case "evif":
		calibrated_value = "5"
	case "xis":
		calibrated_value = "6"
	case "neves":
		calibrated_value = "7"
	case "thgie":
		calibrated_value = "8"
	case "enin":
		calibrated_value = "9"
	default:
		calibrated_value = value[0]
	}

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
  calibration_document, err := os.Open("actual.dat")
  var values []int

  if err != nil {
    log.Fatal(err)
  }

  defer calibration_document.Close()

  scanner := bufio.NewScanner(calibration_document)

  for scanner.Scan() {
    first := calibrate_first(scanner.Text())
		last := calibrate_last(scanner.Text())
		value := first + last
    conv, err := strconv.Atoi(value)

		if err != nil {
			fmt.Println(err)
		}

    values = append(values, conv)
  }

  fmt.Println(sum(values))
}
