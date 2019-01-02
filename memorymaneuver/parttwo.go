package memorymaneuver

type sumFunc func(nums *[]int) int

// CalcRootNodeVal calculates the value of the root node.
func CalcRootNodeVal(nums []int) int {
	node := NewNode(0, nums[0], nums[1])
	traverseNode(&nums, node)
	return node.Value
}

func traverseNode(nums *[]int, node *Node) {
	var sumFunc sumFunc = node.Sum

	if node.HasChildren() {
		children := make(map[int]*Node)

		for i := 1; i <= node.ChildQuantity; i++ {
			childPos := node.Pos + node.Len
			child := NewNode(childPos, (*nums)[childPos], (*nums)[childPos+1])
			traverseNode(nums, child)
			node.Len += child.Len
			children[i] = child
		}

		// Switch the summation func to be based on referenced child node values
		sumFunc = func(n *[]int) int {
			sum := 0
			for _, ref := range node.ChildRefs(n) {
				node, ok := children[ref]
				if ok {
					sum += node.Value
				}
			}
			return sum
		}
	}

	node.Len += node.MetadataQuantity
	node.Value = sumFunc(nums)
}
