package memorymaneuver

// CalcMetadataSum calculates the sum of all metadata
func CalcMetadataSum(nums []int) int {
	node := NewNode(0, nums[0], nums[1])
	return metadataSum(&nums, node)
}

func metadataSum(nums *[]int, node *Node) int {
	sum := 0
	for i := 0; i < node.ChildQuantity; i++ {
		childPos := node.Pos + node.Len
		child := NewNode(childPos, (*nums)[childPos], (*nums)[childPos+1])
		sum += metadataSum(nums, child)
		node.Len += child.Len
	}

	// Append the metadata quantity to get the complete length
	node.Len += node.MetadataQuantity

	// When we know the length we can finally calculate the sum
	return sum + node.sum(nums)
}
