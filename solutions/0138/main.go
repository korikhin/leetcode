package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	for curr := head; curr != nil; curr = curr.Next.Next {
		curr.Next = &Node{Val: curr.Val, Next: curr.Next}
	}
	for curr := head; curr != nil; curr = curr.Next.Next {
		if curr.Random != nil {
			curr.Next.Random = curr.Random.Next
		}
	}

	dummy := &Node{}
	for currNew, currOld := dummy, head; currOld != nil; {
		currNew.Next = currOld.Next
		currOld.Next = currOld.Next.Next

		currOld = currOld.Next
		currNew = currNew.Next
	}

	return dummy.Next
}
