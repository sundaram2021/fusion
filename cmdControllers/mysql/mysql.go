package mysql

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func Mysql() {
	if len(os.Args) < 3 || os.Args[1] != "apply" || os.Args[2] == "" {
	    fmt.Println("Usage: fusion apply -f <file.yaml>")
	    return
	}

	filePath := os.Args[2]
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
	    fmt.Println("Error reading file:", err)
	    return
	}

	var db Database
	err = yaml.Unmarshal(data, &db)
	if err != nil {
	    fmt.Println("Error parsing YAML:", err)
	    return
	}

	if err = ValidateDatabase(db); err != nil {
	    fmt.Println("Validation Error:", err)
	    return
	}

	fmt.Println("Database structure is valid!")
}