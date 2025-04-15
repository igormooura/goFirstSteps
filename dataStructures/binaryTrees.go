package dataStructures

type Tree struct {
	Value int
	Left  *Tree
	Right *Tree
}

func NewTree(value int) *Tree {
	return &Tree{Value: value}
}

func Insert(root *Tree, value int) *Tree {
	if root == nil {
		return NewTree(value)
	}
	if value < root.Value {
		root.Left = Insert(root.Left, value)
	} else {
		root.Right = Insert(root.Right, value)
	}
	return root
}

func Delete(root *Tree, value int) *Tree {
	if root == nil {
		return nil
	}
	if value < root.Value {
		root.Left = Delete(root.Left, value)
	} else if value > root.Value {
		root.Right = Delete(root.Right, value)
	} else {
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}
		minLargerNode := root.Right
		for minLargerNode.Left != nil {
			minLargerNode = minLargerNode.Left
		}
		root.Value = minLargerNode.Value
		root.Right = Delete(root.Right, minLargerNode.Value)
	}
	return root
}

func TraverseInOrder(root *Tree, values *[]int) {
	if root == nil {
		return
	}
	TraverseInOrder(root.Left, values)
	*values = append(*values, root.Value)
	TraverseInOrder(root.Right, values)
}

func PrintTree(root *Tree, level int) {
	if root == nil {
		return
	}
	PrintTree(root.Right, level+1)
	PrintTree(root.Left, level+1)
}
