package main

import "fmt"
//import "os"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

 // Seriously, why is this medium difficulty? :P

func main() {
    // nbFloors: number of floors
    // width: width of the area
    // nbRounds: maximum number of rounds
    // exitFloor: floor on which the exit is found
    // exitPos: position of the exit on its floor
    // nbTotalClones: number of generated clones
    // nbAdditionalElevators: ignore (always zero)
    // nbElevators: number of elevators
    var nbFloors, width, nbRounds, exitFloor, exitPos, nbTotalClones, nbAdditionalElevators, nbElevators int
    fmt.Scan(&nbFloors, &width, &nbRounds, &exitFloor, &exitPos, &nbTotalClones, &nbAdditionalElevators, &nbElevators)

    targets := make([]int, nbFloors) //keeps target position for each floor
    
    for i := 0; i < nbElevators; i++ {
        // elevatorFloor: floor on which this elevator is found
        // elevatorPos: position of the elevator on its floor
        var elevatorFloor, elevatorPos int
        fmt.Scan(&elevatorFloor, &elevatorPos)
        targets[elevatorFloor] = elevatorPos
    }

    targets[exitFloor] = exitPos

    for {
        // cloneFloor: floor of the leading clone
        // clonePos: position of the leading clone on its floor
        // direction: direction of the leading clone: LEFT or RIGHT
        var cloneFloor, clonePos int
        var direction string
        fmt.Scan(&cloneFloor, &clonePos, &direction)
        
        if direction == "RIGHT" && clonePos > targets[cloneFloor] || direction == "LEFT" && clonePos < targets[cloneFloor] { //if clone goes wrong way
            fmt.Println("BLOCK")
        } else {
            fmt.Println("WAIT") // action: WAIT or BLOCK
        }
        
        // fmt.Fprintln(os.Stderr, "Debug messages...")
    }
}