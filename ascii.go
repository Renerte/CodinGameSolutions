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

    var L int
    scanner.Scan()
    fmt.Sscan(scanner.Text(),&L)
    
    var H int
    scanner.Scan()
    fmt.Sscan(scanner.Text(),&H)

    rows := make([][]string,H)
    
    lt := strings.Split("ABCDEFGHIJKLMNOPQRSTUVWXYZ?","")

    scanner.Scan()
    T := strings.Split(strings.ToUpper(scanner.Text()),"")
    for i := 0; i < H; i++ {
        scanner.Scan()
        rows[i] = strings.Split(scanner.Text(),"")
    }

    res := make([]string,H)

    for h := 0; h < H; h++{
        for t := 0; t < len(T); t++ {
            found := false
            for l := 0; l < len(lt)-1; l++ {
                if lt[l] == T[t] {
                    fmt.Fprintln(os.Stderr,"Parsing",l*L,"to",(l+1)*L-1)
                    temp := rows[h][l*L:(l+1)*L]
                    //rows[h][l*(L+1)+1:(l+1)*(L+1)+1]
                    res[h] += strings.Join(temp,"")
                    found = true
                }
            }
            if found == false {
                res[h] += strings.Join(rows[h][(len(lt)-1)*L:(len(lt))*L],"")    
            }
        }
    }

    
    for h := 0; h < H; h++ {
        fmt.Println(res[h])
    }
    // fmt.Fprintln(os.Stderr, "Debug messages...")
    
    //fmt.Println("answer")// Write answer to stdout
}