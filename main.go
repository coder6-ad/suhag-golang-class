package main

import "fmt"
func main() {
user:=map[int]string{
6:"Deepa",
14:"Rajib",

}
user[6]="Mahim"
_, okaygo:=user[6]
fmt.Println(okaygo)
fmt.Println(len(user))
delete(user,6)
}
