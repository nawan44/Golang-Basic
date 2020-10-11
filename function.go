package function

import "fmt"
import "strings"

func main() {
    var names = []string{"Rachmat", "Gunawan"}
    printMessage("halo", names)
}

func printMessage(message string, arr []string) {
    var nameString = strings.Join(arr, " ")
    fmt.Println(message, nameString)
}