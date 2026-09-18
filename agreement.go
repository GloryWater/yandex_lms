package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	var dateStr, firstName, lastName, patronymic string
	var payment, sum float64

	fmt.Scan(&dateStr, &firstName, &lastName, &patronymic)

	for i := 0; i < 3; i++ {
		fmt.Scan(&payment)
		sum += payment
	}

	date, _ := time.Parse("02.01.2006", dateStr)
	date = date.AddDate(0, 0, 15)

	kopecks := int64(math.Round(sum * 100))

	fmt.Printf(
		"Уважаемый, %s %s %s, доводим до вашего сведения, что бухгалтерия сформировала документы по факту выполненной вами работы.\n"+
			"Дата подписания договора: %02d.%02d.%04d. Просим вас подойти в офис в любое удобное для вас время в этот день.\n"+
			"Общая сумма выплат составит %d руб. %d коп.\n\n"+
			"С уважением,\nГл. бух. Иванов А.Е.",
		lastName,
		firstName,
		patronymic,
		date.Day(),
		date.Month(),
		date.Year(),
		kopecks/100,
		kopecks%100,
	)
}
