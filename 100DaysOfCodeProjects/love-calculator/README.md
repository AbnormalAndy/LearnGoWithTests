#Love Calculator v0.1

Write a program that tests the compatibility between two people.

To work out the love score between two people:

1. Take both people's names and check for the number of times the letters in the word TRUE occurs.

2. Then check for the number of times the letters in the word LOVE occurs.

3. Then combine these numbers to make a 2 digit number.

For Love Scores less than 10 or greater than 90, the message should be:

- "Your score is *x*, you go together like coke and mentos."

For Love Scores between 40 and 50, the message should be:

- "You score is *y*, you are alright together."

Otherwise, the message will just be their score, e.g.:

- "Your socre is *z*."

e.g.

name1 = "Angela Yu"
name2 = "Jack Bauer"

T occurs 0 times

R occurs 1 time

U occurs 2 times

E occurs 2 times

Total = 5

L occurs 1 time

O occurs 0 times

V occurs 0 times

E occurs 2 times

Total = 3

Love Score = 53

Print: "Your score is 53."

Test

Name1		Name2			Score
Brad Pitt	Jennifer Aniston	73
Prince William	Kate Middleton		67
Ashton Kutcher	Mila Kunis		63
Angela Yu	Jack Bauer		53
David Beckham	Victoria Beckham	45
Mario		Princess Peach		43
Kanye West	Kim Kardashian		42
Truth Truth	Love			94
QA		Violin			02
