package split

func ByCount(list *[]string, count int) *[][]string {

	data := *list

	result := [][]string{}

	min_index := 0
	max_index := len(data)/count + 1

	for i := 0; i < count; i++ {

		result = append(result, []string{})
		result[i] = data[min_index:max_index]
		min_index, max_index = max_index, max_index+len(data)/count

	}

	return &result
}
