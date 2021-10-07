# Cobra
[Источник](https://github.com/spf13/cobra/tree/master/cobra)  
Инструкция по установке:
```
go get github.com/spf13/cobra/cobra
cobra init --pkg-name github.com/p12s/kazanexpress-chucknorris-jokes .

cobra add serve
cobra add config
cobra add create -p 'configCmd'
```
Настройки:
```
touch .kazanexpress-chucknorris-jokes.yaml

nano .kazanexpress-chucknorris-jokes.yaml

author: p12s <working-tam@yandex.com>
license: MIT
```

Пример запуска команд:
```
go run main.go serve
go run main.go help serve
```

## Пример добавления команды "random" без параметров
1. Генерация файлов:
```
cobra add random
```
2. Редактирование в авто-сгенерированных файлах текстов
3. Добавление логики в команду (содержится в /internal/)
4. Вызвать хелп:
```
go run main.go help random
```
Выполнить команду:
```
go run main.go random
```

## Пример добавления команды "dump -n 2" с параметрами
1. Генерация файлов:
```
cobra add dump
```
из коробки выполняется команда:
```
go run main.go dump
```

2. Добавление возможности передать флаг:
```
go run main.go dump --number 2
или
go run main.go dump -n 2
```
Доработка функции:
```
var (
	number int

	dumpCmd = &cobra.Command{
		Use:   "dump",
		Short: "Upload and save into files n random unique jokes",
		Long: `Upload n random unique jokes (n is set by the user, default value is 5) 
for each of the existing categories and save them to text files - 
one for each of the categories.`,
		Run: func(cmd *cobra.Command, args []string) {
			number, _ := cmd.Flags().GetInt("number")
			internal.Dump(number)
		},
	}
)

func init() {
	rootCmd.AddCommand(dumpCmd)
	dumpCmd.Flags().IntVarP(&number, "number", "n", 5, "number of jokes in each category")
}
```
