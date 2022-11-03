package main

func main() {

}

func MinMax(arr []int) [2]int {

	min, max := arr[0], arr[0]

	for _, v := range arr[1:] {
		if v < min {
			min = v
		} else if v > max {
			max = v
		}
	}

	return [2]int{min, max}
}
