package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	/*a := app.New()//Калькулятор
	w := a.NewWindow("Недокалькулятор")
	w.Resize(fyne.NewSize(600, 400))

	label1 := widget.NewLabel("Введите перове число")
	entry1 := widget.NewEntry()

	label2 := widget.NewLabel("Введите второе число")
	entry2 := widget.NewEntry()

	answer := widget.NewLabel("Результа:")

	btn := widget.NewButton("Посчитай", func() {
		n1, err := strconv.ParseFloat(entry1.Text, 64)
		n2, er := strconv.ParseFloat(entry2.Text, 64)

		if err != nil || er != nil {
			answer.SetText("Ошибка ввода")
		} else {
			sum := n1 + n2
			sub := n1 - n2
			mul := n1 * n2
			div := n1 / n2

			answer.SetText(fmt.Sprintf("(+) %f\n(-) %f\n(*) %f\n(/) %f\n", sum, sub, mul, div))
		}

	})

	w.SetContent(container.NewVBox(
		label1,
		entry1,
		label2,
		entry2,
		btn,
		answer,
	))

	w.ShowAndRun()*/

	a := app.New()
	w := a.NewWindow("Authorization")
	w.Resize(fyne.NewSize(600, 500))

	title := widget.NewLabel("Заказ Машины")

	name_label := widget.NewLabel("Напишите свое имя")
	name_entry := widget.NewEntry()

	Parametrs := widget.NewLabel("\n Парматеры машины:")
	cheks := widget.NewCheckGroup([]string{"стоимость больше 10000$", "Темного цвета", "Спорткар", "Внедорожник"}, func(s []string) {})

	label_sex := widget.NewLabel("Укажите свой пол")
	radio_sex := widget.NewRadioGroup([]string{"Мужской", "Женский"}, func(s string) {})

	Res_label := widget.NewLabel("Статус заказа: ")

	result := widget.NewLabel("")

	btn := widget.NewButton("Оформить заказ", func() {
		Res_label.SetText("Статус заказа: оформлен")

		result.SetText(name_entry.Text + " Пол: " + radio_sex.Selected + "\n" + "Хочет заказать машину со следущими парметрами:" + "\n")
		for _, i := range cheks.Selected {
			result.SetText(result.Text + i + "\n")
		}
	})

	w.SetContent(container.NewVBox(
		title,
		name_label,
		name_entry,
		Parametrs,
		cheks,
		label_sex,
		radio_sex,
		Res_label,
		btn,
		result,
	))

	w.ShowAndRun()
}
