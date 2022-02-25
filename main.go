package main

import (
	"autogate/api"
	"fmt"
	"os"
	"strconv"
)

const (
	UserName      = "dennis"
)

type Person struct {
	Name         string
	Age          int
	AddressLine1 string `json:"address_line_1"`
	AddressLine2 string `json:"address_line_2"`
}

type IPerson interface {
	GetName() string
}

func (person *Person) GetName() string {
	return person.Name
}

func process(person IPerson) {
	canPrint, err := strconv.ParseBool(os.Getenv("ENV_CAN_PRINT"))
	if err != nil {
		fmt.Println("ignore: " + err.Error())
	}

	if canPrint {
		message := fmt.Sprintf("welcome %s \n", person.GetName())
		_ = fmt.Sprintf(message)
		fmt.Printf(message)
	} else {
		fmt.Println("cannot print at this time!")
	}
}

func main() {
	person := Person{Name: UserName}
	process(&person)

	// init api
	api.HandleRequests()
}
