package main

import "fmt"
//import "os"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
    // W: width of the building.
    // H: height of the building.
    var W, H int
    fmt.Scan(&W, &H)

    // N: maximum number of turns before game over.
    var N int
    fmt.Scan(&N)

    var X0, Y0 int
    fmt.Scan(&X0, &Y0)

    ljh := W
    ljv := H

    for {
        // BOMB_DIR: the direction of the bombs from batman's current location (U, UR, R, DR, D, DL, L or UL)
        hm := 0
        vm := 0
        var BOMB_DIR string
        fmt.Scan(&BOMB_DIR)
        for i := 0; i < len(BOMB_DIR); i++ {
          switch BOMB_DIR[i] {
          case "U"[0]:
            vm = -1
          case "L"[0]:
            hm = -1
          case "D"[0]:
            vm = 1
          case "R"[0]:
            hm = 1
          }
        }
        th := X0 + ljh * hm / 2 //target coords
        tv := Y0 + ljv * vm / 2

        //fmt.Fprintln(os.Stderr, th, tv, W, H, BOMB_DIR, hm, vm, ljh, ljv) //debug data

        ljv = ljv / 2 + 1; //update last jump lengths
        ljh = ljh / 2 + 1;

        if th >= W { //if wrong coords
            th = W - 1
        } else if th < 0 {
            th = 0
        }

        if tv >= H {
            tv = H - 1
        } else if tv < 0 {
            tv = 0
        }

        fmt.Println(th,tv) // the location of the next window Batman should jump to.

        X0 = th //update position
        Y0 = tv
    }
}
//belbel