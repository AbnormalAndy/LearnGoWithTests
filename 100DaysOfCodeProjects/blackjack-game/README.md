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

To-Do:
- Computer can finish hand immediately.
    - Print results in stages when revealing after player's hand.
    - Will have to count the length of the `computerHand` slice to get an accurate print out.
    - May not be able to print the hand all at once?

- Take player input: "Hit / H" and "Stay / S".
    - If player goes over a sum of 21, should instantly lose.
    - Display the computer's hand after an instant loss?

- Make a separate function to evaluate winner?
    - Compare computer's hand versus player's hand.
    - Would return string.
    - Easier to just do in main.go file?

