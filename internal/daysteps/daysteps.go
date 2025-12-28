package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")
	if len(parts) != 2 {
		return fmt.Errorf("неверный формат данных: ожидается 2 части, получено %d", len(parts))
	}

	if strings.Contains(parts[0], " ") || strings.Contains(parts[1], " ") {
		return fmt.Errorf("неверный формат данных: пробелы не допускаются")
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("ошибка парсинга шагов: %v", err)
	}
	if steps <= 0 {
		return fmt.Errorf("количество шагов должно быть положительным")
	}
	ds.Steps = steps

	duration, err := time.ParseDuration(parts[1])
	if err != nil {
		return fmt.Errorf("ошибка парсинга длительности: %v", err)
	}
	if duration <= 0 {
		return fmt.Errorf("длительность должна быть положительной")
	}
	ds.Duration = duration

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {

	distance := spentenergy.Distance(ds.Steps, ds.Height)

	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", fmt.Errorf("ошибка расчета калорий: %v", err)
	}

	result := fmt.Sprintf("Количество шагов: %d.\n", ds.Steps)
	result += fmt.Sprintf("Дистанция составила %.2f км.\n", distance)
	result += fmt.Sprintf("Вы сожгли %.2f ккал.\n", calories)

	return result, nil
}
