package main

import "fmt"
import "strconv"
import "strings"

//import "os"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
	var R int
	fmt.Scan(&R)

	var L int
	fmt.Scan(&L)

	r := strconv.Itoa(R)

	res := "x" // 'x' at beginning to support multi-digit Rs

	for k := 1; k < L; k++ {
		temp := res
		res = ""
		for i := 0; i < len(temp); i++ {
			tchar := string(temp[i])
			rpt := 1
			for j := i + 1; j < len(temp); j++ {
				tchar2 := string(temp[j])
				if tchar2 == tchar || (tchar == r && tchar2 == "x") {
					rpt++
				} else {
					break
				}
				i = j
			}
			res += strconv.Itoa(rpt)
			res += tchar
		}
	}

	res_array_temp := strings.Split(res, "")

	res_array := make([]interface{}, len(res_array_temp)) //convert string array to interface one
	for i, v := range res_array_temp {
		if v == "x" {
			res_array[i] = r
			continue
		}
		res_array[i] = v
	}

	// fmt.Fprintln(os.Stderr, "Debug messages...")

	fmt.Println(res_array...) // Write answer to stdout
}
