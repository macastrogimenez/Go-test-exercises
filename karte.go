// Recently, Pero has been into robotics, so he decided to make a robot that checks whether a deck of poker cards is complete.

// He’s already done a fair share of work—he wrote a programme that recognizes the suits of the cards.
//  For simplicity’s sake, we can assume that all cards have a suit and a number.

// The suit of the card is one of the characters P, K, H, T, and the number of the card is an integer between
//  and
// . The robot labels each card in the format
//  where
//  is the suit and
//  is the number. If the card’s number consists of one digit, then
// . For example, the card of suit P and number
//  is labelled P09.

// A complete deck has
//  cards in total—for each of the four suits there is exactly one card with a number between
//  and
// .

// The robot has read the labels of all the cards in the deck and combined them into the string
// . Help Pero finish the robot by writing a programme that reads the string made out of card labels
// and outputs how many cards are missing for each suit. If there are two exact same cards in the deck,
// output GRESKA (Croatian for ERROR).

// Input
// The first and only line of input contains the string
//  (
// ), containing all the card labels.
// Output
// If there are two exact same cards in the deck, output “GRESKA”. Otherwise,
// the first and only line of output must consist of 4 space-separated numbers:
// how many cards of the suit P, K, H, T are missing, respectively.

// ACTION PLAN

// create a MasterSuitsArray with four different suitArrays -----------> DONE
// one for P, K, H, T -> from 0 to 13. -----------> DONE
// position 0 is a counter for the number of elements -----------> DONE
// 1 - 13 are the positions to save elements on the array -----------> DONE
// scanner for reading the input -----------> DONE
// saving the input in a string -----------> DONE
// split the input every time it comes across a letter (regex?) -> save to inputSuitsArray -----------> DONE
// split the input every time it comes across a number (regex?) -> save to inputNumbersArray -----------> DONE
// get the corresponding value from suitsArray and numbersArray, -----------> DONE
// look for the corresponding suitArray (P, K, H, T ) depending on the initial letter -----------> DONE
// check if the corresponding index is empty or taken -> if taken return GRESKA -----------> DONE
// if empty add it to the that array -----------> DONE
// when finished iterating over both suitsArray and numbersArray return the index 0 of each suitArray (P, K, H, T ) -----------> DONE

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func TrimEmpty(arr []string) []string {
	var result []string
	for _, v := range arr {
		if v != "" {
			result = append(result, v)
		}
	}
	return result
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var P [14]int
	var K [14]int
	var H [14]int
	var T [14]int
	var Greska bool
	for i := 0; i < 14; i++ {
		P[i] = 0
		K[i] = 0
		H[i] = 0
		T[i] = 0
	}

	if scanner.Scan() {

		// Scan the input and save it into two arrays of suits and numbers
		input := scanner.Text()
		reNumbers := regexp.MustCompile("[A-Z]")
		reSuits := regexp.MustCompile(`\d{2}`)
		numbers := reNumbers.Split(input, -1)
		suits := reSuits.Split(input, -1)

		//trimming the arrays for empty input
		trimmedNumbers := TrimEmpty(numbers)
		trimmedSuits := TrimEmpty(suits)

		// converting the array of strings of numbers to an array of ints
		var nums []int
		for _, v := range trimmedNumbers {
			i, _ := strconv.Atoi(v)
			nums = append(nums, i)
		}

		// iterate over the arrays of suits and numbers, and based on the suit
		// check if index is taken or not
		for i := 0; i < len(nums); i++ {
			suit := trimmedSuits[i]
			number := nums[i]
			switch suit {
			case "P":
				if P[number] == 0 { // if the value of the index for this card in its suit is 0,
					// 	then it has not been previously added
					P[number] = number // we change the value from 0 to its index
					P[0] += 1          // index 0 is the counter of cards in each suit
				} else {
					Greska = true // else, the card has already been added to the deck
				}
			case "K":
				if K[number] == 0 {
					K[number] = number
					K[0] += 1
				} else {
					Greska = true
				}
			case "H":
				if H[number] == 0 {
					H[number] = number
					H[0] += 1
				} else {
					Greska = true
				}
			case "T":
				if T[number] == 0 {
					T[number] = number
					T[0] += 1
				} else {
					Greska = true
				}
			default:
				Greska = true
				panic("some input is not a valid suit")
			}
		}
		cardsMissingInP := strconv.Itoa(13 - P[0]) // 13 = the total number of cards - card counter for every suit
		cardsMissingInK := strconv.Itoa(13 - K[0])
		cardsMissingInH := strconv.Itoa(13 - H[0])
		cardsMissingInT := strconv.Itoa(13 - T[0])
		totalMissingCards := cardsMissingInP + " " + cardsMissingInK + " " + cardsMissingInH + " " + cardsMissingInT
		if Greska != true {
			fmt.Println(totalMissingCards)
		} else {
			fmt.Println("GRESKA")
		}
	}
}

//PASSEEDDDD!!!!!!
