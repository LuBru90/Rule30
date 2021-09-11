package main

import "fmt"

const SIZE = 250
const ITERS = 200

var RULE = map[int] int {
    0: 0,
    1: 1,
    2: 1,
    3: 1,
    4: 1,
    5: 0,
    6: 0,
    7: 0,
}

func printState(state [SIZE]int) {
    temp := ""
    for _, value := range state {
        if value == 0 {
            temp += " "
        } else {
            temp += "#"
        }
    }
    fmt.Println("->", temp)
}

func applyRule(state [SIZE]int) (newState [SIZE]int) {
    var result int
    for i := 0; i < len(state) - 2 ; i++ {
        out := state[i:i+3]
        result = out[0] + out[1]*2 + out[2]*4
        newState[i+1] = RULE[result]
    }
    return newState
}

func main() {
    var state [SIZE]int
    state[SIZE/2] = 1

    for i := 0; i < ITERS; i++ {
        printState(state)
        state = applyRule(state)
    }
}
