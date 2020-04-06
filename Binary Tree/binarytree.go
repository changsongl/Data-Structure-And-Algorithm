package Binary_Tree

type BinaryTree struct {
	root *Node
}

type Node struct {
	data  int
	left  *Node
	right *Node
}

func New() *BinaryTree {
	return &BinaryTree{}
}

func (bt *BinaryTree) Insert(num int) {
	if bt.root == nil {
		bt.root = &Node{data: num}
		return
	}

	currentNode := bt.root
	for {
		if num > currentNode.data {
			if currentNode.right == nil {
				currentNode.right = &Node{data: num}
				return
			}
			currentNode = currentNode.right
		} else {
			if currentNode.left == nil {
				currentNode.left = &Node{data: num}
				return
			}
			currentNode = currentNode.left
		}
	}
}

func (bt *BinaryTree) Search(num int) (*Node, bool) {
	currentNode := bt.root

	for currentNode != nil {
		if currentNode.data == num {
			return currentNode, true
		} else if num > currentNode.data {
			currentNode = currentNode.right
		} else {
			currentNode = currentNode.left
		}
	}
	return nil, false
}

func (bt *BinaryTree) InOrderTraversal(st *Node, callback func(int)) {
	if st.left != nil {
		bt.InOrderTraversal(st.left, callback)
	}
	callback(st.data)
	if st.right != nil {
		bt.InOrderTraversal(st.right, callback)
	}
}

func (bt *BinaryTree) PreOrderTraversal(st *Node, callback func(int)) {
	callback(st.data)
	if st.left != nil {
		bt.PreOrderTraversal(st.left, callback)
	}
	if st.right != nil {
		bt.PreOrderTraversal(st.right, callback)
	}
}

func (bt *BinaryTree) PostOrderTraversal(st *Node, callback func(int)) {
	if st.left != nil {
		bt.PostOrderTraversal(st.left, callback)
	}
	if st.right != nil {
		bt.PostOrderTraversal(st.right, callback)
	}
	callback(st.data)
}
