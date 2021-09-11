package lists

import "fmt"

func main() {
	var productNames [4]string = [4]string{"Attack Go"}
	productNames[3] = "Mask"

	prices := [4]float64{10.99, 6.99, 420.69, 1337.00}

	fmt.Println(prices)
	fmt.Println(productNames)
	fmt.Println(prices[2])

	featuredPrices := prices[:3]
	fmt.Println(featuredPrices)

	//In this case you do not have to specify the amount of elements in the array. Go will automatically update the array.
	dynamicPrices := []float64{}
	fmt.Println(dynamicPrices)
	// Even though it is dynamic you cannot change/add values using array[index] for an index out of bounds instead use append()
	dynamicPrices = append(dynamicPrices, 9.99)
	fmt.Println(dynamicPrices)
}
