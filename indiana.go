package main

import "fmt"
import "os"
import "bufio"
import "strings"
import "strconv"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

type Room struct { //room with exits
	top       bool
	bottom    bool
	left      bool
	right     bool
	lockedWay string
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// W: number of columns.
	// H: number of rows.
	var W, H int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &W, &H)

	rooms := make([][]int, H)
	roomTypes := make(map[int]Room)

	//big chunk of info about exits for rooms
	roomTypes[1] = Room{bottom: true}
	roomTypes[2] = Room{left: true, right: true}
	roomTypes[3] = Room{bottom: true}
	roomTypes[4] = Room{left: true, bottom: true, lockedWay: "LEFT"}
	roomTypes[5] = Room{right: true, bottom: true, lockedWay: "RIGHT"}
	roomTypes[6] = Room{left: true, right: true, lockedWay: "TOP"}
	roomTypes[7] = Room{bottom: true}
	roomTypes[8] = Room{bottom: true}
	roomTypes[9] = Room{bottom: true}
	roomTypes[10] = Room{left: true, lockedWay: "LEFT"}
	roomTypes[11] = Room{right: true, lockedWay: "RIGHT"}
	roomTypes[12] = Room{bottom: true}
	roomTypes[13] = Room{bottom: true}
	//---------------------------------------

	for i := 0; i < H; i++ {
		rooms[i] = make([]int, W)
	}

	for i := 0; i < H; i++ {
		scanner.Scan()
		LINE := scanner.Text() // represents a line in the grid and contains W integers. Each integer represents one room of a given type.
		roomIds_s := strings.Split(LINE, " ")
		for j := 0; j < W; j++ {
			temp, _ := strconv.ParseInt(roomIds_s[j], 0, 0)
			rooms[i][j] = int(temp)
		}
	}
	// EX: the coordinate along the X axis of the exit (not useful for this first mission, but must be read).
	var EX int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &EX)

	xp := -1

	for {
		var xi, yi int
		var POS string
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &xi, &yi, &POS)

		roomType := rooms[yi][xi]

		switch true {
		case roomTypes[roomType].bottom && roomTypes[rooms[yi+1][xi]].lockedWay != "TOP" && ((roomType != 4 && roomType != 5) || POS != "TOP"):
			fmt.Println(xi, yi+1)
		case roomTypes[roomType].left && roomTypes[rooms[yi][xi-1]].lockedWay != "RIGHT" && xp != xi-1:
			fmt.Println(xi-1, yi)
		case roomTypes[roomType].right && roomTypes[rooms[yi][xi+1]].lockedWay != "LEFT" && xp != xi+1:
			fmt.Println(xi+1, yi)
		}

		xp = xi

		//fmt.Fprintln(os.Stderr, roomTypes[roomType], POS, xi, yi) // debug stuff

		// fmt.Fprintln(os.Stderr, "Debug messages...")

		//fmt.Println("0 0") // One line containing the X Y coordinates of the room in which you believe Indy will be on the next turn.
	}
}
