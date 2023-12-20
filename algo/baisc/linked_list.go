package basic

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(val int) *ListNode {
	return &ListNode{Val: val}
}

// Insert 插入节点
func (node *ListNode) Insert(val int) *ListNode {
	newNode := NewListNode(val)
	newNode.Next = node.Next
	node.Next = newNode
	return node
}

// Delete 删除节点
func (node *ListNode) Delete() *ListNode {
	node.Next = node.Next.Next
	return node
}

// Traverse 遍历节点
func (node *ListNode) Traverse() []int {
	var result []int
	for node != nil {
		result = append(result, node.Val)
		node = node.Next
	}
	return result
}

// Find 查找节点
func (node *ListNode) Find(val int) *ListNode {
	for node != nil {
		if node.Val == val {
			return node
		}
		node = node.Next
	}
	return nil
}

// Reverse 链表反转
func (node *ListNode) Reverse() *ListNode {
	var prev *ListNode
	for node != nil {
		next := node.Next
		node.Next = prev
		prev = node
		node = next
	}
	return prev
}
