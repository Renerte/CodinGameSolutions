package main

import "fmt"
//import "os"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
    for {
        var spaceX, spaceY int
        fmt.Scan(&spaceX, &spaceY)
        
        HM := 0 //highest mountain

        MH := make([]int,8) //slice for mountains

        for i := 0; i < 8; i++ {
            // mountainH: represents the height of one mountain, from 9 to 0. Mountain heights are provided from left to right.
            var mountainH int
            fmt.Scan(&mountainH)
            MH[i] = mountainH
        }
        
        for i := 0; i < 8; i++ { //determine highest
        	if MH[i] > MH[HM] {
        		HM = i
        	}
        }

        if spaceX == HM { //is above highest?
        	fmt.Println("FIRE")
        } else {
        	fmt.Println("HOLD")
        }

        // fmt.Fprintln(os.Stderr, "Debug messages...")
        
        //fmt.Println("HOLD") // either:  FIRE (ship is firing its phase cannons) or HOLD (ship is not firing).
    }
}