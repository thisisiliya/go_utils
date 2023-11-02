package split

func ByLength(list *[]string, max_length int) *[][]string {

	result := [][]string{[]string{}}

	for _, i := range *list {

		last_list := &result[len(result)-1]

		if len(*last_list) < max_length {

			*last_list = append(*last_list, i)
		} else {

			result = append(result, []string{})
			*last_list = append(*last_list, i)
		}
	}

	return &result
}
