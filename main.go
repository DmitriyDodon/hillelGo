package main

import "fmt"

func main() {
	fmt.Println("1 - Зробити конкатинацію рядків та помістити результат у змінну")
	var resString string = "Hello, " + "world!"
	fmt.Println(resString)
	devideFunc()

	fmt.Println("2 - Зробити змінну результатом якої буде форматування рядків за допомогою fmt.Sprintf(...), використовуючи %s та %d, значення змінної вивести на екран")
	resString2 := fmt.Sprintf("%s, %d", resString, 10)
	fmt.Println(resString2)
	devideFunc()

	fmt.Println("3 - Зробити математичні операції: +, -, *, /, з будь якими числами результати помістити в змінні та вивести на екран")
	res := float64(3) + float64(45) - float64(2.33)*float64(5.55)/float64(14.42)
	fmt.Print(res, "\n")
	devideFunc()

	fmt.Println("4 - Створити дві змінні типу float з різними значеннями і написати 4 умовні конструкції if використовуючи оператори: >, >=, <, <=, ==, !=, &&, || також додати else до двух конструкцій, в тілі конструкцій виводити на екран ті операціїї які біли в умові")
	firstFloat, secondFloat := 34.44, 45.55

	if firstFloat > secondFloat {
		fmt.Printf("%.2f > %.2f\n", firstFloat, secondFloat)
	} else {
		fmt.Printf("%.2f < %.2f\n", firstFloat, secondFloat)
	}

	if firstFloat < secondFloat {
		fmt.Printf("%.2f < %.2f\n", firstFloat, secondFloat)
	} else {
		fmt.Printf("%.2f > %.2f\n", firstFloat, secondFloat)
	}

	if firstFloat != secondFloat {
		fmt.Printf("%.2f != %.2f\n", firstFloat, secondFloat)
	}

	if firstFloat == secondFloat {
		fmt.Printf("%.2f == %.2f\n", firstFloat, secondFloat)
	}
	devideFunc()

	fmt.Println("5 - Створити змінну типу string та за допомогою switch case і будь яких умов вивести ії значення на екран (має бути не менше трьох case)")
	var stringTypeVar string = "Hello, world!"
	switch stringTypeVar {
	case "string1":
		fmt.Println("string1")
	case "Hello, world!":
		fmt.Print("Hello, world!", "\n")
	case "string2":
		fmt.Printf("%s\n", "string2")
	case "string3":
		fmt.Printf("%s\n", "string3")
	}
	devideFunc()

	fmt.Println("6 - Створити ще один switch case, але без параметру switch результатом має виводитись надписи з двох будь яких case стейтментів")
	switch {
	case false:
		fmt.Println("False value")
	case true && false:
		fmt.Println("True and false value")
	case true || false:
		fmt.Println("True or false value")
		fallthrough
	case true:
		fmt.Println("True value")
	}
	devideFunc()

	fmt.Println("7 - Створити switch case з викроистанням дефолтного стейтменту, зробити вивід на екран будь чого з дефолтної частини блока")
	switch {
	case true && false:
		fmt.Println("True and false value")
	case false:
		fmt.Println("False value")
	default:
		fmt.Print("Default value\n")
	}
	devideFunc()

	fmt.Println("8 - Створити масив чисел з довжиною 5, інкрементувати елемент під індексом 3, декрементувати елемент під індиксом 4 - результат вивести на екран")
	arr := [5]int{}
	arr[3]++
	arr[4]--
	fmt.Printf("%#v\n", arr)
	devideFunc()

	fmt.Println("9 - Створити два слайси типа string двума способами (з використанням make та без використання make), додати в перший слайс рядок з пункту 1 та вивести на екран, також вивести кількість елементів з другого слайсу другим пареметром")
	firstSlice := []string{}
	secondSlice := make([]string, 2)

	firstSlice = append(firstSlice, resString)
	fmt.Println(firstSlice, len(secondSlice))

}

func devideFunc() {
	fmt.Print("\n\n")
}
