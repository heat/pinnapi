package models

import (
  "fmt"
  
  "github.com/garyburd/redigo/redis"
)

type Dinheiro float64 

type Palpite struct {
  Evento string
  Odd string
}

type Pin struct {
  
  Codigo string
  Valor Dinheiro
  Cliente string
  Palpites []*Palpite
}

func Next() int {
  c := Pool().Get()
  
  defer c.Close()

  n, _ := redis.Int(c.Do("INCR", "next"))
  fmt.Printf("%#v\n", n)

  return n;
}