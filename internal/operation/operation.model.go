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
	UUID          string  `json:"uuid"`
	OperationType string  `json:"operation_type"`
	Amount        float64 `json:"amount"`
	Price         float64 `json:"price"`
	CreatedAt     string  `json:"created_at"`
	UserUUID      string  `json:"user_uuid"`
}

// JSONOperationsResponse : structure to classify JSON operations response
type JSONOperationsResponse struct {
	Code       int         `json:"code"`
	Operations []Operation `json:"operations"`
}

// InitOperationSchema : init table
func InitOperationSchema() {
	_, err := DB.Exec(`
		CREATE TABLE IF NOT EXISTS operations (
		uuid VARCHAR(36) NOT NULL UNIQUE,
		operation_type ENUM('sale', 'purchase') NOT NULL,
		amount DOUBLE NOT NULL,
		price DOUBLE NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
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
		INSERT INTO operations (uuid, operation_type, amount, price, user_uuid)
		VALUES (UUID(), '%s', '%f',	%f, '%s');`,
		operation.OperationType, operation.Amount, operation.Price, operation.UserUUID)

	insert, err := DB.Query(insertOperationQuery)
	insert.Close()

	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

// GetOperations : select all operations by user_uuid
func GetOperations(uuid string) ([]Operation, error) {
	operations := []Operation{}
	getOperationQuery := fmt.Sprintf(` SELECT * FROM operations WHERE user_uuid = '%s';`, uuid)
	rows, err := DB.Query(getOperationQuery)

	if err != nil {
		log.Fatal(err)
		return operations, err
	}
	defer rows.Close()

	for rows.Next() {
		var opt Operation
		if err := rows.Scan(&opt.UUID, &opt.OperationType, &opt.Amount, &opt.Price, &opt.CreatedAt, &opt.UserUUID); err != nil {
			log.Fatal(err)
		}
		operations = append(operations, opt)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
		return operations, err
	}
	return operations, nil
}
