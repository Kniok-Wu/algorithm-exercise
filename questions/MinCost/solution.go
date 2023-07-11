package main

import "fmt"

func Sort(shops [][]int) {
	if len(shops) <= 1 {
		return
	}

	pivot := shops[0]
	i, j := 0, len(shops)-1

	for i < j {
		for shops[j][1] > pivot[1] && j > i {
			j--
		}

		if i < j {
			shops[i] = shops[j]
		}

		for shops[i][1] < pivot[1] && j > i {
			i++
		}

		if i < j {
			shops[j] = shops[i]
		}

	}

	shops[j] = pivot

	Sort(shops[:j])
	Sort(shops[j+1:])
}

func main() {
	var numShop, toBuy int
	shops := make([][]int, 0)
	cost := 0
	fmt.Scanf("%d %d", &numShop, &toBuy)

	for i := 0; i < numShop; i++ {
		shops = append(shops, make([]int, 2))
		fmt.Scanf("%d %d", &shops[i][0], &shops[i][1])
	}

	Sort(shops)

	for i := 0; i < len(shops); i++ {
		if toBuy > shops[i][0] {
			cost += shops[i][0] * shops[i][1]
			toBuy -= shops[i][0]
		} else {
			cost += toBuy * shops[i][1]
			break
		}
	}

	fmt.Println(cost)
}
