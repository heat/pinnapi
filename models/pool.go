package models

import (
  "sync"
  "time"
  
  "github.com/garyburd/redigo/redis"
)
//// sysbilhete.redis.cache.windows.net:6380,password=erZKWFS8Qp35/CBBeDuyT+VQOBzzSKW/Ujx+2DOTAUY=,ssl=True,abortConnect=False
var (
  once sync.Once
  pool *redis.Pool
)

func Pool() (*redis.Pool) {
  once.Do(func() {
    pool = &redis.Pool {
      MaxIdle: 3,
      IdleTimeout: 240 * time.Second,
      Dial: func() (redis.Conn, error) { return redis.Dial("tcp", "sysbilhete.redis.cache.windows.net:6379", redis.DialPassword("erZKWFS8Qp35/CBBeDuyT+VQOBzzSKW/Ujx+2DOTAUY="))},
    }
  })
  
  return pool
}