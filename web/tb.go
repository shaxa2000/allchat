package web

import (
	"bufio"
	"fmt"
	"os"
)

func GetAllData(path string) []string {
	var lines []string
	file, err := os.Open(path)
	if os.IsNotExist(err) {
		return nil
	}
	checkErr(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	checkErr(scanner.Err())

	len := len(lines)

	if len > 13 {
		file.Close()
		err = os.Remove("./db/dataBase.txt")
		checkErr(err)
		return lines
	}
	return lines

}

//
//
//
//
//
//

func SetData(name, message string) {
	options := os.O_WRONLY | os.O_APPEND | os.O_CREATE
	file, err := os.OpenFile("./db/dataBase.txt", options, os.FileMode(0600))
	checkErr(err)
	full := fmt.Sprintf("[%s]: %s", name, message)
	_, err = fmt.Fprintln(file, full)
	checkErr(err)
	defer file.Close()

}
