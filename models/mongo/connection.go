package mongo

import (
	"gopkg.in/mgo.v2"
	"os"
)

var (
	session *mgo.Session
)

func Conn() *mgo.Session {

	return session.Copy()
}

func init() {

	URL_ENV := os.Getenv("MONGO_URL")
	//mongoURL, err := url.Parse(URL_ENV);
	//if (err != nil) {
	//	panic("URL de conexao com banco não definida")
	//}
	//password, ok := mongoURL.User.Password()
	//if(!ok) {
	//	beego.Warn("database password unset")
	//}
	//dialInfo := &mgo.DialInfo{
	//	Addrs:      []string{mongoURL.Host},
	//	Timeout:    60 * time.Second,
	//	Database:   "pinnapi",
	//	Username:   mongoURL.User.Username(),
	//	Password:   password,
	//	DialServer: func(addr *mgo.ServerAddr) (net.Conn, error) {
	//		return tls.Dial("tcp", addr.String(), &tls.Config{})
	//	},
	//}

	//sess, err := mgo.DialWithInfo(dialInfo)

	sess, err := mgo.Dial(URL_ENV);
	if err != nil {
		panic(err)
	}

	session = sess
	session.SetMode(mgo.Monotonic, true)
}
