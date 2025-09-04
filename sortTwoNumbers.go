// In this problem, your program should read two whole numbers (also called integers) from the input,
//  and print them out in increasing order.

// As a refresher, here are some ways to read two numbers from standard input in a few different languages:

// # Python 3
// line = input()
// a, b = line.split()
// a = int(a)
// b = int(b)

// // C++
// // make sure to first "#include <iostream>"
// int a, b;
// std::cin >> a >> b;

// // Java
// // make sure to first "import java.util.Scanner;"
// Scanner s = new Scanner(System.in);
// int a = s.nextInt(), b = s.nextInt();
// Input
// The input contains one line, which has two integers
//  and
// , separated by a single space. The bounds on these values are
// .

// Output
// Output the smaller number first, and the larger number second.

// Sample Input 1	Sample Output 1
// 3 4
// 3 4
// Sample Input 2	Sample Output 2
// 987 23
// 23 987

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input := scanner.Text()
		array := strings.Split(input, " ")
		finalArray := [2]int{}
		for i := 0; i < len(array); i++ {
			num, _ := strconv.Atoi(array[i])
			finalArray[i] = num
		}
		sort.Ints(finalArray[:])
		num1 := finalArray[0]
		num2 := finalArray[1]
		p := strconv.Itoa(num1) + " " + strconv.Itoa(num2)
		fmt.Println(p)
	}
}
