//package bst
// import (
//	 "errors"
// )
//
//type Tree struct {
//	Root *Node
//}
//
//func( t *Tree) Insert(value string ,data string) error {
//	if t.Root == nil{
//		t.Root = &Node{Value: value, Data: data}
//		return nil
//	}
//	return t.Root.Insert(value,data)
//}
//
//func (t *Tree) Find(s string) (string,bool){
//	if t.Root == nil{
//		return "",false
//	}
//	return t.Root.Find(s)
//}
//
//func (t *Tree) Delete (s string) error{
//	if t.Root == nil {
//		return errors.New("empty tree")
//	}
//	fakeParent := &Node{Right: t.Root}
//	err:= t.Root.Delete(s, fakeParent)
//
//	if err != nil{
//		return err
//	}
//
//	if fakeParent.Right == nil {
//		t.Root = nil
//	}
//	return nil
//}
//
////preOrder Traversal
//func (t *Tree) PreOrderTraverse(n *Node, f func(*Node)){
//	if n == nil {
//		return
//	}
//	t.PreOrderTraverse(n.Left,f)
//	f(n)
//	t.PreOrderTraverse(n.Right,f)
//}
//
////preOrder Traversal
//func (t *Tree) InOrderTraverse(n *Node, f func(*Node)){
//	if n == nil {
//		return
//	}
//	f(n)
//	t.InOrderTraverse(n.Left,f)
//	t.InOrderTraverse(n.Right,f)
//}
//
//func (t *Tree) PostOrderTraverse(n *Node, f func(*Node)){
//	if n == nil {
//		return
//	}
//	t.PostOrderTraverse(n.Left,f)
//	t.PostOrderTraverse(n.Right,f)
//	f(n)
//}