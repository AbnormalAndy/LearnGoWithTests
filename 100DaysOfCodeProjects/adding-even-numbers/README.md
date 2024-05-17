# Adding Even Numbers

Write a program that calculates the sum of all the even numbers from 1 to X. If X is 100, then the first even number would be 2 and the last one is 100.

i.e. 2 + 4 + 6 + 8 + 10 ... + 98 + 100

Important, there should only be 1 print statement in the console output. It should just print the final total and not every step of the calculation.

Also, constrain the inputs to only take numbers from 0 to a MAX of 1000.

## Example Input 1

10

## Example Output 1

30

## Example Input 2

52

## Example Output 2

702

Learning Point: If taking an integer as user input, can use fmt.Scanln - bufio / os are not required. To catch the error in main.go, needed to use an else statement with the if statement. The switch in addingevennumbers.go did not require an expression. Maybe could have combined the two invalid tests in addingevennumbers_test.go, but believe it is testing two different invalid behaviors so this may be ideal.

