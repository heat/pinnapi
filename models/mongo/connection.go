package mongo

import (
	"crypto/tls"
	"net"
	"time"

	"gopkg.in/mgo.v2"
)

var (
	session *mgo.Session
)

func Conn() *mgo.Session {

	return session.Copy()
}

func init() {

	dialInfo := &mgo.DialInfo{
		Addrs:      []string{"sysbetv2.documents.azure.com:10255"},
		Timeout:    60 * time.Second,
		Database:   "pinnapi",
		Username:   "sysbetv2",
		Password:   "foygzUQ59JzVWV4xVSNKsbQibHIlruiUsWlxRhRWgSHa591d8ewCrgyx8YjyJsyjDbM9zTp0loM7rOwQXJOoDg==",
		DialServer: func(addr *mgo.ServerAddr) (net.Conn, error) { return tls.Dial("tcp", addr.String(), &tls.Config{}) },
	}

	sess, err := mgo.DialWithInfo(dialInfo)

	if err != nil {
		panic(err)
	}

	session = sess
	session.SetMode(mgo.Monotonic, true)
}
