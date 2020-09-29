var month map[string]int
month = map[string]int{}
month["Decembe"] = 40
month["February"] = 30

fmt.Println("Decembe", month["Decembe"])
fmt.Println("May", month["May"])

var aaa map[string]int
aaa["one"] = 1

aaa = map[string]int{}
aaa["one"] = 1

//vertical Method
var chicken1 = map[string]int{"Decembe": 50, "February": 40}

//horizontal Method
var chicken2 = map[string]int{
    "Decembe":  50,
    "February": 40,
}

//for - range

var month = map[string]int{
    "Decembe":  50,
    "February": 40,
    "August":    34,
    "September":    67,
}

for key, val := range month {
    fmt.Println(key, "  \t:", val)
}


//Delete Item Map
var month = map[string]int{"Decembe": 50, "February": 40}

fmt.Println(len(month)) // 2
fmt.Println(month)

delete(month, "Decembe")

fmt.Println(len(month)) // 1
fmt.Println(month)

// with Key
var month = map[string]int{"Decembe": 50, "February": 40}
var value, isExist = month["May"]

if isExist {
    fmt.Println(value)
} else {
    fmt.Println("item is not exists")
}

//Combination slice & map
var months = []map[string]string{
    map[string]string{"name": "month blue",   "gender": "male"},
    map[string]string{"name": "month red",    "gender": "male"},
    map[string]string{"name": "month yellow", "gender": "female"},
}

for _, month := range months {
    fmt.Println(month["gender"], month["name"])
}


