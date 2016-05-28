package main

import (
	"bufio"
	"log"
	"os"
	"strings"

	"gopkg.in/mgo.v2"

	"github.com/DStatIt/account"
)

const (
	server   = "mongodb://104.197.212.163:27017,146.148.102.23:27017/socialmedia_test"
	socialDB = "socialmedia_test"
)

func main() {
	session, err := mgo.Dial(server)
	if err != nil {
		log.Fatal(err)
	}

	accountRepo := account.MGOAccount{
		Database: session.DB(socialDB),
	}

	file, err := os.Open("accounts.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result := strings.Split(scanner.Text(), ":")
		a := account.Account{
			Type: account.Instagram,

			Username:      result[0],
			Password:      result[1],
			Email:         result[2],
			EmailPassword: result[3],
			PhoneNumber:   result[4],

			Device: account.GenerateDevice(),
		}

		log.Printf("Account: %+v", a)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
