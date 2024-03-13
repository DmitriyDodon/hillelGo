package main

import (
	"flag"
	"fmt"
)

const defaultRaceDistance = 1000
const defaultTigerSpeed = 70
const defaultTurtleSpeed = 20
const defaultFishSpeed = 50

// Створити структуру Animal з полями runSpeed, name, voice
type Animal struct {
	runSpeed int
	name     string
	voice    string
}

// Для Animal створити метод WinnerSay та LoserSay, вони мають виводити щось на екран
func (a Animal) WinnerSay() {
	fmt.Println("I've won! ", a.voice)
}

func (a Animal) LoserSay() {
	fmt.Println("I've lost!", a.voice)
}

func (a Animal) CheckIfEndInRace(r Race) bool {
	return (r.TimeFromStart * a.runSpeed) >= r.Distance
}

// Створити структури Turtle, Tiger, Fish з принаймні двома унікальними полями (наприклад кількість
// плавників, колір смуг, розмір панциря тощо)
type Turtle struct {
	Animal
	habitat    string
	shellColor string
}

type Tiger struct {
	Animal
	furLengthInMm int
	age           int
}

type Fish struct {
	Animal
	skinType             string
	isLivingInFreshWater bool
}

// Створити структуру Race котра буде мати поля Distance int,  Turtle, Tiger, Fish, метод Start(),
// метод CreateTeam() котрий має приймати і заповнювати поля (Turtle, Tiger, Fish)
type Race struct {
	Distance      int
	Turtle        Turtle
	Tiger         Tiger
	Fish          Fish
	Finished      []RacePerticipant
	TimeFromStart int
}

type RacePerticipant struct {
	Animal
	Kind                          string
	TimeFromStartLeftWhenFinished int
}

func (r Race) checkIfParticipantHasFinished(p RacePerticipant) bool {
	for _, finished := range r.Finished {
		if finished.Kind == p.Kind {
			return true
		}
	}
	return false
}

// Метод Start має запускати "гонку" - в безкінцевому циклі всі тварини мають "бігати"
// дистанцію - 1 ітерація має символізувати 1 одиницю часу за яку тварини пробігають
// дистанцію котра дорівнює їх швидкості (runSpeed) (наприклад у тигра швидкість
// 50 - за одну ітерацію він пробігає 50 метрів з всієї дистанції)
func (r *Race) Start() {
	var racePerticipant RacePerticipant
	r.Finished = make([]RacePerticipant, 0, 3)
	r.TimeFromStart = 0
	for {
		if r.Fish.CheckIfEndInRace(*r) {
			racePerticipant = RacePerticipant{
				Kind: "Риба",
			}
			if !r.checkIfParticipantHasFinished(racePerticipant) {
				racePerticipant.TimeFromStartLeftWhenFinished = r.TimeFromStart
				racePerticipant.Animal = r.Fish.Animal
				r.Finished = append(r.Finished, racePerticipant)
			}
		}
		if r.Turtle.CheckIfEndInRace(*r) {
			racePerticipant = RacePerticipant{
				Kind: "Черепаха",
			}
			if !r.checkIfParticipantHasFinished(racePerticipant) {
				racePerticipant.TimeFromStartLeftWhenFinished = r.TimeFromStart
				racePerticipant.Animal = r.Turtle.Animal
				r.Finished = append(r.Finished, racePerticipant)
			}
		}
		if r.Tiger.CheckIfEndInRace(*r) {
			racePerticipant = RacePerticipant{
				Kind: "Тигр",
			}
			if !r.checkIfParticipantHasFinished(racePerticipant) {
				racePerticipant.TimeFromStartLeftWhenFinished = r.TimeFromStart
				racePerticipant.Animal = r.Tiger.Animal
				r.Finished = append(r.Finished, racePerticipant)
			}
		}
		if len(r.Finished) >= 3 {
			break
		}
		r.TimeFromStart++
	}

	fmt.Printf("%s переміг - час забігу становить %d ітерацій\n", r.Finished[0].Kind, r.Finished[0].TimeFromStartLeftWhenFinished)
	r.Finished[0].WinnerSay()
	fmt.Printf("Другий прийшов %s - час забігу становить %d ітерацій\n", r.Finished[1].Kind, r.Finished[1].TimeFromStartLeftWhenFinished)
	fmt.Printf("Третiм прийшов %s - час забігу становить %d ітерацій\n", r.Finished[2].Kind, r.Finished[2].TimeFromStartLeftWhenFinished)
	r.Finished[2].LoserSay()
}

func NewRace(distance int) *Race {
	race := Race{
		Distance: distance,
	}
	return &race
}

func (r *Race) CreateTeam(turtle Turtle, tiger Tiger, fish Fish) {
	r.Turtle = turtle
	r.Tiger = tiger
	r.Fish = fish
}

func main() {
	var distance, tigerSpeed, turtleSpeed, fishSpeed int

	flag.IntVar(&distance, "distance", defaultRaceDistance, "Set race distance")
	flag.IntVar(&tigerSpeed, "tigerSpeed", defaultTigerSpeed, "Set tiger speed")
	flag.IntVar(&turtleSpeed, "turtleSpeed", defaultTurtleSpeed, "Set turtle speed")
	flag.IntVar(&fishSpeed, "fishSpeed", defaultFishSpeed, "Set fish speed")

	flag.Parse()

	fmt.Println(distance)

	turtle := Turtle{
		Animal: Animal{
			runSpeed: turtleSpeed,
			name:     "Pitt",
			voice:    "ya-ya",
		},
		habitat:    "swamp",
		shellColor: "green",
	}

	tiger := Tiger{
		Animal: Animal{
			runSpeed: tigerSpeed,
			name:     "Chester",
			voice:    "roar",
		},
		furLengthInMm: 20,
		age:           10,
	}

	fish := Fish{
		Animal: Animal{
			runSpeed: fishSpeed,
			name:     "Nemo",
			voice:    "blop blop",
		},
		skinType:             "covered with scales",
		isLivingInFreshWater: false,
	}

	race := NewRace(distance)

	race.CreateTeam(turtle, tiger, fish)

	race.Start()

}
