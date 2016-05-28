package main

import (
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/DStatIt/account"
)

const (
	server   string = "mongodb://104.197.212.163:27017,146.148.102.23:27017/socialmedia_test"
	socialDB string = "socialmedia_test"
)

func main() {

	session, err := mgo.Dial(server)
	if err != nil {
		log.Println(err)
	}
	db := session.DB(socialDB)
	accountRepo := account.MGOAccount{db}

	p := account.Proxy{
		IP:   "192.1.1.2",
		Port: 123,
	}
	if err := accountRepo.AddProxy(bson.ObjectIdHex("57463e10537bc07bfa9a56df"), p); err != nil {
		log.Println(err)
	}
}
