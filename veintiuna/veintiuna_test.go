package main

import (
	"github.com/DATA-DOG/godog"
  "fmt"
)


func aBase(arg1 string) error {
  Hand = arg1
	return nil
}

func theDealerSumsTheCards() error {
  sum_cards()
	return nil
}

func isCorrect(arg1 string) error {
  if Total != arg1 {
		return fmt.Errorf("expected %s as total, but there is %s", arg1, Total)
	}
	return nil
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^a base "([^"]*)"$`, aBase)
	s.Step(`^the dealer sums the cards$`, theDealerSumsTheCards)
	s.Step(`^"([^"]*)" is correct$`, isCorrect)
}
