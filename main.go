package main

// func main() {
// 	name := "Erfuuan"
// 	age := 25
// 	// fmt.Println(name, "is", age, "years old")
// 	print(name, "  is : ", age, " years old\n")
// }

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

// func printArray(arr *[3]int) {
// 	fmt.Println(*arr)
// }

// func main() {
// 	arr := [3]int{1, 2, 3}
// 	ptr := &arr

// 	fmt.Println(ptr)
// 	fmt.Println(*ptr) // output: [1 2 3]
// 	fmt.Println(ptr)
// 	printArray(ptr) // output: [1 2 3]
// }

//-----------

// func sayHello() {
// 	for i := 1; i <= 5; i++ {
// 		fmt.Println("Hello!")
// 		time.Sleep(500 * time.Millisecond)
// 	}
// }
//----------------
// func main() {
// 	go sayHello()
// 	fmt.Println("Main function is running...")
// 	time.Sleep(3 * time.Second)
// }
//----------------
// func printMessage(msg string) {
// 	for i := 0; i < 9; i++ {
// 		fmt.Println(msg)
// 		time.Sleep(1000 * time.Millisecond)
// 	}
// }

//	func main() {
//		go printMessage("Hello from Goroutine!") // اجرا در پس‌زمینه
//		printMessage("Hello from Main!")         // اجرا در main
//	}
//
// ----------------
//
//	func main() {
//		str := "Hello"
//		bytes := []byte(str)
//		fmt.Println(bytes) //ASCII
//		fmt.Printf("Characters: %s\n", bytes)
//	}
//
// ----------------
//
//	func main() {
//		str := "سلام"
//		runes := []rune(str)
//		fmt.Println("Runes : ", runes)
//		fmt.Printf("Characters: %c\n", runes)
//	}
//
// ----------------
//
//	func main() {
//		str := "سلام"
//		fmt.Println("Byte length:", len(str))         // تعداد بایت‌ها
//		fmt.Println("Rune length:", len([]rune(str))) // تعداد کاراکترها
//	}
//
// ----------------
// func main() {
// 	// آرایه‌ای از byte (فقط مناسب برای حروف انگلیسی)
// 	byteArray := []byte{'H', 'e', 'l', 'l', 'o'}

// 	// آرایه‌ای از rune (مناسب برای تمام کاراکترهای یونیکد)
// 	runeArray := []rune{'س', 'ل', 'ا', 'م'}

// 	fmt.Println("Byte array:", byteArray)
// 	fmt.Println("Rune array:", runeArray)
// }
// ----------------

// func main() {
// 	str := "سلام"

// 	fmt.Println("Using for range:")
// 	for i, r := range str {
// 		fmt.Printf("Index: %d, Rune: %c, Unicode: %U\n", i, r, r)
// 	}

//		fmt.Println("\nUsing byte loop:")
//		for i := 0; i < len(str); i++ {
//			fmt.Printf("Byte %d: %x\n", i, str[i])
//		}
//	}
//
// ----------
// func main() {
// 	str := "Hello"
// 	bytes := []byte(str)    // تبدیل رشته به byte
// 	newStr := string(bytes) // تبدیل مجدد به string

//		fmt.Println("Original:", str)
//		fmt.Println("Bytes:   ", bytes)
//		fmt.Println("New String:", newStr)
//	}
//
// ----------
// func main() {
// 	str := "سلام"
// 	runes := []rune(str)    // تبدیل رشته به rune
// 	newStr := string(runes) // تبدیل مجدد به string

//		fmt.Println("Original:", str)
//		fmt.Println("Runes:", runes)
//		fmt.Println("New String:", newStr)
//	}
//
// ----------
// func main() {
// 	user := map[string]int{
// 		"ali":   12,
// 		"erfan": 20,
// 	}
// 	fmt.Println(user)
// }
// ----------
// func pause() {
// 	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
// }
// func sendMsg(msg string) {
// 	pause()
// 	fmt.Println(msg)
// }
// func main() {
// 	sendMsg("hello")
// 	go sendMsg("t1")
// 	go sendMsg("t2")
// 	go sendMsg("t3")
// 	sendMsg("main")
// 	time.Sleep(1 * time.Second)
// }
// ----------

// func pause() {
// 	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
// }

// func sendMsg(msg string, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	pause()
// 	fmt.Println(msg)
// }

// func main() {
// 	var wg sync.WaitGroup

// 	wg.Add(3)

// 	// go func(msg string) {
// 	// 	defer wg.Done()
// 	// 	pause()
// 	// 	fmt.Println(msg)
// 	// }("test1")

// 	go sendMsg("test1", &wg)
// 	go sendMsg("test2", &wg)
// 	go sendMsg("test3", &wg)

// 	wg.Wait()
// }

// func main() {
// 	// var x int
// 	// const x int = 0
// 	x := 0
// 	fmt.Println(x)
// }

// func main() {
// 	start := time.Now()
// 	var s string

// 	for i := 0; i < 10000; i++ {
// 		s += "hi"
// 	}
// 	fmt.Println(s)
// 	fmt.Println(time.Since(start))

// }

// func main() {
// 	start := time.Now()
// 	var sb strings.Builder

// 	for i := 0; i < 10000; i++ {

// 		sb.WriteString("hi")
// 	}

// 	fmt.Println(sb.String())
// 	fmt.Println(time.Since(start))
// }

// func main() {

// 	var x int8 = math.MaxInt8
// 	fmt.Printf("B=%b , d=%d", x, x)
// }

// func main() {
// 	var nums []int
// 	fmt.Println(len(nums))
// 	fmt.Println(cap(nums))
// }

// func main() {
// 	var sl []int = make([]int, 3, 4)
// 	sl[0] = 1
// 	sl[1] = 1
// 	sl[2] = 2
// 	sl = append(sl, 4, 5, 6)

// 	fmt.Println(sl)
// 	fmt.Println(len(sl), cap(sl))
// }
