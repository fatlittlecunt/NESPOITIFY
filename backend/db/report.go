package db

import (
	"fmt"
	strc "music/structs"
	"strconv"

	"github.com/xuri/excelize/v2"
)

func CreateReport() error {
	// Создание нового XLSX файла
	f := excelize.NewFile()
	var sum int

	// Создание нового листа
	sheetName := "Report"
	index, err := f.NewSheet(sheetName)
	if err != nil {
		return err
	}

	// Запись данных в ячейки
	f.SetCellValue(sheetName, "A1", "Артист")
	f.SetCellValue(sheetName, "B1", "Жанр")
	f.SetCellValue(sheetName, "C1", "Прослушиваний")

	artists, err := GetInfo()
	if err != nil {
		return err
	}
	var i int
	var data strc.ArtistForReport
	for i, data = range *artists {
		f.SetCellValue(sheetName, "A"+strconv.Itoa(i+2), data.ArtistName)
		f.SetCellValue(sheetName, "B"+strconv.Itoa(i+2), data.Genre)
		f.SetCellValue(sheetName, "C"+strconv.Itoa(i+2), data.Popularity)
		sum = sum + data.Popularity
	}
	f.SetCellValue(sheetName, "C"+strconv.Itoa(i+3), sum)

	// Установка активного листа
	f.SetActiveSheet(index)

	// Сохранение файла
	if err := f.SaveAs("D:/res/report.xlsx"); err != nil {
		fmt.Println(err)
		return err
	} else {
		fmt.Println("Файл успешно сохранен!")
	}
	return nil
}
