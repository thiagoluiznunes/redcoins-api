package operation

import (
	"database/sql"
	"errors"
	"fmt"
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
func CreateOperation(opt Operation) error {
	insertOperationQuery := fmt.Sprintf(`
		INSERT INTO operations (uuid,	operation_type,	amount,	price,user_uuid)
		VALUES (UUID(),	'%s',	%f,	%f,	'%s');`, opt.OperationType, opt.Amount, opt.Price, opt.UserUUID)

	insert, err := DB.Query(insertOperationQuery)
	insert.Close()

	if err != nil {
		return err
	}
	return nil
}

// GetOperationsByID : select all operations by user_uuid
func GetOperationsByID(uuid string) ([]Operation, error) {
	operations := []Operation{}
	getOperationsQuery := fmt.Sprintf(`SELECT *	FROM operations	WHERE user_uuid = '%s';`, uuid)
	rows, err := DB.Query(getOperationsQuery)

	if err != nil {
		return operations, err
	}
	defer rows.Close()

	for rows.Next() {
		var opt Operation
		if err := rows.Scan(&opt.UUID, &opt.OperationType, &opt.Amount, &opt.Price, &opt.CreatedAt, &opt.UserUUID); err != nil {
			return operations, err
		}
		operations = append(operations, opt)
	}

	if err := rows.Err(); err != nil {
		return operations, err
	}
	return operations, nil
}

// GetOperationsByDate : select all operations by date
func GetOperationsByDate(date string) ([]Operation, error) {
	operations := []Operation{}
	getOperationsQuery := fmt.Sprintf(`SELECT * FROM operations
		WHERE created_at >= '%s 00:00:00' AND created_at <='%s 23:59:59';`, date, date)

	rows, err := DB.Query(getOperationsQuery)
	if err != nil {
		return operations, err
	}
	defer rows.Close()

	for rows.Next() {
		var opt Operation
		if err := rows.Scan(&opt.UUID, &opt.OperationType, &opt.Amount, &opt.Price, &opt.CreatedAt, &opt.UserUUID); err != nil {
			return operations, err
		}
		operations = append(operations, opt)
	}

	if err := rows.Err(); err != nil {
		return operations, err
	}
	return operations, nil
}

// GetOperationsByParam : retrive all operations by param
func GetOperationsByParam(param string, data string) ([]Operation, error) {
	operations := []Operation{}
	var selectUserQuery string
	var uuid string
	switch param {
	case `email`:
		selectUserQuery = fmt.Sprintf(`SELECT uuid FROM users	WHERE email = '%s'`, data)
		row := DB.QueryRow(selectUserQuery).Scan(&uuid)
		if row == sql.ErrNoRows || uuid == "" {
			return operations, errors.New("user: User not found")
		}
	case `name`:
		selectUserQuery = fmt.Sprintf(`SELECT uuid FROM users	WHERE name = '%s'`, data)
		row := DB.QueryRow(selectUserQuery).Scan(&uuid)
		if row == sql.ErrNoRows || uuid == "" {
			return operations, errors.New("user: User not found")
		}
	}

	getOperationsQuery := fmt.Sprintf(`SELECT *	FROM operations	WHERE user_uuid = '%s';`, uuid)
	rows, err := DB.Query(getOperationsQuery)
	if err != nil {
		return operations, err
	}
	defer rows.Close()

	for rows.Next() {
		var opt Operation
		if err := rows.Scan(&opt.UUID, &opt.OperationType, &opt.Amount, &opt.Price, &opt.CreatedAt, &opt.UserUUID); err != nil {
			return operations, err
		}
		operations = append(operations, opt)
	}

	if err := rows.Err(); err != nil {
		return operations, err
	}
	return operations, nil
}
