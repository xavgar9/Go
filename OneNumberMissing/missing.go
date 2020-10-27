package main

//import "fmt"

var exists = struct{}{}

type Set struct {
    M map[int]struct{}
}

func newSet() *Set {
    s := &Set{}
    s.M = make(map[int]struct{})
    return s
}

func (s *Set) add(value int) {
    s.M[value] = exists
}

func (s *Set) delete(value int) {
	delete(s.M, value)
}

func (s *Set) contains(value int) bool {
    _, c := s.M[value]
    return c
}

func (s *Set) getAll() []int{
	var array []int	
	for key := range s.M{ 
		array = append(array, key)
	}
	return array
}

func missing(array []int) int{
	set := newSet()
	
	for i:=1; i<=len(array); i++{ set.add(i) }

	for i:=0; i<len(array); i++{
		set.delete(array[i])
	}
	return set.getAll()[0]
}
/*
func main(){
	array := []int{1,2,4,5}
	fmt.Println(solve(array))
}
*/