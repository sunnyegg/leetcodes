package mergetwosortedlists

type ListNode struct {
	Val  int
	Next *ListNode
}

type Head struct {
	head *ListNode
}

type Merged struct {
	merged []int
}

func Merge(list1 *ListNode, list2 *ListNode) *ListNode {
	output := &Head{}
	var values1 []int
	var values2 []int
	var merged Merged

	// jadiin slice int biar lebih mudah untuk sort
	for list1 != nil {
		values1 = append(values1, list1.Val)
		list1 = list1.Next
	}
	for list2 != nil {
		values2 = append(values2, list2.Val)
		list2 = list2.Next
	}

	// create linked list
	merged.mergeSort(values1, values2)
	for _, v := range merged.merged {
		output.add(v)
	}

	return output.head
}

func (l *Head) add(val int) {
	node := &ListNode{
		Val: val,
	}

	if l.head == nil {
		l.head = node
		return
	}

	curr := l.head
	for curr.Next != nil {
		curr = curr.Next
	}
	curr.Next = node
}

func (m *Merged) mergeSort(slice1, slice2 []int) {
	// jadiin satu slice
	for _, v := range slice1 {
		m.merged = append(m.merged, v)
	}
	for _, v := range slice2 {
		m.merged = append(m.merged, v)
	}

	// sort
	for i := 0; i < len(m.merged); i++ {
		pivot := i // yg palng kiri sebagai pivot

		// iterasi i+1 buat bandingin sama i
		for j := i + 1; j < len(m.merged); j++ {
			// kalo pivot lebih besar, berarti j yg kecil (swap)
			if m.merged[pivot] > m.merged[j] {
				pivot = j
			}
		}

		// kalo pivot itu udah paling kecil, lanjut iterasi
		if pivot == i {
			continue
		}

		// swap
		m.merged[i], m.merged[pivot] = m.merged[pivot], m.merged[i]
	}
}
