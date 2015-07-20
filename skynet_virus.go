package main

import "fmt"
//import "os"


/*
This solution works with all challenges, but you can't get "at least 50 links left" achievement because it randomly gets it (not always) :/
*/
type Node struct {
	isGateway bool
	nearestGateway int
	linkedTo []int
}

func main() {
    // N: the total number of nodes in the level, including the gateways
    // L: the number of links
    // E: the number of exit gateways
    var N, L, E int
    fmt.Scan(&N, &L, &E)

    nodes := make([]Node,N)
    gateways := make([]int,E)
    wnodes := make([]int,0,10)

    for i := 0; i < N; i++ {
    	nodes[i].nearestGateway = -1 //struct would set it to default which is actual node id, so I could break whole code
    	nodes[i].linkedTo = make([]int,0,10) //make an empty slice for nodes linked to this one, with capacity of 10
    }
    
    for i := 0; i < L; i++ {
        // N1: N1 and N2 defines a link between these nodes
        var N1, N2 int
        fmt.Scan(&N1, &N2)
        nodes[N1].linkedTo = append(nodes[N1].linkedTo, N2)
        nodes[N2].linkedTo = append(nodes[N2].linkedTo, N1)
    }
    for i := 0; i < E; i++ {
        // EI: the index of a gateway node
        var EI int
        fmt.Scan(&EI)
        nodes[EI].isGateway = true
        gateways[i] = EI
        for j := 0; j < len(nodes[EI].linkedTo); j++ {
        	nodes[nodes[EI].linkedTo[j]].nearestGateway = EI
        	wnodes = append(wnodes,nodes[EI].linkedTo[j])
        }
    }
    for i := 0; i < N; i++ {
    	if nodes[i].nearestGateway != -1 {
    		for j := 0; j < len(nodes[i].linkedTo); j++ {
    			if nodes[nodes[i].linkedTo[j]].nearestGateway == -1 {
    				nodes[nodes[i].linkedTo[j]].nearestGateway = i
    			}
    		}
    	}
    }
    for {
        // SI: The index of the node on which the Skynet agent is positioned this turn
        var SI int
        fmt.Scan(&SI)

        if nodes[SI].nearestGateway == -1 {
            temp := -1
            for i := 0; i < len(nodes[SI].linkedTo); i++{
                for j := 0; j < len(wnodes); j++ {
                    if nodes[SI].linkedTo[i] == wnodes[j] {
                        temp = i
                        break
                    }
                }
                if temp > -1 {
                    break
                }
            }
            fmt.Println(SI,nodes[SI].linkedTo[temp])
        } else {
            fmt.Println(SI,nodes[SI].nearestGateway)
            nodes[SI].nearestGateway = -1
        }
        
        // fmt.Fprintln(os.Stderr, "Debug messages...")
        
        //fmt.Println("0 1") // Example: 0 1 are the indices of the nodes you wish to sever the link between
    }
}