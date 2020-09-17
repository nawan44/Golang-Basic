for a :=0; 
	a < 5;
	a++{
		fmt.Println("Number", 1)
	}

var b = 0
for b < 5 {
	fmt.Println("Number1", 1)
	b++
}

var c = 0
for {
	fmt.Println("Number2", c)
	c++
	if c == 7 {
		break
	}
}

for d := 1;
	d <= 9;
	d++ {
		if d % 2 == 1{
			continue
		}

		if d > 8 {
			break
		}
		fmt.Println("Number", d)
	}

for e := 0;
	e < 6;
	e++ {
		for f := e;
		f < 5;
		f++ {
			fmt.Prin(f, "")
		}
		fmt.Println
	}

outerLoop:
for i := 0; i < 5; i++ {
	for j := 0; j < 5; j++ {
		if i == 3 {
			break outerLoop
		}
		fmt.Print("matriks [", i, "][", j, "]", "\n")
	}
}

