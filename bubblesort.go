package datstr

func bubbleSort(arr []int) []int {
	end := len(arr)
	for i := 0; i < len(arr); i++ {
		for j := 0; j < end-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
		end--
	}
	return arr
}
