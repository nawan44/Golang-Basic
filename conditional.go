// if, else if, & else

var grade = 7

if grade == 8 {
	fmt.Println("Pass Cumlaude")
} else if grade > 5 {
	fmt.Println("Pass")
} else if grade == 4 {
	fmt.Println("Almost graduated")
}else {
	fmt.Println("Not Pass. Your Grade %d\n", grade)
}

var point = 8840.0

if percent := point / 100; percent >= 100 {
    fmt.Printf("%.1f%s perfect!\n", percent, "%")
} else if percent >= 70 {
    fmt.Printf("%.1f%s good\n", percent, "%")
} else {
    fmt.Printf("%.1f%s not bad\n", percent, "%")
}

var value = 6
switch point {
case 8 :
	fmt.Println("Cumlaude")
case 5, 6, 7 :
	fmt.Println("Pass")
default :
	fmt.Println("Not Pass")
	fmt.Println("you can be better!")

}


var value = 6

switch {
case value == 8:
    fmt.Println("perfect")
case (value < 8) && (value > 3):
    fmt.Println("awesome")
    fallthrough
case value < 5:
    fmt.Println("you need to learn more")
default:
    {
        fmt.Println("not bad")
        fmt.Println("you need to learn more")
    }
}



