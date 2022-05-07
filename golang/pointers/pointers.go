package pointers

import "fmt"

type Product struct {
	ID       int64
	Type     string
	Name     string
	FullName string
}

func Run() {
	pointer := &Product{
		1, "Type", "Name", "",
	}
	fullNamePointer := getFullNamePointer(pointer)
	fmt.Println("fullNamePointer", fullNamePointer)
	fmt.Println("pointer", pointer)

	value := Product{
		1, "Type", "Name", "",
	}
	fullNameValue := getFullNameValue(value)
	fmt.Println("fullNameValue", fullNameValue)
	fmt.Println("value", value)
}

func getFullNamePointer(in *Product) string {
	in.FullName = in.Type + in.Name
	return in.FullName
}

func getFullNameValue(in Product) string {
	in.FullName = in.Type + in.Name
	return in.FullName
}
