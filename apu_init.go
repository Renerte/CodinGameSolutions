package main

import "fmt"
import "os"
import "bufio"
//import "strings"
import "strconv"

/**
 * Don't let the machines win. You are humanity's last hope...
 **/

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    // width: the number of cells on the X axis
    var width int
    scanner.Scan()
    fmt.Sscan(scanner.Text(),&width)
    
    // height: the number of cells on the Y axis
    var height int
    scanner.Scan()
    fmt.Sscan(scanner.Text(),&height)

    node := make([][]bool,height) //slice for nodes' rows

    for i := 0; i < height; i++ {
        node[i] = make([]bool,width) //slice - nodes' row
    }
    
    for i := 0; i < height; i++ { //find nodes
        scanner.Scan()
        line := scanner.Text() // width characters, each either 0 or .
        for j := 0; j < len(line); j++ {
            if string(line[j]) == "0" {
                node[i][j] = true
            }
        }
    }

    for i := 0; i < height; i++ {
        for j := 0; j < width; j++ {
            if node[i][j] { //is there a node?
                res := strconv.Itoa(j)+" "+strconv.Itoa(i)+" " //start of response
                for k := j+1; k < width; k++ {
                    if node[i][k] { //is there a node?
                        res += strconv.Itoa(k)+" "+strconv.Itoa(i)+" " //inner response
                        break
                    }
                }
                if len(res) < 7 { //was node found?
                    res += "-1 -1 "
                }
                for k := i+1; k < height; k++ {
                    if node[k][j] { //is there a node?
                        res += strconv.Itoa(j)+" "+strconv.Itoa(k) //end response
                        break
                    }
                }
                if len(res) < 11 { //was node found?
                    res += "-1 -1"
                }
                fmt.Println(res) //output response
            }
        }
    }
    
    // fmt.Fprintln(os.Stderr, "Debug messages...")
    
    //fmt.Println("0 0 1 0 0 1") // Three coordinates: a node, its right neighbor, its bottom neighbor
}