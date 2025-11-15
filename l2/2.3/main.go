package main

import (
	"fmt"
	"os"
)

// В файле runtime/runtime2.go можно посмотреть устройство интерфейса
// type iface struct {
//  метаданные об интерфейсе и о типе хранимого значения
// 	tab  *itab
//  ссылка на хранимое значение, которое реализует интерфейс
// 	data unsafe.Pointer
// }

// А это устройство пустого интерфейса
// type eface struct {
//  информация о типе хранимого значения
// 	_type *_type
//  ссылка на хранимое значение, которое реализует интерфейс
// 	data  unsafe.Pointer
// }

// Данный метод не будет возвращать нулевой интерфейс,
// так как тут интерфейс приводится к типу os.PathError
// и получается, что data в интерфесе nil, а tab содержит инфорацию
func Foo() error {
	var err *os.PathError = nil
	return err
}

// А в этом методе интерфейс не приводится ни к какому типу,
// поэтому он будет nil
func Foo2() error {
	var err error = nil
	return err
}

type Empty interface {
}

func main() {
	// ненулевой интерфейс с нулевыми данными
	err := Foo()
	// выводит данные интерфейса, которые = nil
	fmt.Println(err)
	// сравнивает интерфейс с nil, интерфейс != nil
	fmt.Println(err == nil)

	// нулевой интерфейс
	err2 := Foo2()
	// выводит данные интерфейса, которые = nil
	fmt.Println(err2)
	// сравнивает интерфейс с nil, интерфейс = nil
	fmt.Println(err2 == nil)
}
