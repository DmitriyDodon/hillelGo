package main

import "fmt"

// 1 - Створити структуру Game, в якій буде 5 полей. Одне з яких TotalPlayers з типом int. Інші будь які.
type Game struct {
	TotalPlayers int
	Title        string
	Genre        string
	FieldType    string
	Inventer     string
}

// 2 - Створити структуру TableGame  - де буде вбудована структура Game та ще одне будь яке поле.
type TableGame struct {
	Game
	FigureCount int
}

func main() {

	// 3 - Створити обʼєкт та заповнити всі поля.
	chess := TableGame{
		Game: Game{
			TotalPlayers: 2,
			Title:        "chess",
			Genre:        "educational",
			FieldType:    "board of 64 squares arranged",
			Inventer:     "unknown",
		},
		FigureCount: 32,
	}

	// 4 - Створити слайс з типом TableGame додати туди обʼєкт з 3 (третього) пункту.
	tableGames := make([]TableGame, 0, 3)
	oldCappacity := cap(tableGames)
	tableGames = append(tableGames, chess)

	// 5 - Створити ще 2 обʼєкти і також додати їх до слайсу.
	checkers := TableGame{
		Game: Game{
			TotalPlayers: 2,
			Title:        "checkers",
			Genre:        "educational",
			FieldType:    "board of 64 squares arranged",
			Inventer:     "Frenchman",
		},
		FigureCount: 24,
	}

	pocker := TableGame{
		Game: Game{
			TotalPlayers: 10,
			Title:        "pocker",
			Genre:        "gambling",
			FieldType:    "table",
			Inventer:     "unknown",
		},
		FigureCount: 52,
	}

	tableGames = append(tableGames, checkers, pocker)

	// 6 - Переконатись що при виконанні 5-го пункта в слайсі не буде виникати аллокація памʼяті (якщо буде, то змінити код так щоб аллокація не виникала)
	cappacityDifference := oldCappacity - cap(tableGames)
	if cappacityDifference == 0 {
		fmt.Printf("Різниця в capacity %d. Отже аллокацiя пам'ятi не вiдбулася.\n\n", cappacityDifference)
	} else {
		fmt.Printf("Різниця в capacity %d. Отже аллокацiя пам'ятi вiдбулася.\n\n", cappacityDifference)
	}

	// 7 - Всі елементи зі слайсу записати в мапу map[int]TableGame - де ключом буде індекс елементу зі слайса. За допогою for{}
	tableGamesMap := make(map[int]TableGame)

	for i, v := range tableGames {
		tableGamesMap[i] = v
	}

	// 8 - вивести на екран мапу (кожне значення з нового рядка)
	// 9 - За допомогою циклу знайти сумму всіх TotalPlayers мапі (в усіх елементах)
	totalPlayersSum := 0
	for _, v := range tableGamesMap {
		fmt.Printf("%+v\n", v)
		totalPlayersSum += v.TotalPlayers
	}

	// 10 - Сумму вивести на екран
	fmt.Printf("\nTotalPlayersSum: %d\n", totalPlayersSum)
}
