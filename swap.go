package main
import "fmt"
func pani(a *int){
if *a=nil{
fmt.Println("The pointer must be empty")
}else{
fmt.Println(a)
}
}
func Namaskar(*name *string){
*name="HEllo User" + *name
}
func tooglebool(a *bool){
*a=!*a

}
func reset_to_zero(a *int){
*a=0
}
func swap(a, b *int) {
temp:=*a
*a = *b
*b = temp
}

func main() {
x :=10
y :=11
fmt.Println("x =",x)
fmt.Println("y =",y)
swap(&x,&y)
fmt.Println("After swap",swap)
fmt.Println("x =",x)
fmt.Println("y =",y)
z:=10
fmt.Println("z =",z)
reset_to_zero(&z)
fmt.Println("z =",z)
d:=true
fmt.Println("Boolean =" ,d)
tooglebool(&d)
fmt.Print("Boolen =",d)
name :="suhag"
Namaskar(&name)
fmt.Println(name)
g:=0
pani(&g)
p=4
pani(&p)
fmt.Println(g)
fmt.Println(p)
}


