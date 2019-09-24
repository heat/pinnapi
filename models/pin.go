package models

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"pinnapi/models/mongo"
	mgo "gopkg.in/mgo.v2"
	"strconv"
	"sync/atomic"
)

type Dinheiro float64

var (
	id int64 = 0
)
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

func init() {

	//tenta verificar se ja existe
	conn := mongo.Conn()
	defer conn.Close()

	c := conn.DB("pinnapi").C("pin")
	var lastPIN Pin
	var lastID = 10000;

	iter := c.Find(bson.M{}).Iter();

	for iter.Next(&lastPIN) {
		currID, _ := strconv.Atoi(lastPIN.ID);
		if currID > lastID {
			lastID = currID;
		}
	}


	if lastID > 0 {
		atomic.StoreInt64(&id, int64(lastID));
	}
}

func Next() int64 {
	atomic.AddInt64(&id, 1);



	fmt.Println("consultando")
	fmt.Println(id)

	readID := atomic.LoadInt64(&id);
	return readID;
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
