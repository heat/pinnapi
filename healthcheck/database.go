package healthcheck

import (
	"pinnapi/models/mongo"
)

type DatabaseCheck struct {
}

func (dc *DatabaseCheck) Check() error {

	conn := mongo.Conn()

	return conn.Ping()
}
