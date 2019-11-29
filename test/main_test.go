package test

import (
	"testing"
)

func TestRunner(t *testing.T) {
	// <setup code>
	t.Run("Test Create User", UserSignUp)
	t.Run("Test User Login", UserLogin)
	t.Run("Test Create Operation", CreateOperation)
	t.Run("Test Get Operation By User", GetOperationByUser)
	t.Run("Test Get Operation By Date", GetOperationByDate)
	t.Run("Test Delete Operation", DeleteOperation)
	t.Run("Test Delete User", DeleteUser)
	// <tear-down code>
}
