package main

import "fmt"

var a bool = true;

func computeValue() int {
	return 11
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func addsub(ab, bb int) (add , sub int){
	add = ab + bb
	sub = ab - bb
	return
}

func printParam(args ...int){
	for _, n := range args {
		fmt.Printf("Number is %d\n", n)
	}
}

func ReadWrite() {
	for i:=1; i <= 5 ; i++ {
		defer fmt.Printf("file closed... %d\n", i)
	}
	
	fmt.Println("working on file...")
	fmt.Println("Done")
	
}

func main(){

	if a {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}

	if x := computeValue(); x < 10 {
		fmt.Println("Less than 10")
	} else if x > 10 {
		fmt.Println("more than 10")
	} else if x == 10 {
		fmt.Println("Equal")
	}
	sum := 0
	for x := 1; x <= 10; x++ {
		fmt.Printf("%d ", x)
		sum += x
	}
	fmt.Printf("\n")
	fmt.Println(sum)

	for sum <= 200 {
		sum += 55
	}
	fmt.Println(sum)

	val := map[string]int{"a":1,"b":2,"c":3}

	for k,v := range val {
		fmt.Printf("%s:%d\n",k,v)
	}

	types := "a"

	switch types {
		case "a": fallthrough
		case "e": fallthrough
		case "i": fallthrough
		case "o": fallthrough
		case "u":
			fmt.Println("Vowel")
		default:
			fmt.Println("Consonant")
	}

	g, h := addsub(3,1)

	fmt.Println(max(1, 2))
	fmt.Printf("%d %d", g, h)

	printParam(1, 2, 3, 4)

	ReadWrite()
}