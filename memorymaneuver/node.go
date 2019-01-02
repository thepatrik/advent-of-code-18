package memorymaneuver

// Node type
type Node struct {
	Pos              int
	ChildQuantity    int
	MetadataQuantity int
	Len              int
}

// NewNode constructs a new Node
func NewNode(pos, childQuantity, metadataQuantity int) *Node {
	return &Node{Pos: pos, ChildQuantity: childQuantity, MetadataQuantity: metadataQuantity, Len: 2}
}

func (node *Node) sum(nums *[]int) int {
	sum := 0
	for i := 1; i <= node.MetadataQuantity; i++ {
		pos := node.Pos + node.Len - i
		sum += (*nums)[pos]
	}
	return sum
}
