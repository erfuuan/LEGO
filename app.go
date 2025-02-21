package main

//	func main() {
//	     name := "Erfuuan"
//	     age := 25
//	    fmt.Println(name, "is", age, "years old")
//	}
//
// ?=====================================================================
//
//	func main() {
//		if age==18 {
//			fmt.Println("bache")
//		}else{
//			fmt.Println("bozorg")
//		}
//	}
//
// ?=====================================================================
// * for loop
//
//	func main() {
//		for i := 1; i <= 5; i++ {
//			fmt.Println("level", i)
//		}
//	}
//
// ------------
//
//	func main() {
//		for {
//			fmt.Println("infinity")
//		}
//	}
//
// ------------
//
//	func main() {
//		for {
//			fmt.Println("infinity")
//			break
//		}
//	}
//
// ------------
//
//	func main() {
//		num := 0
//		for num < 5 {
//			fmt.Println("num is:", num)
//			num++
//		}
//	}
//
// -------------
//
//	func main() {
//		nums := []int{1, 2, 3, 4, 5, 6, 7}
//		for index, value := range nums {
//			fmt.Println("Index:", index, "Value:", value)
//		}
//	}
//
// -------------
//
//	func main() {
//		nums := []int{1, 2, 3, 4, 5, 6, 7}
//		for _, value := range nums {
//			fmt.Println("Value:", value)
//		}
//	}
// -------------
//	func main(){
//		j := 0
//		for j < 5 {
//			fmt.Println(j)
//			j++
//		}
//	}
//
// -------------
// func main() {
// 	ages := map[string]int{"ali": 2, "sara": 3}
// 	for key, value := range ages {
// 		fmt.Println(key, value)
// 	}
// }
// -------------
//	func main() {
//		for i := 1; i <= 10; i++ {
//			if i == 5 {
//				fmt.Println("break")
//				break
//			}
//			fmt.Println((i))
//		}
//	}
//
// -------------
// func main() {
// 	for i := 1; i <= 5; i++ {
// 		if i == 3 {
// 			fmt.Println(("continue"))
// 			continue
// 		}
// 		fmt.Println(i)
// 	}
// }
//?=====================================================================
// func greet(name string){fmt.Println("hi :", name)}
// func add (a int , b int) int{return a+b}

//	func main(){
//		greet("erfan")
//		result:=add(3,4)
//		fmt.Println(result)
//	}
//
// -------------
// Recursive Function :
//
//	func factorial(n int) int {
//		if n == 0 {
//			return 1
//		}
//		return n * factorial(n-1)
//	}
//
//	func main() {
//		fmt.Println(factorial(5))
//	}
//
// -------------
// variadic function
//
//	func sum(nums ...int) int {
//		total := 0
//		for _, num := range nums {
//			total += num
//		}
//		return total
//	}
//
//	func main() {
//		fmt.Println(sum(1, 2, 3, 4, 5))
//	}
//
// -------------
//First-Class Function
// func main() {
// 	a := func(a, b int) int {
// 		return a + b
// 	}
// 	fmt.Println(a(1, 2))
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

//	func main(){
//		go sayHi()
//		fmt.Println("goodBye")
//	}
//
// ?=====================================================================
// array
//
//	func main(){
//		var numbers[5] int
//		numbers[0]=0
//		numbers[1]=1
//		numbers[2]=2
//		numbers[3]=3
//		numbers[4]=4
//		fmt.Println(numbers[1])
//	}
//
// ?=====================================================================
// Slices
//
//	func main(){
//		var fruits []string
//		fruits=append(fruits,"benana")
//		fruits=append(fruits,"apple")
//		fmt.Println(fruits[0])
//		fmt.Println(fruits)
//	}
//
// ?=====================================================================
// Maps
//
//	func main(){
//		var score map[string]int
//		score=make(map[string]int)
//		score["Erfan"]=1
//		score ["ali"]=2
//		score ["sara"]=3
//		fmt.Println(score["Erfan"])
//	}
//
// ?=====================================================================
// func main(){
// switch/case
// switch day := "Monday"; day {
// case "Monday":
//
//	fmt.Println("Start of the week")
//
// case "Friday":
//
//	fmt.Println("End of the week")
//
// default:
//
//	    fmt.Println("Midweek")
//	}
//
// }
// ?=====================================================================
//* Pointer
// func main() {
// 	var p *int
// 	fmt.Println(p)
// }

// func main() {
// 	name := "erfuuan"
// 	ptrName := &name
// 	fmt.Println((ptrName))
// 	fmt.Println((*ptrName))
// }

// func main() {
// 	name := "erfan"
// 	ptrName := &name
// 	*ptrName = "ali"
// 	fmt.Println((name))
// }

// func changeValue(x int) {
// 	x = 20
// }

// func main() {
// 	num := 10
// 	changeValue(num)
// 	fmt.Println(num)
// }

// func changeValue(x *int) {
// 	*x = 20
// }
// func main() {
// 	num := 10
// 	changeValue(&num)
// 	fmt.Println(num)
// }
