// Игра в котрой пользователь должен угадать число
package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)
func  main() {
	// Генератор случайных чисел (15,16)
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	fmt.Println("Я выбрал случайное число от 1 до 100.")
	fmt.Println("Сможешь угадать его ?")
	//Ввод числа с клавиатуры
	reader := bufio.NewReader(os.Stdin)
	//создание цикла для ограничения попыток на решение задачи
	var success bool
	for guesses := 0; guesses < 10; guesses++ {
		fmt.Println("у Вас осталось ", 10-guesses, "попыток.")

		fmt.Print("угадай число: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		// строка 30 - удаления симбвола новой строки
		input = strings.TrimSpace(input)
		// преобразуем строку в число
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}

		if guess < target {
			fmt.Println("Oops. ваше число меньше загадываемого. ")
		} else if guess > target {
			fmt.Println("Oops. ваше число больше загадываемого.  ")
		} else {
			success = true
			fmt.Println("Отличная работа!!! Ты справился.")
			break
		}
	}
	if success != true{
		fmt.Println("Плохо, пробуй еще")
		fmt.Println("Я загадал число :",target)
	}
}

