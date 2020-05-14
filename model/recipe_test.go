package model_test

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"recipe-api/model"
)

var _ = Describe("Recipe", func() {
	It("Should get types from database", func() {
		rows := sqlmock.NewRows([]string{
			"id", "name",
		}).AddRow(
			123,
			"Protein")

		db, mock, err := sqlmock.New()
		if err != nil {
			fmt.Println("an error was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		database := model.DB{
			DB: db,
		}

		mock.ExpectQuery("SELECT").WillReturnRows(rows)
		recipeTypes, getErr := database.GetTypes()
		if getErr != nil {
			Fail("error getting types" + getErr.Error())
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			Fail("failed db expectation")
		}

		Expect(recipeTypes[0].Id).To(Equal(123))
		Expect(recipeTypes[0].Name).To(Equal("Protein"))
	})

	It("Should get methods from database", func() {
		rows := sqlmock.NewRows([]string{
			"id", "name",
		}).AddRow(
			123,
			"Pan")

		db, mock, err := sqlmock.New()
		if err != nil {
			fmt.Println("an error was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		database := model.DB{
			DB: db,
		}

		mock.ExpectQuery("SELECT").WillReturnRows(rows)
		recipeMethods, getErr := database.GetMethods()
		if getErr != nil {
			Fail("error getting types" + getErr.Error())
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			Fail("failed db expectation")
		}

		Expect(recipeMethods[0].Id).To(Equal(123))
		Expect(recipeMethods[0].Name).To(Equal("Pan"))
	})
})
