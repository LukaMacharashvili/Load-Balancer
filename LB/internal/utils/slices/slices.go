package slices

func GetIndexOfElementInSliceString(slice []string, element string) int {
	s := slice
	for i, v := range s {
		if v == element {
			return i
		}
	}

	return -1
}

func RemoveElementString(slice *[]string, element string) {
	idxOfTargetUrl := GetIndexOfElementInSliceString(*slice, element)

	if idxOfTargetUrl != -1 {
		(*slice) = append((*slice)[:idxOfTargetUrl], (*slice)[idxOfTargetUrl+1:]...)
	}
}

func AddTargetUrlToEnv(slice *[]string, element string) {
	if GetIndexOfElementInSliceString(*slice, element) != -1 {
		*slice = append(*slice, element)
	}
}
