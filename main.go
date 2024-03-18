package main

import "fmt"

const SIZE = 5

type Node struct {
	val   string
	Left  *Node
	Right *Node
}

type Queue struct {
	Head   *Node
	Tail   *Node
	Length int
}

type Cache struct {
	Queue Queue
	Hash  Hash
}

func NewCache() *Cache {
	return &Cache{Queue: NewQueue(), Hash: Hash{}}
}

func NewQueue() Queue {
	head := &Node{}
	tail := &Node{}

	head.Right = tail
	tail.Left = head

	return Queue{Head: head, Tail: tail, Length: 0}

}

func (c *Cache) Check(str string) {
	node := &Node{}

	if val,ok := c.Hash[str];ok{
		node = c.Remove(val)
	}else{
		node = &Node{val: str}
	}

	c.Add(node)
	c.Hash[str] = node
}

func (c *Cache) Remove(n *Node) *Node {
	fmt.Printf("Remove: %s", n.val)

	left := n.Left
	right := n.Right

	left.Right = right
	right.Left = left

	c.Queue.Length -= 1
	delete(c.Hash, n.val)
	return n

}

func (c *Cache) Add(n *Node) {
	fmt.Printf("add: %s\n", n.val)
	temp := c.Queue.Head.Right

	c.Queue.Head.Right = n
	n.Left = c.Queue.Head
	n.Right = temp
	temp.Left = n

	c.Queue.Length += 1

	if c.Queue.Length > SIZE {
		c.Remove(c.Queue.Tail.Left)
	}
}

func (c *Cache) Display() {
	c.Queue.Display()
}

func (q *Queue) Display() {
	node := q.Head.Right

	fmt.Printf("%d - [", q.Length)

	for i := 0; i < q.Length; i++ {
		fmt.Printf("{%s}", node.val)
		if i < q.Length {
			fmt.Printf("<--->")
		}
		node = node.Right
	}
	fmt.Println("]")
}

type Hash map[string]*Node

func main() {
	fmt.Println("START CACHE")

	cache := NewCache()

	for _, word := range []string{
		"parrot", "avacado", "tree", "potato", "tomato", "cat", "dog", "horse", "cat",
	} {
		cache.Check(word)
		cache.Display()
	}	
}
