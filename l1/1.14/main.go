package main

import "fmt"

func main() {
	fmt.Println(getType(5))
	fmt.Println(getType("str"))
	fmt.Println(getType(false))
	fmt.Println(getType(make(chan int)))
	fmt.Println(getType(make(chan string)))
	fmt.Println(getType(make(chan bool)))
	fmt.Println(getType(make(chan struct{})))
	fmt.Println(getType(make(chan int, 1)))
	fmt.Println(getType(make(chan string, 2)))
	fmt.Println(getType(make(chan bool, 3)))
	fmt.Println(getType(make(chan struct{}, 4)))
	fmt.Println(getType(5.6))

}

func getType(val interface{}) string {
	switch v := val.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	case chan int:
		return fmt.Sprintf("chan int, buffer: %d", cap(v))
	case chan string:
		return fmt.Sprintf("chan string, buffer: %d", cap(v))
	case chan bool:
		return fmt.Sprintf("chan bool, buffer: %d", cap(v))
	case chan struct{}:
		return fmt.Sprintf("chan struct, buffer: %d", cap(v))
	default:
		return "unknown type"
	}

}
