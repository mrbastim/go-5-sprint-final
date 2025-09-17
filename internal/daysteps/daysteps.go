package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/mrbastim/go-5-sprint-final/internal/personaldata"
	"github.com/mrbastim/go-5-sprint-final/internal/spentenergy"
)

type DaySteps struct {
	// TODO: добавить поля
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
	stepStr, durationStr, ok := strings.Cut(datastring, ",")
	if !ok {
		return fmt.Errorf("неверное количество данных")
	}
	steps, err := strconv.Atoi(stepStr)
	if err != nil {
		return fmt.Errorf("неверный формат количества шагов: %v", err)
	} else if steps <= 0 {
		return fmt.Errorf("количество шагов должно быть больше нуля")
	}
	duration, err := time.ParseDuration(durationStr)
	if err != nil {
		return fmt.Errorf("неверный формат длительности: %v", err)
	}
	if duration <= 0 {
		return fmt.Errorf("длительность должна быть больше нуля")
	}
	ds.Duration, ds.Steps = duration, steps
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	distance := spentenergy.Distance(ds.Steps, ds.Personal.Height)
	callories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Personal.Weight, ds.Personal.Height, ds.Duration)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		ds.Steps, distance, callories), nil
}
