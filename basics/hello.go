package main

import "fmt"
import "errors"

var abcd = "1234"

const(
    x = iota  // x == 0
    y = iota  // y == 1
    z = iota  // z == 2
    w  // If there is no expression after the constants name, it uses the last expression,
    //so it's saying w = iota implicitly. Therefore w == 3, and y and z both can omit "= iota" as well.
)

const v = iota // once iota meets keyword `const`, it resets to `0`, so v = 0.

const (
  e, f, g = iota, iota, iota // e=0,f=0,g=0 values of iota are same in one line.
)

const PI = 3.14558364
var trueS = true

var a int32
var b int32

const(
	i = 100
	j = 200
)

var numbers map[string] int


var complex complex64 = 5 + 5i

func main(){
	av := false

	var arr [10]int

	newarr := []byte{'a','b','c','d'}

	maps := make(map[string]int)
	
	maps["abcd"] = 1
	maps["abcde"] = 2
	fmt.Println(maps)
	fmt.Println(len(maps))

	g := newarr[0:1]

	fmt.Println(newarr)
	fmt.Println(g)
	
	newarr[0] = 'b'

	fmt.Println(newarr)
	fmt.Println(g)

	arr[0] = 45
	arr[1] = 115

	arr2 := [3]int{1,2,3}
	arr3 := [...]int{1,2,3,4}
	md_arr := [2][4]int{{1,1,2,3}, {5,6,7,8}}

	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(md_arr)

	m := `
	hello world
	abcd
	`
	fmt.Println(arr)
	fmt.Println(x + y + z + w)

	err := errors.New("Error occured")
	if err != nil {
		fmt.Print(err)
	}

	s := "hello"
	c := []byte(s) // convert string to byte
	c[0] = 'c'
	s2 := string(c)

	fmt.Println(m)
	fmt.Println(s2)
	fmt.Printf("String: %s\n", s[1:])


	fmt.Println(complex)
	fmt.Println(trueS && av)
	fmt.Println(abcd);
	fmt.Println(PI)
	fmt.Printf("Hello World\n")
}