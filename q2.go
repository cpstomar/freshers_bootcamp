package main
import "fmt"
type Node struct {
	val string
	left *Node
	right *Node
}
func preorder(root *Node){
	if root!=nil {
		fmt.Printf("%s ", root.val)
	}

	if root.left != nil {
		preorder(root.left)
	}
	if root.right != nil {
		preorder(root.right)
	}
}
func postorder(root *Node){
	if root.left != nil {
		postorder(root.left)
	}
	if root.right != nil {
		postorder(root.right)
	}
	if root!=nil {
		fmt.Printf("%s ", root.val)
	}

}
func main(){
	tree := [5]Node{}
	tree[0].val = "+"
	tree[1].val = "a"
	tree[2].val = "-"
	tree[0].left = &tree[1]
	tree[0].right = &tree[2]
	tree[3].val = "b"
	tree[4].val = "c"
	tree[2].left = &tree[3]
	tree[2].right = &tree[4]
	preorder(&tree[0])
	fmt.Println("")
	postorder(&tree[0])
}
