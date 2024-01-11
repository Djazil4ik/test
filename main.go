package main

import "fmt"

func main() {
	//Часть Mamc
	var n int
	fmt.Print("Введите число возрастных групп:\n")
	fmt.Scan(&n)
	var array [][2]int
	var n1, n2 int
	for i := 0; i < n; i++ {
		fmt.Printf("Введите возрастную группу №%d:\n", i+1)
		fmt.Println("Min:")
		fmt.Scan(&n1)
		fmt.Println("Max:")
		fmt.Scan(&n2)
		array = append(array, [2]int{n1, n2})
	}
	fmt.Println(array)
	//Часть респондента
	var old int
	fmt.Println("Введите возраст респондента:")
	fmt.Scan(&old)
	for i := 1; i < n+1; i++ {

		if old >= array[i-1][0] && old <= array[i-1][1] {
			fmt.Println(i)
			break
		}
		if i == n {
			fmt.Println(98)
		}
	}
}
