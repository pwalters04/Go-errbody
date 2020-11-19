package ClosestValueBST

type BST struct {
	Value int
	Left  *BST
	Right *BST
}

func (tree *BST) FindClosestValue(target int) int {
	return tree.findClosestValue(target,tree.Value)
}

func (tree *BST) findClosestValue(target,closest int) int{

	if absValueDiff(target,closest) > absValueDiff(target,tree.Value){
		closest = tree.Value
	}

	if target < tree.Value && tree.Left !=nil{
		return tree.Left.findClosestValue(target,closest)
	}else if target > tree.Value && tree.Right !=nil{
		return tree.Right.findClosestValue(target,closest)
	}

	return closest
}


func absValueDiff (x ,y int) int{
	if x > y{
		return x-y
	}
	return y-x
}
