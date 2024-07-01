// package main

// import "fmt"

// type contactInfo struct {
// 	email   string
// 	zipcode int
// }

// type person struct {
// 	firstname string
// 	lastname  string
// 	contact   contactInfo
// }

// func main() {

// 	jim := person{
// 		firstname: "Jim",
// 		lastname:  "Dunphy",
// 		contact: contactInfo{
// 			email:   "jim@gmail.com",
// 			zipcode: 24656,
// 		},
// 	}

// 	jimmypointer := &jim
// 	jimmypointer.updatename("Jimmy")

// 	jim.print()
// }

// func (pointedvalue *person) updatename(newfirstame string) {
// 	pointedvalue.firstname = newfirstame
// }

// func (p person) print() {
// 	fmt.Printf("%+v", p)
// }

package main

import "fmt"

func main() {
	name := "bill"

	namePointer := &name

	fmt.Println(&namePointer)
	printPointer(namePointer)
}

func printPointer(namePointer *string) {
	fmt.Println(&namePointer)
}
