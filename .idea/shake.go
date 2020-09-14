package main

import "fmt"

func main() {
	a := [10]int{10,9,8,7,6,5,4,3,2,1}
	flag:=true
	fmt.Println(a)
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
		for k:=len(a)-1;k>0;k--{
			if a[k]<a[k-1]{
				var n = a[k-1]
				a[k-1]=a[k]
				a[k]=n
				flag=true

			}
		}
	}

	fmt.Println(a)
}