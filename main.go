package main

import (
	"fmt"
	"time"
)

const SexMale = "male"
const SexFemale = "female"

var SexMap = map[string]string{
	SexMale:   SexFemale,
	SexFemale: SexMale,
}

type HumanError struct {
	When time.Time
	What string
}

type Human struct {
	Name     string
	Birthday time.Time
	Sex      string
	IsTrans  bool
}

func (human *Human) ChangeSex() error {
	if human.Sex == "" {
		message := fmt.Sprintf("Не можем поменять %v пол: пол не указан", human.Name)
		return &HumanError{time.Now(), message}
	}

	newSex, ok := SexMap[human.Sex]
	if !ok {
		message := fmt.Sprintf("%v не может поменять пол: неизвестный текущий пол", human.Name)
		return &HumanError{time.Now(), message}
	}

	if human.IsTrans {
		message := fmt.Sprintf("%v уже менял(а) пол. Повторная операция его (её) убьет", human.Name)
		return &HumanError{time.Now(), message}
	}

	err := human.IsMinor()

	if err != nil {
		return err
	}

	human.Sex = newSex

	return nil
}

func (human *Human) IsMinor() error {
	eighteenYearsAgo := time.Now().AddDate(-18, 0, 0)

	if eighteenYearsAgo.Before(human.Birthday) {
		message := fmt.Sprintf("%v слишком молод(а) для транс перехода", human.Name)
		return &HumanError{time.Now(), message}
	}

	return nil
}

func (human *HumanError) Error() string {
	return fmt.Sprintf("%v %v", human.When.Format("01.02.2006 15:04:05"), human.What)
}

func (human *Human) ShowTransitionComplete() {
	oldSex, ok := SexMap[human.Sex]

	if !ok {
		fmt.Println("Странная ошибка. Не удалось получить старый пол пользователя ")
		return
	}

	fmt.Printf(
		"%v Пол %v успешно изменен: c %v на %v \n",
		time.Now().Format("01.02.2006 15:04:05"),
		human.Name,
		oldSex,
		human.Sex,
	)
}

func main() {
	humans := [3]Human{
		{
			"Иванов Иван Иваныч",
			time.Date(2006, time.May, 1, 12, 0, 0, 0, time.UTC),
			SexMale,
			false,
		},
		{
			"Трамп Иванка Дональдавна",
			time.Date(1989, time.July, 1, 12, 0, 0, 0, time.UTC),
			SexFemale,
			true,
		},
		{
			"Барак Обама Хз",
			time.Date(2012, time.July, 1, 12, 0, 0, 0, time.UTC),
			SexMale,
			false,
		},
	}

	for _, human := range humans {
		err := human.ChangeSex()
		if err != nil {
			fmt.Println(err)
		} else {
			human.ShowTransitionComplete()
		}
	}
}
