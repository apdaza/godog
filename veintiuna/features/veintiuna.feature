# file veintiuna feature
Feature: The dealer for the game of 21
  In order to play 21
  As a dealer
  I need to be able to count cards


  Scenario Outline: Get hand total
    Given a base "<hand>"
    When the dealer sums the cards
    Then "<total>" is correct

    Examples:
      | hand          | total |
      | 5,7           | 12    |
      | 5,Q           | 15    |
      | Q,Q,A         | 21    |
      | J,Q,K         | 30    |
      | K,Q,A         | 21    |
      | Q,A           | 21    |
      | A,A,A         | 13    |
