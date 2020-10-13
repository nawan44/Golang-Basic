package variadic

import "fmt"

func variadic) {
    var avg = calculate(2, 4, 3, 1, 4, 3, 9, 5, 8, 7)
    var msg = fmt.Sprintf("Average : %.2f", avg)
    fmt.Println(msg)
}

func calculate(numbers ...int) float64 {
    var summary int = 0
    for _, number := range numbers {
        summary += number
    }

    var avg = float64(summary) / float64(len(numbers))
    return avg
}