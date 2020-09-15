var value = (((3+7)%4)*2 - 1)/5

var isEqual = (value == 2)
fmt.Printf("value %d (%t) \n", value, isEqual)

var left = false
var right = true

var leftAndRight = left && right
fmt.Printf("left && right \t(%t) \n", leftAndRight)

var leftOrRight = left || right
fmt.Printf("left || right \t(%t) \n", leftOrRight)

var leftReverse = !left
fmt.Printf("!left \t\t(%t) \n", leftReverse)