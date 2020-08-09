package main

import "fmt"

type listNode struct {
	value int
	next  *listNode
}

var head *listNode

func main() {

	add(&head, 1)
	add(&head, 2)
	//add(&head, 10)
	//add(&head, 11)
	//add(&head, 12)
	fmt.Println(head)
	CreateCycle(&head)
	println(ExistenceOfCycle(head))

	//print(head)
	//fmt.Println(At(2))
}

func add(head **listNode, value int) {
	if nil == *head {
		*head = new(listNode)
		(*head).value = value
		return
	}
	i := *head
	for ; i.next != nil; i = i.next {
		// todo   handle cycle

	}
	i.next = new(listNode)
	i.next.value = value
}

//CreateCycle function create cycle in linked list
func CreateCycle(head **listNode) {
	i := *head
	for ; i.next != nil; i = i.next {
	}
	i.next = (*head).next
	fmt.Println("created cycle")

}

// ExistenceOfCycle function check is there cycle or not
func ExistenceOfCycle(head *listNode) bool {

	node1 := head
	node2 := head

	for ; node1.next != nil; node1 = node1.next {
		if node1.next.next == nil {
			fmt.Println("no;there isn't cycle")
			return false
		}
		node2 = node2.next.next
		if node1.next == node2.next {
			fmt.Println("yes: there is cycle")
			return true
		}
	}
	fmt.Println("no;there isn't cycle")
	return false
}
