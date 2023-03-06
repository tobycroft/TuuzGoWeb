package Log

import (
	"fmt"
	"github.com/tobycroft/gorose-pro"
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
	if err != nil {
		Write(file_name, "", "", err.Error())
	}
}

func Err(err error) {
	if err != nil {
		Write("Error", "", "", err.Error())
	}
}

func Errs(err error, log string) {
	fmt.Println(log, err)
	if err != nil {
		Write("Error", log, "", err.Error())
	}
}

// Database err
func Drr(err error) {
	if err != nil {
		Write("Database", "", "", err.Error())
	}
}

func Crr(logs error) {
	if logs != nil {
		Write("Common", "", "", logs.Error())
	}
}

func Crrs(logs error, discript string) {
	fmt.Println(logs, discript)
	if logs != nil {
		Write("Common", "", discript, logs.Error())
	}
}

func Dbrr(err error, log string) {
	fmt.Println(err, log)
	if err != nil {
		Write("Dberror", log, "", err.Error())
	}
}

func DBrrsql(err error, db gorose.IOrm, log string) {
	fmt.Println(err, log)
	if err != nil {
		Write("Dberror", log, db.LastSql(), err.Error())
	}
}
