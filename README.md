# Repo for GO exercises

## 28/08/2025 - First steps in Go

Did 3 basic Kattis exercises so far, basically reading input, and returning concatenated input (x3), also basic addition of inputs.

## 05/09/2025 - More Kattis exercises in Go

Solved 3 more Kattis exercises, including a simple addition problem, sorting the provided input and returning it as well as the problem known as Karte (Card deck checker) which was more challenging and interesting.

### Karte
The file karte.go is a Go program that checks whether a deck of poker cards is complete based on a string input of card labels. Each card label consists of a suit (P, K, H, T) and a two-digit number (01–13). The program:

Parses the input string to extract suits and numbers.
Tracks which cards are present for each suit.
Detects duplicate cards (outputs "GRESKA" if any are found).
Otherwise, outputs how many cards are missing for each suit (P, K, H, T).
It is designed to help verify the completeness of a deck and detect errors in card input.

In my solution for the card deck checker in karte.go, I used several key programming concepts:

- Arrays: I used fixed-size arrays to keep track of which cards are present for each suit.
- Loops: I used for loops to initialize the arrays and to iterate over the parsed card data.
- Conditionals: I used if and switch statements to check for duplicate cards, handle different suits, and control the flow of the program.
- Functions: I wrote a helper function (TrimEmpty) to process arrays and remove empty elements.
- String Manipulation & Regular Expressions: I used regular expressions to split the input string into suits and numbers.
- Error Handling: I checked for invalid input and duplicate cards, and I output "GRESKA" when necessary.
- Input/Output: I read input from the user and printed the results to the console.
- Type Conversion: I converted strings to integers for card numbers and back to strings for output.