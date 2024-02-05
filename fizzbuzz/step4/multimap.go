package step4

func multiFizzBuzzMap(numbers []int) map[int]string {
	result := make(map[int]string, len(numbers))

	for _, number := range numbers {
		result[number] = fizzbuzz(number)
	}
	return result
}
