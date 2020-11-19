package bst

import(

)

type Tree struct {
	root *TreeNode
}

func(node *TreeNode) Inorder(fn func(node *TreeNode)){
	if node != nil{
		node.Left.Inorder(fn)
		fn(node)
		node.Right.Inorder(fn)
	}
}

func(t *Tree) Flatten() []string {
	var keys []string
	fn:= func(node *TreeNode){
		keys = append(keys,node.Data)
	}
	t.root.Inorder(fn)
	return keys
}

func newNode(data string) *TreeNode{
	var node TreeNode
	node.Data = data
	return &node
}

func (t *Tree) Add(key string){
	node:=newNode(key)
	t.insert(node)
}

func (t *Tree) insert (node *TreeNode){
	if t.root  == nil {
		t.root = node
	}
	//find correct position
	var ptr *TreeNode = nil
	rootNode := t.root

	for rootNode !=nil{
		ptr = rootNode
		if node.Data < rootNode.Data{
			rootNode = rootNode.Left
		}else{
			rootNode = rootNode.Right
		}
	}
	// inset Node
	if node.Data < ptr.Data{
		ptr.Left = node
	}else{
		ptr.Right = node
	}
	node.Parent = ptr
}

func (node *TreeNode) IsLeaf() bool{
	return node.Right==nil && node.Left==nil
}

