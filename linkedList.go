package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	head   *node
	length int
}

func (l *linkedList) addNode(n *node) {
	temp := l.head
	l.head = n
	l.head.next = temp
	l.length++
}

func (l linkedList) printAll() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next
		l.length--
	}
	fmt.Printf("\n")
}

func (l *linkedList) deleteNode(value int) {

	if l.head.data == value {
		l.head = l.head.next
		l.length--
		return
	}

	start := l.head
	for start.next.data != value {
		if start.next.next == nil {
			return
		}
		start = start.next
	}
	start.next = start.next.next
	l.length--
}

func (l *linkedList) deleteLastNode() {

	start := l.head
	for start.next != nil {
		if start.next.next == nil {
			l.length--
			return
		}
		start = start.next
	}
}

func (l *linkedList) editNode(index, value int) {

	if index == 1 {
		start := l.head
		start.data = value
	} else {
		start := l.head
		i := 0
		val := start
		for i < index {
			val = start
			start = start.next
			i++
		}
		val.data = value

	}
}
func (l *linkedList) valueAt(index int) {

	if index == 1 {
		start := l.head
		fmt.Println(start.data)
	} else {
		start := l.head
		i := 0
		for i <= index+1 {
			if i == index {
				fmt.Println("Result is--->", start.data)
				return
			} else if i == index-1 {
				i++
			} else {
				start = start.next
				i++
			}

		}

	}
}
func main() {
	mylist := linkedList{}
	node1 := &node{data: 1}
	node2 := &node{data: 2}
	node3 := &node{data: 3}
	node4 := &node{data: 4}
	node5 := &node{data: 5}
	node6 := &node{data: 6}
	node7 := &node{data: 7}
	mylist.addNode(node1)
	mylist.addNode(node2)
	mylist.addNode(node3)
	mylist.addNode(node4)
	mylist.addNode(node5)
	mylist.addNode(node6)
	mylist.addNode(node7)
	mylist.printAll()

	//mylist.deleteNode(4)
	//mylist.printAll()
	//mylist.deleteLastNode()
	//mylist.editNode(4, 44)
	//mylist.printAll()
	//mylist.valueAt(2)

}
