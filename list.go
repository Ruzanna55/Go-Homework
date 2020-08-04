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
	add(&head, 10)
	add(&head, 11)
	fmt.Println(head)
	fmt.Println("Length =", Length())
	//Delete(3)
	Insert(4, 5)
	print(head)
	fmt.Println(At(2))
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

func print(head *listNode) {
	i := head
	for ; i != nil; i = i.next {
		println(i.value)
	}
}

//Length returns the list item count
func Length() uint {
	var count uint
	i := head
	for ; i != nil; i = i.next {
		count++
	}
	return count
}

//Delete the item at given position or return error if the position is outside the list boundaries
func Delete(position uint) bool {
	i := head
	var count uint

	if position > Length() {
		fmt.Println("the position is outside the list boundaries")
		return false
	}
	for ; i != nil; i = i.next {
		count++
		if count == (position - 1) {
			i.next = i.next.next
			fmt.Println("the element was deleted")

		}

	}
	return true
}

//At return the value at given position, or return error if the index is outside the list boundaries
func At(position uint) (int, bool) {
	i := head
	var count uint
	var valueOfPosition int
	if position > Length() {
		fmt.Println("the position is outside the list boundaries")
		return 0, false
	}
	for ; i != nil; i = i.next {
		count++
		if count == position {
			valueOfPosition = i.value

		}
	}
	return valueOfPosition, true
}

// Insert add the value at the given position
func Insert(value int, position uint) {
	i := head
	var count uint
	node := new(listNode)
	node.value = value
	if position > (Length() + 1) {
		fmt.Println("the position is outside the list boundaries")
	}

	for ; i != nil; i = i.next {
		count++

		if count == (position - 1) {
			a := i.next
			i.next = node
			node.next = a
			fmt.Println("value was  added  at the given position")
		}
	}

}
