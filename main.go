package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Документ - платежное поручение
type document struct {
	sum          string // Сумма
	date         string // Дата
	sender       string // Плательщик1
	recipient    string // Получатель1
	senderAcc    string // Расчетный счет отправителя
	recipientAcc string // Расчетный счет получателя
}

var ourCompanyAcc string // Информация из файла - наш расчетный счет
var totalIncome string   // Информация из файла - ВсегоПоступило
var totalOutcome string  // Информация из файла - ВсегоСписано

var calcIncome float64 = 0  // вычисляемая сумма всех приходов на наш счет
var calcOutcome float64 = 0 // вычисляемая сумма всех расходов с нашего счета

var documents []document

func main() {

	// файл для обработки указывается в качестве параметра запуска утилиты
	if len(os.Args) < 2 {
		fmt.Println("Не указан файл!")
		return
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	currentDoc := document{}

	// построчный перебор файла
	for fileScanner.Scan() {
		currentString := fileScanner.Text()

		// начало блока нового документа в файле - сбрасываем текущий документ для заполнения новыми данными
		if currentString == "СекцияДокумент=Платежное поручение" {
			currentDoc = document{}
		}

		// конец блока текущего документа в файле - добавляем текущий документ к слайсу документов
		if currentString == "КонецДокумента" {
			documents = append(documents, currentDoc)
			continue
		}

		s := strings.Split(currentString, "=")

		switch s[0] {
		case "Плательщик1":
			currentDoc.sender = s[1]
		case "Получатель1":
			currentDoc.recipient = s[1]
		case "Дата":
			currentDoc.date = s[1]
		case "Сумма":
			currentDoc.sum = s[1]
		case "ПлательщикРасчСчет":
			currentDoc.senderAcc = s[1]
		case "ПолучательРасчСчет":
			currentDoc.recipientAcc = s[1]
		case "ВсегоПоступило":
			totalIncome = s[1]
		case "ВсегоСписано":
			totalOutcome = s[1]
		case "РасчСчет":
			ourCompanyAcc = s[1]
		}
	}

	file.Close()

	for _, doc := range documents {
		if sum, err := strconv.ParseFloat(doc.sum, 64); err == nil {
			// если счет получателя в платежке - это наш счет, значит это приход, иначе - расход
			if doc.recipientAcc == ourCompanyAcc {
				calcIncome += sum
			} else {
				calcOutcome += sum
			}
		}
	}

	var verdictIn, verdictOut string

	if totalIncome == fmt.Sprintf("%.2f", calcIncome) {
		verdictIn = "СОВПАДАЮТ"
	} else {
		verdictIn = "!!! РАЗЛИЧАЮТСЯ !!!"
	}

	if totalOutcome == fmt.Sprintf("%.2f", calcOutcome) {
		verdictOut = "СОВПАДАЮТ"
	} else {
		verdictOut = "!!! РАЗЛИЧАЮТСЯ !!!"
	}

	fmt.Println("")
	fmt.Println("Платежных поручений в файле:", len(documents))
	fmt.Println("")
	fmt.Println("ВСЕГО ПОСТУПИЛО: заявлено в файле ->", totalIncome, " / ", fmt.Sprintf("%.2f", calcIncome), "<- вычислено --", verdictIn)
	fmt.Println("ВСЕГО СПИСАНО  : заявлено в файле ->", totalOutcome, " / ", fmt.Sprintf("%.2f", calcOutcome), "<- вычислено --", verdictOut)
	fmt.Println("")

	// fmt.Println(documents)
}
