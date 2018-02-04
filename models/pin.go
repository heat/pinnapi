package models

import (
	"fmt"
	"pinnapi/models/mongo"

	"github.com/garyburd/redigo/redis"
	mgo "gopkg.in/mgo.v2"
)

type Dinheiro float64

type Palpite struct {
	Evento string `bson:"evento" json:"evento"`
	Odd    string `bson:"odd" json:"odd"`
}

type Pin struct {
	ID       string     `bson:"_id"      json:"_id,omitempty"`
	Codigo   string     `bson:"codigo"     json:"codigo,omitempty"`
	Valor    Dinheiro   `bson:"valor_aposta" json:"valor_aposta,omitempty"`
	Cliente  string     `bson:"cliente"     json:"cliente,omitempty"`
	Palpites []*Palpite `bson:"palpites"  json:"palpites"`
}

func Next() int {
	c := Pool().Get()

	defer c.Close()

	n, _ := redis.Int(c.Do("INCR", "next"))

	return n
}

func (p *Pin) FindById(id string) (code int, err error) {
	conn := mongo.Conn()
	defer conn.Close()

	fmt.Println("consultando")
	fmt.Println(id)

	c := conn.DB("pinnapi").C("pin")
	err = c.FindId(id).One(p)

	if err != nil {
		if err == mgo.ErrNotFound {
			code = ErrNotFound
		} else {
			code = ErrDatabase
		}
	} else {
		code = 0
	}
	return
}

func (p *Pin) Insert() (code int, err error) {
	conn := mongo.Conn()
	defer conn.Close()

	c := conn.DB("pinnapi").C("pin")

	err = c.Insert(p)

	if err != nil {
		if mgo.IsDup(err) {
			code = ErrDupRows
		} else {
			code = ErrDatabase
		}
	} else {
		code = 0
	}

	return
}
