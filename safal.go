///package main
//import"fmt"
//var name string="suhag"
//var nirajan string
////var age int
////var temp float64
////var choice bool
////func main(){
////name :="suha"
//
////fmt.Println(name,nirajan,age,temp,choice)
/////}
//
//package main
//import"fmt"
//
//type rectangle struct{
//length int
//breath int
//}
//func (rect*rectangle) perimeter(){
//return 2*(rect.length+rect.breath)
//}
//func(rect*rectangle) area(){
//return rect.length*rect.breath
//func main(){
//var rect rectangle
//fmt.Println("Enter the value of breath and length")
//fmt.Scanln(&rect.length,&rect.breath)
//area:=rect.length* rect.breath
//perimeter:=2*(rect.length+rect.breath)
//fmt.Println("\nThe area  of rectangle is ",area)
//fmt.Println("\nTHe perimeter of rectangle is ",perimeter)
//fmt.Println("\n perimeter=",rect.perimeter)
//fmt.Println("\n area=",rect.area)
//}

package main
import "fmt"

func suhag (l,b float64)(float64,float64){
area:=l*b
perimeter:=2*(l+b)
return area,perimeter

}
func main(){
l:=10.5
b:=11.4
area,_:=suhag(l,b)
fmt.Println("The area of the rectangle is",area)
fmt.Print("The perimeter of the rectangle is",)
}
