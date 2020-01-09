package main

func main() {
	sales := []int {12_000,8_000,15_000,8_000}
	sum:=0

	for _, value := range sales{
	sum += bonus(value)


	}

}
func  bonus(money int) int{
	const priceForBonus = 10_000
	const percentForBonus = 5
	result := 0
	if money >= priceForBonus{
		result = (money - priceForBonus) * percentForBonus /100
	}

return result
}