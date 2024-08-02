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
- Easy to sum two cards. Decided to utilize conditional expression to evaluate whether an ace should be converted to a one. A pitfall could be that this function only intakes two cards.
- Perhaps convert to intaking a slice. Iterate through the slice for 11's to convert to 1's if necessary.

DealCard Function:
- Not sure how to test the randomness of this function. Not sure if the randomness should be tested. Instead, evaluated intake of a deck size of 13 (an array) and returning an integer.


To-Do:
- Convert AddCards function to intake a slice.
- May need to adjust conditional statement.
    - Could take a total and, if greater than 21, iterate to find an 11 and make it a 1.
- DrawCard should append to a slice that is the computer and player hands in main.go file.

