package main

import "fmt"
import "math"
import "sort"

//import "os"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

/*
	Big thanks to user "nicolas_patrois" for pointing me in the right direction :P
	https://forum.codingame.com/t/network-cabling-puzzle-discussion/41/54
*/

func main() {
	var N int
	fmt.Scan(&N)

	var leftX, rightX, topY, bottomY, distX, distY, medianY, nearestY, midY int
	buildingY := make([]int, N)
	statY := make(map[int]int)
	listY := make([]int, N)

	for i := 0; i < N; i++ {
		var X, Y int
		fmt.Scan(&X, &Y)
		if i == 0 || X < leftX {
			leftX = X
		}
		if i == 0 || X > rightX {
			rightX = X
		}
		if i == 0 || Y < topY {
			topY = Y
		}
		if i == 0 || Y > bottomY {
			bottomY = Y
		}
		buildingY[i] = Y
		statY[Y]++
		listY[i] = Y
	}

	sort.Ints(listY)

	if len(listY)%2 == 1 {
		medianY = listY[(len(listY)+1)/2-1]
	} else {
		medianY = (listY[len(listY)/2-1] + listY[len(listY)/2]) / 2
	}

	nearestY = int(math.Pow(2.0, 30.0))

	for k, _ := range statY {
		if int(math.Abs(float64(k-medianY))) < nearestY {
			nearestY = int(math.Abs(float64(k - medianY)))
			midY = k
		}
	}

	distX = rightX - leftX

	for i := 0; i < N; i++ {
		distY += int(math.Abs(float64(buildingY[i] - midY)))
	}

	fmt.Println(distX + distY) // Write answer to stdout
}
