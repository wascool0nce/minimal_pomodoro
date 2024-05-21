package cmd

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var pomodoroCmd = &cobra.Command{
	Use:   "start",
	Short: "Запустить таймер pomodoro, на вход требуется 3 аргумента: рабочее время, время перерыва и количество помодоров. ",
	Run:   pomodoroTimer,
}

func pomodoroTimer(_ *cobra.Command, args []string) {
	if len(args) < 3 {
		fmt.Println("Ошибка: недостаточно аргументов. Требуется 3 аргумента: рабочее время, время перерыва и количество помодоров.")
	}

	workDuration, err := parseDuration(args[0])
	if err != nil {
		fmt.Printf("Ошибка при преобразовании рабочего времени: %v\n", err)
		return
	}

	breakDuration, err := parseDuration(args[1])
	if err != nil {
		fmt.Printf("Ошибка при преобразовании времени перерыва: %v\n", err)
		return
	}

	numPomodors, err := strconv.Atoi(strings.TrimSpace(args[2]))
	if err != nil {
		fmt.Printf("Ошибка при преобразовании количества помодоров: %v\n", err)
		return
	}

	for i := 0; i < numPomodors; i++ {
		fmt.Println("Запуск pomodoro, время работать!")
		startTimer(workDuration)
		fmt.Println("Время перерыва, пора отдохнуть")
		startTimer(breakDuration)
	}
	fmt.Println("Все готово! Хорошая работа!")

}

func parseDuration(s string) (time.Duration, error) {
	number, err := strconv.Atoi(strings.TrimSpace(s))
	if err != nil {
		return 0, err
	}
	return time.Duration(number) * time.Minute, nil
}

func startTimer(duration time.Duration) {
	timer := time.NewTimer(duration)

	<-timer.C
	fmt.Println("Time's up!")
}

func init() {
	rootCmd.AddCommand(pomodoroCmd)
}
