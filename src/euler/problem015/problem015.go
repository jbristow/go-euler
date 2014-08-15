/*
Starting in the top left corner of a 2×2 grid, and only being able to move to the right and down, there are exactly 6 routes to the bottom right corner.

How many such routes are there through a 20×20 grid?

https://projecteuler.net/problem=15
*/
package euler

func Problem015() int {

	listA := []int{1, 1}
	for i := 2; i <= 40; i++ {
		listB := []int{1}
		for j := 1; j < len(listA); j++ {
			listB = append(listB, listA[j]+listA[j-1])
		}
		listA = append(listB, 1)
	}
	return listA[len(listA)/2]
}
