package main

import "fmt"


func cs(b suhag,newval string){
b.name=newval
}
type suhag struct {
	name string
	age  int
}

func change_value(n *suhag, newval string) {
	n.name = newval 
}
func pbv ( b suhag , newval string)  suhag{
b.name=newval
return  b
}



func main() {
	s1 := suhag{
		name: "Adhikari",
		age:  19,
	}

	var s2 suhag
	s2.name = "Nirajan"
	s2.age = 20

	fmt.Println(s1.name)
	fmt.Println(s1.age)
	fmt.Println(s2.name)
	fmt.Println(s2.age)

	fmt.Println("\nValue after Change:")
	change_value(&s1,"suhag")

	fmt.Println(s1.name)
cs(s2,"mahim")
fmt.Println(s2.name)
fmt.Println(s2.name)
ayush :=pbv(s2 , "Rohan")
fmt.Println(ayush.name)

}
