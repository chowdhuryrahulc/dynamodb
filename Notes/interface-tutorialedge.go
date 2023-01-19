package main

import "fmt"

/*
type Handler struct {
	handler.Interface
	Controller        product.Interface
	Rules             Rules.Interface
}
handler.Interface: implements all the funcs present in the interface as struct funcs (for more explaination see below)
Controller  product.Interface: funcs inside product.interface are implimented INSIDE struct funcs
  (for more explaination, see internal/handler/product/product.go. Just delete the 2 lines ie: "Controller product.Interface", an the next)
*/

/*
* TutorialEdge:
	we implement a interface named Employee whic has a function named GetName()
	and this function is implemented as methord by 2 structs named Engineer and Manager
	func printdetails takes in interface, ut since both the structs implement interface,
	so printdetails also takes in the 2 structs
  todo Try implimenting in another struct, using something like eg: handlers.interfaces

  you can also write cli in go
  interfaces: adv build reusable and extensible code
  we do not need to write "implements interface", like in flutter. Just use the interface methods
  Like wall socket example of connection btw 2 stuff
  Implements abstraction and polymorphism
*/

type Employee interface {
	GetName() string
}

type Engineer struct {
	Name string
}

func (e *Engineer) GetName() string {
	return "Engineer Name: " + e.Name
}

type Manager struct {
	Name string
}

func (m *Manager) GetName() string {
	return "Manager Name: " + m.Name
}

func PrintDetails(e Employee) {
	fmt.Println(e.GetName())
}

// Type assersion in interfaces
func myfun(a interface{}) {
	val := a.(string)
	fmt.Println("Value: ", val)
}

func main() {
	engineer := &Engineer{Name: "Elliot"}
	manager := &Manager{Name: "Donna"}

	PrintDetails(engineer) // output: Engineer Name: Elliot
	PrintDetails(manager)  // output: Manager Name: Donna

	//*******************************************************************
	m := Manager{Name: "from geeks for geeks"}
	fmt.Println(m) // output: m value is:  {from geeks for geeks}
	//*******************************************************************

	// Type assersion: we put a type in the interface (in this case, a string)
	var val interface{} = "geeks part2"
	myfun(val)

}
