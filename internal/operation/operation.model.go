package operation

import (
	"database/sql"
	"errors"
)

// DB : database instance
var DB *sql.DB

// Operation : user model
type Operation struct {
	uuid          string
	opertaionType string
	amount        float64
	userUUID      string
}

// InitOperationSchema : init table
func InitOperationSchema() {
	_, err := DB.Exec(`
		CREATE TABLE IF NOT EXISTS operations (
		uuid VARCHAR(36) NOT NULL UNIQUE,
		opertaion_type VARCHAR(36) NOT NULL,
		amount DOUBLE NOT NULL,
		creatAt TIMESTAMP,
		user_uuid VARCHAR(36) NOT NULL,
		CONSTRAINT pk_uuid PRIMARY KEY (uuid),
		CONSTRAINT fk_user_uuid FOREIGN KEY (user_uuid)	REFERENCES users(uuid))
		ENGINE=InnoDB DEFAULT CHARSET=utf8;`)

	if err != nil {
		panic(err)
	}
}

// CreateOperation : insert new operation in operations table
func CreateOperation(operation Operation) error {
	return errors.New("user: Not found")
}
