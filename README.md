## Утилита для проверки файла-выгрузки из клиент-банка

Утилита проверяет файл, который указывается в качестве параметра запуска для нее.

В качестве файла используется структурированная выгрузка из клиент-банка ДБО, подготовленная для 1С.

Утилита проверяет совпадают-ли заявленные в файле суммы поступления и списания (***ВсегоПоступило=*** и ***ВсегоСписано=***) с фактическими поступлениями и списаниями на расчетный счет организации, фигурирующими в платежных поручениях, содержащихся в том же файле (блоки ***СекцияДокумент=Платежное поручение***).


### Обрабатываемый файл
Файл ***data.txt*** содержит пример заполнения, в части значимых для проекта полей.

Формат файла ***data.txt*** должен быть **UTF-8**.


### Запуск / Сборка
Пример запуска утилиты:
> go run .\main.go "D:\GO\path\data.txt"

Сборка и использование исполняемого файла:
> go build .\main.go
>
> main.exe "D:\GO\path\data.txt"




