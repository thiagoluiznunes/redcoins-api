package operation

import (
	"database/sql"
	"fmt"
	"log"
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
		creat_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
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
	insertOperationQuery := fmt.Sprintf(`
		INSERT INTO users (uuid, operation_type, amount, user_uuid, createAt)
		VALUES (UUID(), '%s',	'%f',	'%s')`, operation.opertaionType, operation.amount, operation.userUUID)
	insert, err := DB.Query(insertOperationQuery)
	insert.Close()

	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
