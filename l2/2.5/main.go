package main

import (
	"fmt"
	"unsafe"
)

// В файле runtime/runtime2.go можно посмотреть устройство интерфейса
// type iface struct {
//  метаданные об интерфейсе и о типе хранимого значения
// 	tab  *itab
//  ссылка на хранимое значение, которое реализует интерфейс
// 	data unsafe.Pointer
// }

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	// ... do something
	return nil
}

func main() {
	// err и err2 нулевые интерфейсы типа error
	// data unsafe.Pointer и tab *itab равны nil
	var err error
	var err2 error

	// интерфейс err приводится к типу customError и становится ненулевым
	// то есть data unsafe.Pointer указывает на nil, но в tab *itab
	// теперь храняться данные о типе data
	err = test()

	// так как интерфейс ненулевой, то выведет error
	foo(err)
	// показывает указатели data и tab,
	// tab не нулевой, data нулевой
	checkIface(err)
	// так как интерфейс нулевой, то выведет ok
	foo(err2)
	// показывает указатели data и tab,
	// tab нулевой, data нулевой
	checkIface(err2)
}

func foo(err error) {
	if err != nil {
		println("error")
		return
	}
	println("ok")
}

func checkIface(err error) {

	type iface struct {
		tab  unsafe.Pointer
		data unsafe.Pointer
	}

	i := (*iface)(unsafe.Pointer(&err))

	fmt.Printf("itab pointer = %p\n", i.tab)
	fmt.Printf("data pointer = %p\n", i.data)
}
