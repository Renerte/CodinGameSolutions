package main

import "fmt"
import "os"
import "bufio"
import "strings"
import "math"
import "strconv"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    // N: the number of temperatures to analyse
    var N int
    scanner.Scan()
    fmt.Sscan(scanner.Text(),&N)
    
    scanner.Scan()
    
    TEMPS := strings.Split(scanner.Text()," ") // the N temperatures expressed as integers ranging from -273 to 5526

    res := 10000 //to make its overwritten by any value

    for i := 0; i < N; i++ {
    	tint,_ := strconv.ParseInt(TEMPS[i],10,0)
    	temp := float64(tint)
    	if math.Abs(temp) < math.Abs(float64(res)) { //nearer to 0?
    		res = int(temp)
    	} 

    	if temp == float64(res)*-1{ //is positive while as close as nearest?
    		res = int(math.Abs(temp))
    	}
    }

    if res == 10000 { //is not overwritten?
    	res = 0
    }
    
    // fmt.Fprintln(os.Stderr, "Debug messages...")
    
    fmt.Println(int(res))// Write answer to stdout
}