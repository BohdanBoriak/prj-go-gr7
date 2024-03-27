package main

import (
	"fmt"
	"math/rand"
	"prj-go/domain"
	"strconv"
	"time"
)

const (
	totalPoints       = 50
	pointsPerQuestion = 10
)

var id uint64 = 1

func main() {
	fmt.Println("Вітаємо у грі MATHREVENGE!")
	time.Sleep(1 * time.Second)

}

func play() domain.User {
	for i := 5; i > 0; i-- {
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}

	myPoints := 0
	start := time.Now()
	for myPoints < totalPoints {
		x, y := rand.Intn(100), rand.Intn(100)
		fmt.Printf("%v + %v = ", x, y)

		ans := ""
		fmt.Scan(&ans)

		ansInt, err := strconv.Atoi(ans)
		if err != nil {
			fmt.Println("Спробуй ще!")
		} else {
			if ansInt == x+y {
				myPoints += pointsPerQuestion
				points := totalPoints - myPoints
				fmt.Println("Правильно, ти набрав", myPoints, "очок!")
				fmt.Printf("Залишилось набрати %v!\n", points)
			} else {
				fmt.Println("Спробуй ще!")
			}
		}
	}
	end := time.Now()
	timeSpent := end.Sub(start)

	fmt.Printf("Молодчинка, впорався всього за %v!", timeSpent)
	fmt.Println("Введіть своє ім'я:")

	name := ""
	fmt.Scan(&name)

	// var user domain.User
	// user.Id = id
	// user.Name = name
	// user.Time = timeSpent

	user := domain.User{
		Id:   id,
		Name: name,
		Time: timeSpent,
	}
	id++

	return user
}
