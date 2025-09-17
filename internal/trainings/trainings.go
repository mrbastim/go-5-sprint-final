package trainings

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/mrbastim/go-5-sprint-final/internal/personaldata"
	"github.com/mrbastim/go-5-sprint-final/internal/spentenergy"
)

type Training struct {
	// TODO: добавить поля
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
	data := strings.Split(datastring, ",")
	if len(data) != 3 {
		return fmt.Errorf("неверное количество данных")
	}
	t.Steps, err = strconv.Atoi(data[0])
	if err != nil {
		return fmt.Errorf("неверный формат количества шагов: %v", err)
	}
	t.TrainingType = data[1]

	t.Duration, err = time.ParseDuration(data[2])
	if err != nil {
		return fmt.Errorf("неверный формат длительности тренировки: %v", err)
	}
	return nil
}

func (t Training) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	distance := spentenergy.Distance(t.Steps, t.Personal.Height)
	meanSpeed := spentenergy.MeanSpeed(t.Steps, t.Personal.Height, t.Duration)

	switch t.TrainingType {
	case "Бег":
		callories, err := spentenergy.RunningSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("Тип тренировки: Бег\nДлительность: %.2f\nДистанция: %.2f\nСкорость: %.2f\nСожгли калорий: %.2f\n",
			t.Duration.Hours(), distance, meanSpeed, callories), nil
	case "Ходьба":
		callories, err := spentenergy.WalkingSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("Тип тренировки: Ходьба\nДлительность: %.2f\nДистанция: %.2f\nСкорость: %.2f\nСожгли калорий: %.2f\n",
			t.Duration.Hours(), distance, meanSpeed, callories), nil
	default:
		return "", fmt.Errorf("неизвестный тип тренировки")
	}
}
