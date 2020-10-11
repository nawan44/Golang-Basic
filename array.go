var names [4]string
names[0] = "HTML"
names[1] = "CSS"
names[2] = "Javascript"
names[3] = "Web"

fmt.Println(names[0], names[1], names[2], names[3])

var vegetables = [4]string{"spinach", "tomato", "garlic", "peanut"}

fmt.Println("Jumlah element \t\t", len(fruits))
fmt.Println("Isi semua element \t", fruits)

var val = [...]int{2, 3, 2, 4, 3}

fmt.Println("data array \t:", val)
fmt.Println("jumlah elemen \t:", len(numvalbers))

var numbers1 = [2][3]int{[3]int{3, 2, 3}, [3]int{3, 4, 5}}
var numbers2 = [2][3]int{{3, 2, 3}, {3, 4, 5}}

fmt.Println("numbers1", numbers1)
fmt.Println("numbers2", numbers2)