package memorymaneuver

// Node type
type Node struct {
	Pos              int
	ChildQuantity    int
	MetadataQuantity int
	Value            int
	Len              int
}

// NewNode constructs a new Node
func NewNode(pos, childQuantity, metadataQuantity int) *Node {
	return &Node{Pos: pos, ChildQuantity: childQuantity, MetadataQuantity: metadataQuantity, Len: 2}
}

// HasChildren checks if the node has any children or not
func (node *Node) HasChildren() bool {
	return node.ChildQuantity > 0
}

// ChildRefs returns slice with child ref positions
func (node *Node) ChildRefs(nums *[]int) []int {
	refs := make([]int, 0)
	for i := 1; i <= node.MetadataQuantity; i++ {
		pos := node.Pos + node.Len - i
		refs = append(refs, (*nums)[pos])
	}
	return refs
}

// Sum calculates the sum of all metadata values
func (node *Node) Sum(nums *[]int) int {
	sum := 0
	for i := 1; i <= node.MetadataQuantity; i++ {
		pos := node.Pos + node.Len - i
		sum += (*nums)[pos]
	}
	return sum
}
