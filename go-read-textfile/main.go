package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	//"unicode/utf8"
)

func main() {
	fmt.Println("Go read text file sample app")

	exeFileName, err := os.Executable()
	if err != nil {
		panic(err)
	}

	exePath := filepath.Dir(exeFileName)

	fileName := filepath.Join(exePath, "testfile.txt") //`E:\Projects_E_Go\go-example-util\testfile.txt`

	//открываем файл, читаем его строки
	fileHandle, err := os.Open(fileName)

	if err != nil {
		fmt.Println("Помилка відкриття файлу:", err)
		return
	}

	fileScanner := bufio.NewScanner(fileHandle)

	fileScanner.Split(bufio.ScanLines)

	const radiationRune rune = '☢'

	radiationRunesCount := 0
	radiationRunesCount2 := 0
	linesInFile := 0

	for fileScanner.Scan() {
		line := fileScanner.Text()
		linesInFile++

		//пройти строку, и найти тот символ что нам нужен - способ через range
		for _, runeValue := range line {
			if runeValue == radiationRune {
				radiationRunesCount++
			}
		} //for

		//альтернативный подсчет символов:
		//runesCount := utf8.RuneCountInString(line)
		runesSlice := []rune(line) //это преобразует в массив рун

		for i := 0; i < len(runesSlice); i++ {
			itRune := runesSlice[i]
			if itRune == radiationRune {
				radiationRunesCount2++
			}
		} //for i
	}

	fileHandle.Close()

	fmt.Println("Lines in file:", linesInFile)
	fmt.Println("Total radiation runes:", radiationRunesCount)
	fmt.Println("Total radiation runes2:", radiationRunesCount2)

	fmt.Println("Finished")
}
