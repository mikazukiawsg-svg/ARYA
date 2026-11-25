package main

import "fmt"

func main() {
nilai := []int{70, 85, 90, 65, 88}

for i := 0; i < len(nilai); i++{
	fmt.Println(i, nilai[i])
}
for i, v := range nilai {
	fmt.Println(i, v)
}
for _, v := range nilai {
	fmt.Println(v)
}

}