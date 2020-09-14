
package main

import "fmt"

func main() {
	a := [10]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	fmt.Println(a)
	var k = 0
	var n = 1
	for n < len(a)-1 {
		if a[n] > a[k] {
			k = n
			n = n + 1
		}

		if a[n] < a[k] {
			var i = a[k]
			a[k] = a[n]
			a[n] = i
			n = k
			k = k - 1

			if k == -1 {
				k = n
				n = k + 1
			}

		}
	}

	fmt.Println(a)
}
