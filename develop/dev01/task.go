package main

import (
	"fmt"
	"log"
	"time"

	"github.com/beevik/ntp"
)

/*
=== Базовая задача ===
Создать программу печатающую точное время с использованием NTP библиотеки.Инициализировать как go module.
Использовать библиотеку https://github.com/beevik/ntp.
Написать программу печатающую текущее время / точное время с использованием этой библиотеки.
Программа должна быть оформлена с использованием как go module.
Программа должна корректно обрабатывать ошибки библиотеки: распечатывать их в STDERR и возвращать ненулевой код выхода в OS.
Программа должна проходить проверки go vet и golint.
*/



func main() {
	now := getCurrentTime()
	fmt.Println("Current time:", now)

	accurate, _ := accurateTime("0.beevik-ntp.pool.ntp.org")
	fmt.Println("Exact time:", accurate)
}

func accurateTime(host string) (time.Time, error) {
	response, err := ntp.Query("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		log.Fatal(err)
	}
	time := time.Now().Add(response.ClockOffset)
	return time, nil
}

func getCurrentTime() time.Time {
	time := time.Now()

	return time
}