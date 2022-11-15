package repository_test

import (
	"forum/internal/models"
	"forum/internal/repository"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestSignup(t *testing.T) {
	tests := []struct {
		desc        string
		mockClosure func(mock sqlmock.Sqlmock)
		args        models.User
		expected    error
	}{
		{
			desc: "success signed up",
			mockClosure: func(mock sqlmock.Sqlmock) {
				mock.ExpectPrepare("INSERT INTO User").
					ExpectExec().
					WithArgs("loli@gmail.com", "neo", "Aibek", "Abdikhalyk", "Qwerty123", "male", "19").WillReturnResult(sqlmock.NewResult(1, 1))
			},
			args: models.User{Email: "loli@gmail.com", Nickname: "neo", FirstName: "Aibek", LastName: "Abdikhalyk", Password: "Qwerty123", Gender: "male", Age: "19"},
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			handler, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			test.mockClosure(mock)
			db := repository.NewRepository(handler)
			if err = db.Signup(&test.args); err != test.expected {
				t.Errorf("was expected %v while putting post, got %v", test.expected, err)
			}
			mock.ExpectClose()
			if err = handler.Close(); err != nil {
				t.Error(err)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}
