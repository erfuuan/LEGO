
package main
import "fmt"

// func main() {
//      name := "Erfuuan"
//      age := 25
//     fmt.Println(name, "is", age, "years old")
// }
//?=====================================================================
// func main() {
// 	if age==18 {
// 		fmt.Println("bache")
// 	}else{
// 		fmt.Println("bozorg")
// 	}
// }
//?=====================================================================

// func main() {
// 	for i:=1; i<=5; i++ {
// 		fmt.Println("level", i)	
// 	}
// }

// func main(){
// 	for {
// 		fmt.Println("infinity")
// 	}
// }
//?=====================================================================
// func greet(name string){fmt.Println("hi :", name)}
// func add (a int , b int) int{return a+b}

// func main(){
// 	greet("erfan")
// 	result:=add(3,4)
// 	fmt.Println(result)
// }
//?=====================================================================
// type Person struct {
// 	Name string 
// 	Age int
// }
// func main(){
//     p := Person{Name: "Erfuuan", Age: 25}
// 	fmt.Println(p)
// }
//?=====================================================================
// func sayHi (){
// 	fmt.Println("hi")
// }

// func main(){
// 	go sayHi()
// 	fmt.Println("goodBye")
// }
//?=====================================================================
//array
// func main(){
// 	var numbers[5] int
// 	numbers[0]=0
// 	numbers[1]=1
// 	numbers[2]=2
// 	numbers[3]=3
// 	numbers[4]=4
// 	fmt.Println(numbers[1])
// }
//?=====================================================================
//Slices
// func main(){
// 	var fruits []string 
// 	fruits=append(fruits,"benana")
// 	fruits=append(fruits,"apple")
// 	fmt.Println(fruits[0])
// 	fmt.Println(fruits)
// }
//?=====================================================================
//Maps
func main(){
	var score map[string]int
	score=make(map[string]int)
	score["Erfan"]=1
	score ["ali"]=2
	score ["sara"]=3
	fmt.Println(score["Erfan"])
}