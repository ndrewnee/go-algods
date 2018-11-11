package exercises

func Height(tree []int) int {
	treeMap := treeToMap(tree)
	root := treeMap[-1][0]
	return calcHeight(treeMap, root)
}

func treeToMap(tree []int) map[int][]int {
	treeMap := make(map[int][]int)
	for value, parent := range tree {
		treeMap[parent] = append(treeMap[parent], value)
	}

	return treeMap
}

func calcHeight(treeMap map[int][]int, root int) int {
	height := 1
	children := treeMap[root]
	for _, child := range children {
		childHeight := calcHeight(treeMap, child)
		height = maxOf(height, childHeight+1)
	}

	return height
}
