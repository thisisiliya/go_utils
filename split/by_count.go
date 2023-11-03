package split

func ByCount(list []string, count int) *[][]string {

	result := [][]string{}

	switch len(list) {

	case 0:
		return &[][]string{}

	case 1:
		return &[][]string{{list[0]}}

	case count:
		for _, v := range list {

			result = append(result, []string{v})
		}

		return &result

	default:
		min_index := 0
		max_index := len(list) / count

		if len(list)%count != 0 {

			max_index++
		}

		for i := 0; i < count; i++ {

			result = append(result, []string{})
			result[i] = list[min_index:max_index]
			min_index, max_index = max_index, max_index+len(list)/count

		}

		return &result
	}
}
