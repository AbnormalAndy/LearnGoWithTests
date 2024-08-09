# Blackjack

- The deck is unlimited in size.
- There are no jokers.
- The Jack/Queen/King all count as 10.
- The Ace can count as 11 or 1.
- Use the following list as the deck of cards:
    - cards = [11, 2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 10, 10]
- The cards in the list have equal probability of being drawn.
- Cards are not removed from the deck as they are drawn.
- The computer is the dealer.

To-Do:
- Deal card to dealer then player.
- Deal another card to dealer then player.
- Dealer's second card is hidden.
- Dealer must be delt another card if hand total is less than 17.
- Player sees current total.
- Player may ask for another card.
- Ace will count as 11 until score is greater than 21 where it will then count as 1.

Win Condition:
- Player score higher than dealer score.
- Player score is 21.

Lose Condition:
- Player score is greater than 21.
- Player score is less than dealer score.


AddCards Function:
- Created a function that, if the sum was greater than 21, would evaluated the hand and, if card was an eleven and sum was greater than 21, would minus ten from the sum.
    - Will have to test this further with the test cases.

DealCard Function:
- Not sure how to test the randomness of this function. Not sure if the randomness should be tested. Instead, evaluated intake of a deck size of 13 (an array) and returning an integer.

EvaluateWinner Function:
- Evaluated player and dealer sums and return true if player wins.
- Added another `and` statement to case three and four.
    - This prevents the player from winning when the player and dealer have a hand greater than 21.


Learning Point:
- Added a conditional expression with another for loop to AddCards function that would minus 10 from the hand if the hand contained an ace and was greater than 21. An obstacle was to only change one ace at a time as needed and this portion of the code did this.
- Moved code that was in main.go file to blackjack.go file under the EvaluateWinner function. It felt like this portion of code needed to be tested thoroughly so wrote a test and moved it to the blackjack.go file.

