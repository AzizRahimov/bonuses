package main

func main() {

}

func bonus(sales []int) int {
	const priceForBonus = 10_000
	const percentForBonus = 5
	result := 0
	for _, value := range sales {

		if value > priceForBonus {
			result += (value - priceForBonus) * percentForBonus / 100

		}

	}
	return result

}
