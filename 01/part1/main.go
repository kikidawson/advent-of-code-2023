package main

import (
  "fmt"
  "os"
  "log"
  "bufio"
  "regexp"
  "strconv"
  "strings"
)

func calibrate(value string) (calibrated_int int) {
  regexp := regexp.MustCompile("[0-9]+")
  ints := regexp.FindAllString(value, -1)
  ints = strings.Split(strings.Join(ints, ""), "")
  first := ints[0]
  intsLength := len(ints)
  last := ints[intsLength-1]
  calibrated_value := first + last

  conv, err := strconv.Atoi(calibrated_value)

  if err != nil {
    fmt.Println(err)
  }

  calibrated_int = conv

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
    value := calibrate(scanner.Text())
    values = append(values, value)
  }

  sum := 0
  values_length := len(values)

  for i := 0; i < values_length; i++ {
    sum += values[i]
  }

  fmt.Println(sum)
}
