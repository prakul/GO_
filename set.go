//Implementation of a Set in GOLANG using maps
// Set can contain a string value and supports
// Add, Remove, Exists operation
package main

import "fmt"

type set struct {
	set map[string]bool
}

func NewSet() *set {
	s := &set{make(map[string]bool)}
	return s
}

func (s *set) Add(value string) {
	s.set[value] = true
}

func (s *set) Remove(value string) {
	delete(s.set, value)

}

func (s *set) Exists(value string) bool {
	_, ok := s.set[value]
	return ok
}


func main() {
	//Create an empty set class
	set := NewSet()

	//Add elements to the set
	set.Add("k1")
	set.Add("k2")
	fmt.Println("set:", set)

	//Check if an element exists in set
	fmt.Println("k2", set.Exists("k2"))
	fmt.Println("k3", set.Exists("k3"))

	//Remove an element
	set.Remove("k2")
	fmt.Println("k2",set.Exists("k2"))
	fmt.Println("k1", set.Exists("k1"))
}

