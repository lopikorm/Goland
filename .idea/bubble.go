package main

import "fmt"

func main() {
	a := [10]int{10,5,8,7,9,5,3,3,2,1}
	fmt.Println(a)

	flag:=true

	for flag{
		flag=false

		for i:=0;i<len(a)-1;i++ {

			if a[i]>a[i+1]{
				var j = a[i]
				a[i]=a[i+1]
				a[i+1]=j
				flag=true
			}
		}
	}
	fmt.Println(a)
}
