package main

import "fmt"
import "os"
import "bufio"
//import "strings"
import "strconv"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    scanner.Scan()

    font := " !\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz" //all basic characters ordered by their position in ASCII

    charMap := make(map[string]int)


    for i := 0; i < len(font); i++ {
    	charMap[string(font[i])] = i+32 //assign ids to individual characters
    }

    MESSAGE := scanner.Text()

    idMap := make(map[int]int)

    for i := 0; i < len(MESSAGE); i++ { //get ids for message
    	temp := string(MESSAGE[i])
    	id := charMap[temp]
    	if id != 0 {
    		idMap[i] = id
    	}
    }

    res := ""

    res_string := ""

    for i := 0; i < len(MESSAGE); i++ { //make one long binary string
    	temp := strconv.FormatInt(int64(idMap[i]),2) //get binary string for this character
    	for j := len(temp); j < 7; j++ {
    		temp = "0" + temp
    	}
    	res_string += temp
    }

	for i := 0; i < len(res_string); i++ { //convert to Chuck Norris' Unary :P
		tchar := string(res_string[i])
		if tchar == "0" {
			res += "0"
		}
		res += "0 "
		rpt := 1 //repeats of same character
		for k := i+1; k < len(res_string); k++ {
			tchar2 := string(res_string[k]) //i know, ambitious variable naming :P
			if tchar2 == tchar {
				rpt++
			} else {
				break
			}
			i = k //skip repeated letters in parent loop
		}
		for k := 0; k < rpt; k++ {
			res += "0"
		}
		if (i+1 != len(res_string)){
			res += " "
		}
	}

    
    // fmt.Fprintln(os.Stderr, "Debug messages...")
    
    fmt.Println(res)// Write answer to stdout
}