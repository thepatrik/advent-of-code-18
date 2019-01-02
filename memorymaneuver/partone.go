package memorymaneuver

// CalcMetadataSum calculates the sum of all metadata
func CalcMetadataSum(nums []int) int {
	node := NewNode(0, nums[0], nums[1])
	traverse(&nums, node)
	return node.Value
}

func traverse(nums *[]int, node *Node) {
	for i := 0; i < node.ChildQuantity; i++ {
		childPos := node.Pos + node.Len
		child := NewNode(childPos, (*nums)[childPos], (*nums)[childPos+1])
		traverse(nums, child)
		node.Value += child.Value
		node.Len += child.Len
	}

	// Append the metadata quantity to get the complete length
	node.Len += node.MetadataQuantity

	// When we know the length we can finally calculate the value
	node.Value += node.Sum(nums)
}
