package groupietracker

/*------------------------------------------------------*/
/*					Alphabet Sort :						*/
/*------------------------------------------------------*/
func alphabetSort(list Artists, left int, right int) int {
	pivot := list[right]
	i := left - 1
	for j := left; j < right; j++ {
		if list[j].Name < pivot.Name {
			i++
			list[i], list[j] = list[j], list[i]
		}
	}
	list[i+1], list[right] = list[right], list[i+1]
	return i + 1
}

/*------------------------------------------------------*/
/*						Date Sort :						*/
/*------------------------------------------------------*/
func datesSort(list Artists, left int, right int) int {
	pivot := list[right]
	i := left - 1
	for j := left; j < right; j++ {
		if list[j].CreationDate < pivot.CreationDate {
			i++
			list[i], list[j] = list[j], list[i]
		}
	}
	list[i+1], list[right] = list[right], list[i+1]
	return i + 1
}

/*------------------------------------------------------*/
/*					Members Sort :						*/
/*------------------------------------------------------*/
func membersSort(list Artists, left int, right int) int {
	pivot := list[right]
	i := left - 1
	for j := left; j < right; j++ {
		if len(list[j].Members) < len(pivot.Members) {
			i++
			list[i], list[j] = list[j], list[i]
		}
	}
	list[i+1], list[right] = list[right], list[i+1]
	return i + 1
}

/*------------------------------------------------------*/
/*					Partitions Sort :					*/
/*------------------------------------------------------*/
func partitionsSort(list Artists, leftIndex int, rightIndex int, nameSort string) Artists {
	var pivotIndex int // number to determine where to split the array
	if leftIndex < rightIndex {
		switch nameSort {
		case "alphabet":
			pivotIndex = alphabetSort(list, leftIndex, rightIndex)
		case "dates":
			pivotIndex = datesSort(list, leftIndex, rightIndex)
		case "members":
			pivotIndex = membersSort(list, leftIndex, rightIndex)
		}
		partitionsSort(list, leftIndex, pivotIndex-1, nameSort)  // sort left side of the array
		partitionsSort(list, pivotIndex+1, rightIndex, nameSort) // sort right side of the array
	}
	return list
}

/*------------------------------------------------------*/
/*					Sort Selection :					*/
/*------------------------------------------------------*/
func GetSort(sortName string, data Artists) Artists {
	switch sortName {
	case "alphabet":
		return partitionsSort(data, 0, len(data)-1, "alphabet")
	case "date":
		return partitionsSort(data, 0, len(data)-1, "dates")
	case "members":
		return partitionsSort(data, 0, len(data)-1, "members")
	}
	return nil
}
