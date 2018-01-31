package main

import (
  "fmt"
	"strings"
  "strconv"

)

var Hand string
var Total string

func sum_cards()  {
  cards := strings.Split(Hand, ",")
  value := 0
  // Display all elements.
  for i := range cards {
    value = value + value_card(cards[i])
  }
  // Length is 3.
  fmt.Println(len(cards))
  Total = strconv.Itoa(value)
}

func value_card(c string) int{
  switch c {
  case "A":
    return 1
  case "J":
    return 10
  case "Q":
    return 10
  case "K":
    return 10
  default:
    val, _ := strconv.Atoi(c)
    return val

  }
}
