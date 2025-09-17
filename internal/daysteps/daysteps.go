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
	data := strings.Split(datastring, ",")
	if len(data) != 2 {
		return fmt.Errorf("неверное количество данных")
	}
	ds.Steps, err = strconv.Atoi(data[0])
	if err != nil {
		return fmt.Errorf("неверный формат количества шагов: %v", err)
	}
	ds.Duration, err = time.ParseDuration(data[1])
	if err != nil {
		return fmt.Errorf("неверный формат длительности: %v", err)
	}
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
