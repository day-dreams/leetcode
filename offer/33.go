package offer

func verifyPostorder(postorder []int) bool {
	if len(postorder) <= 1 {
		return true
	}

	root := postorder[len(postorder)-1]

	cur := 0

	for cur < len(postorder)-1 {
		if postorder[cur] < root {
			cur++
			continue
		}

		break
	}

	// [0,cur) < root
	mark := cur
	for cur < len(postorder)-1 {
		if postorder[cur] > root {
			cur++
			continue
		}

		break
	}

	// [cur,) > root

	if cur < len(postorder)-1 {
		return false
	}

	if !verifyPostorder(postorder[0:mark]) {
		return false
	}
	return verifyPostorder(postorder[mark : len(postorder)-1])

}
