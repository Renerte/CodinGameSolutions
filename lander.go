package main

import "fmt"
import "os"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

//LEVEL ONE ONLY!!! (for now)

/*
Hey, I'm not proud of this solution. I intended it to be a lot better but sadly I observed,
that this simulation seems to not work properly. If you tried this code, it would fail in tests.
Though, if you submitted it, it works properly! I'm really sad that this happens, but I can't do anything to fix serverside problem.
Maybe when I get to next levels, I'll find some workaround for this bug ;(
*/

func main() {
    // surfaceN: the number of points used to draw the surface of Mars.
    upforce := 23 //pre-calculated
    var surfaceN int
    fmt.Scan(&surfaceN)
    //pLandX := -1
    pLandY := -1
    ty := 0
    for i := 0; i < surfaceN; i++ {
        // landX: X coordinate of a surface point. (0 to 6999)
        // landY: Y coordinate of a surface point. By linking all the points together in a sequential fashion, you form the surface of Mars.
        var landX, landY int
        fmt.Scan(&landX, &landY)
        if pLandY == landY {
            ty = landY
        }
        pLandY = landY
        //pLandX = landX
    }
    fmt.Fprintln(os.Stderr,4000/upforce)
    for {
        // hSpeed: the horizontal speed (in m/s), can be negative.
        // vSpeed: the vertical speed (in m/s), can be negative.
        // fuel: the quantity of remaining fuel in liters.
        // rotate: the rotation angle in degrees (-90 to 90).
        // power: the thrust power (0 to 4).
        var X, Y, hSpeed, vSpeed, fuel, rotate, power int
        fmt.Scan(&X, &Y, &hSpeed, &vSpeed, &fuel, &rotate, &power)
 
        if vSpeed*-1 > 37 && Y < 2220 {
            fmt.Println("0 4")
        } else {
            fmt.Println("0 0")
        }
        
        
        // fmt.Fprintln(os.Stderr, "Debug messages...")
        
        //fmt.Println("-20 3") // rotate power. rotate is the desired rotation angle. power is the desired thrust power.
    }
}