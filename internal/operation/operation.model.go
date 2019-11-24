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
	price         float64
	userUUID      string
}

// InitOperationSchema : init table
func InitOperationSchema() {
	_, err := DB.Exec(`
		CREATE TABLE IF NOT EXISTS operations (
		uuid VARCHAR(36) NOT NULL UNIQUE,
		operation_type ENUM('sale', 'purchase') NOT NULL,
		amount DOUBLE NOT NULL,
		prince DOUBLE NOT NULL,
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
	fmt.Println(operation)
	insertOperationQuery := fmt.Sprintf(`
		INSERT INTO operations (uuid, operation_type, amount, price, user_uuid)
		VALUES (UUID(), '%s', '%f',	%f, '%s');`,
		operation.opertaionType, operation.amount, operation.price, operation.userUUID)

	insert, err := DB.Query(insertOperationQuery)
	insert.Close()

	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
