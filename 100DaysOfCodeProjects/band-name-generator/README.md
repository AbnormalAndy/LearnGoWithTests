Band Name Generator

1. Create a greeting for the program.

2. Ask the user for the city that they grew up in.

3. Ask the user for the name of a pet.

4. Combine the name of their city and pet and show them their band name.

5. Make sure the input cursor shows on a new line.


Learning Point: Scan(), Scanln(), and Scanf() were unable to take multiple word inputs. Bufio package was used to take inputs such as New York and not cause the program to skip the second question.

ReadString from the Bufio packages reads until the first occurrence of delim in the input, returning the string containing the data up to and including the delimiter ("\n").

Since "\n" was included in the string inputs collected from the user, TrimSuffix from the Strings package was used to remove these. Not removing these causes the output to be wonky.

