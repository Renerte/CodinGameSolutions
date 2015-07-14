package main

import "fmt"
import "os"
import "bufio"
import "strings"
//import "strconv"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    // N: Number of elements which make up the association table.
    var N int
    scanner.Scan()
    fmt.Sscan(scanner.Text(),&N)
    
    // Q: Number Q of file names to be analyzed.
    var Q int
    scanner.Scan()
    fmt.Sscan(scanner.Text(),&Q)

    MIME := make(map[string]string) //map of minetypes (key - extension)
    
    for i := 0; i < N; i++ {
        // EXT: file extension
        // MT: MIME type.
        var EXT, MT string
        scanner.Scan()
        fmt.Sscan(scanner.Text(),&EXT, &MT)
        EXT = strings.ToLower(EXT) //lowercase extension
        MIME[EXT] = MT
    }
    for i := 0; i < Q; i++ {
        scanner.Scan()
        FNAME := scanner.Text() // One file name per line.
        FNAME = strings.ToLower(FNAME) //filename to lowercase
        fmt.Fprintln(os.Stderr,FNAME)
        array := strings.Split(FNAME,".")
        if len(array) > 1 { //extension exists?
            mtype, prs := MIME[array[len(array)-1]]
            if prs { //mimetype exists?
                fmt.Println(mtype)
            } else {
                fmt.Println("UNKNOWN")
            }
        } else {
            fmt.Println("UNKNOWN")
        }
    }
    
    // fmt.Fprintln(os.Stderr, "Debug messages...")
    
    //fmt.Println("UNKNOWN") // For each of the Q filenames, display on a line the corresponding MIME type. If there is no corresponding type, then display UNKNOWN.
}