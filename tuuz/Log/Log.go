package Log

import (
	"fmt"
	"log"
	"os"
)

func Write(file_name string, logs string, discript string, other string) {
	file, err := os.OpenFile("log/"+file_name+".log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		log.Fatalln(err)
	} else {
		logger := log.New(file, "", log.LstdFlags)
		logger.Println(logs, discript, other)
		file.Close()
	}
}

func Error(file_name string, err error) {
	Write(file_name, "", "", err.Error())
}

func Err(err error) {
	Write("Error", "", "", err.Error())
}

func Errs(err error, log string) {
	Write("Error", log, "", err.Error())
}

//Database err
func Drr(err error) {
	Write("Database", "", "", err.Error())
}

func Crr(logs error) {
	Write("Common", "", "", logs.Error())
}

func Crrs(logs error, discript string) {
	Write("Common", "", discript, logs.Error())
}

func Dbrr(err error, log string) {
	fmt.Println(err.Error())
	Write("Dberror", log, "", err.Error())
}
