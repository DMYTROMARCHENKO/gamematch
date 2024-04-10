package main

import (
	"fmt"
	"math/rand"
	"time"
)

func welcomePlayer() {
	fmt.Println("Ласкаво просимо до математичної гри!")
	fmt.Println("Давайте подивимося, як швидко ви досягнете 50 балів.")
	fmt.Println("Готуйся...")
}

func countdown() {
	for i := 5; i > 0; i-- {
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}
}

func generateQuestion() (int, int) {
	source := rand.NewSource(time.Now().UnixNano())
	rand := rand.New(source)
	num1 := rand.Intn(10) + 1
	num2 := rand.Intn(10) + 1
	return num1, num2
}

func calculateAnswer(num1, num2 int) int {
	return num1 + num2
}

func mainGame() {
	score := 0
	startTime := time.Now()

	for score < 50 {
		num1, num2 := generateQuestion()
		answer := calculateAnswer(num1, num2)

		fmt.Printf("Скільки буде, якщо %d + %d?\n", num1, num2)
		var userAnswer int
		fmt.Print("Ваша відповідь: ")
		fmt.Scan(&userAnswer)

		if userAnswer == answer {
			score += 5
			fmt.Printf("Правильно! Ваш рахунок зараз %d. Тобі потрібно %d більше очок для перемоги.\n", score, 50-score)
		} else {
			fmt.Println("Невірно! Спробуйте знову.")
		}
	}

	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)
	fmt.Printf("Щиро вітаю! Ви набрали 50 балів %.2f секунд.\n", elapsedTime.Seconds())
}

func main() {
	welcomePlayer()
	countdown()
	mainGame()
}
