package main

import "fmt"
import "os"
import "bufio"
import "strings"
import "strconv"
import "math"

// Convert between radians and degrees
var z float64 = math.Pi/180

func Rad (d float64) float64 {
    return d*z
}

func Deg (r float64) float64 {
    return r/z
}
//------------------------------------

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    var LONS string
    scanner.Scan()
    fmt.Sscan(scanner.Text(),&LONS)
    
    var LATS string
    scanner.Scan()
    fmt.Sscan(scanner.Text(),&LATS)

    tLON,_ := strconv.ParseFloat(strings.Replace(LONS,",",".",-1),64) //making float64 from input while replacing commas with points
    tLAT,_ := strconv.ParseFloat(strings.Replace(LATS,",",".",-1),64)

    LON := Rad(tLON) 
    LAT := Rad(tLAT)
    
    var N int

    scanner.Scan()
    fmt.Sscan(scanner.Text(),&N)

    DEFIB := make([][]string, N) //slice for defibrillators infos
    
    for i := 0; i < N; i++ {
        scanner.Scan()
        DEFIB[i] = strings.Split(strings.Replace(scanner.Text(),",",".",-1),";") //split info into array
    }

    DIST := 10000000.0
    NDEF := N-1
    
    for i := 0; i < N; i++ { 
        tDLON,_ := strconv.ParseFloat(DEFIB[i][4],64) //retrieve defib position
        tDLAT,_ := strconv.ParseFloat(DEFIB[i][5],64)
        DLON := Rad(tDLON) //convert degrees to radians
        DLAT := Rad(tDLAT)
        x := (DLON - LON)*math.Cos((LAT+DLAT)/2) //provided equation
        y := DLAT - LAT
        d := math.Sqrt(x*x+y*y)*6371
        fmt.Fprintln(os.Stderr,d)
        if d < DIST {
            NDEF = i
            DIST = d
        }
    }
    // fmt.Fprintln(os.Stderr, "Debug messages...")
    
    fmt.Println(DEFIB[NDEF][1])// Write answer to stdout
}