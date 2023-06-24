package main

import (
	"fmt"
	"io/fs"
	"os"
	"path"
	"sync"
)

func main() {
	fmt.Println("Go search files and directories - goroutine sample app")

	dir := `E:\Images`

	fmt.Println("Start folder:", dir)

	var wg sync.WaitGroup
	wg.Add(2)

	go _ReadDirectories(dir, &wg)
	go _SimpleCounter(10, &wg)

	wg.Wait() //ждем пока завершатся оба потока

	fmt.Println("Finished")
}

func _ReadDirectories(dirPath string, wg *sync.WaitGroup) {
	defer wg.Done() //defer означает что вызов выполнится перед выходом из ф-ции
	_ReadDir(dirPath)
}

func _ReadDir(dirPath string) {

	dirItems, err := os.ReadDir(dirPath)
	if err != nil {
		panic(err)
	}

	for i, fsItem := range dirItems {

		itemType := fsItem.Type()

		//по маске ищем только файлы
		if itemType&fs.ModeType != 0 {

			if fsItem.IsDir() {

				subDirPath := path.Join(dirPath, fsItem.Name())
				_ReadDir(subDirPath)
			}

		} else {

			//time.Sleep(2 * time.Second)

			fileName := fsItem.Name()
			matchedPng, _ := path.Match("*.png", fileName)
			matchedZip, _ := path.Match("*.zip", fileName)

			if matchedPng || matchedZip {
				fullName := path.Join(dirPath, fsItem.Name())
				fmt.Println(i, ": ", fullName)
			}
		}

	} //for

}
