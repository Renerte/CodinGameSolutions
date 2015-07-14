package main

import "fmt"
//import "math"
//import "os"
import "sort"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
    var N int
    fmt.Scan(&N)
    
    Pi := make([]int,N)

    for i := 0; i < N; i++ {
        fmt.Scan(&Pi[i])
    }

    sort.Ints(Pi)
    
    res := 10000001

    for i := 0; i < N-1; i++ {
    	//fmt.Fprintln(os.Stderr,"Sprawdzam",i,"-",Pi[i],",",N)
    	if Pi[i+1]-Pi[i] < res {
			//fmt.Fprintln(os.Stderr,int(math.Abs(float64(Pi[i]-Pi[j]))),"jest mniejsze! (",i,Pi[i],j,Pi[j],")")
			res = Pi[i+1]-Pi[i]
		}
    }
    // fmt.Fprintln(os.Stderr, "Debug messages...")
    
    fmt.Println(res)// Write answer to stdout
}