package main 
import "fmt"

type Node struct{
 	value int
 	next *Node
}

type SingleList struct{
	len int
	head *Node
}

func (L *SingleList) add(val int){
	node := &Node{value: val}
	if L.head == nil{
		L.head = node
	} else{
		node.next = L.head
		L.head = node
	}
	L.len++
	return
}

func (L *SingleList) size() int{
	return L.len
}

func main(){
	/*
	lista := ListNode{value: 1}
	lista.next = &ListNode{value: 2}
	fmt.Println(lista)
	fmt.Println(lista.next.value)
	*/
	//array := []int{1,2,3,4,5}
	lista := &SingleList{len: 0}
	fmt.Println(lista, lista.size)
	lista.add(1)
	fmt.Println(lista, lista.size)
	lista.add(2)
	fmt.Println(lista, lista.size)
}
