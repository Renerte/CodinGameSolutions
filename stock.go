package main

import "fmt"
import "os"
import "bufio"

//import "strings"
import "strconv"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var n int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &n)

	stocks := make([]int, n)

	scanner.Split(bufio.ScanWords) //make scanner scan per-word

	for i := 0; i < n; i++ {
		scanner.Scan()
		stocks[i], _ = strconv.Atoi(scanner.Text()) //string to int
	}

	fmt.Fprintln(os.Stderr, stocks)

	maxLoss := 0 //largest loss counted
	low := n - 1 //lowest stock value
	loss := 0

	for i := n - 2; i >= 0; i-- {
		if stocks[i] > stocks[low] {
			if stocks[low]-stocks[i] < loss {
				loss = stocks[low] - stocks[i]
			}
		} else if stocks[i] < stocks[low] {
			if maxLoss > loss {
				maxLoss = loss
				loss = 0
			}
			low = i
		}
	}

	if maxLoss > loss {
		maxLoss = loss
		loss = 0
	}

	// fmt.Fprintln(os.Stderr, "Debug messages...")

	fmt.Println(maxLoss) // Write answer to stdout
}
