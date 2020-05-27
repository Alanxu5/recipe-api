package gateway_test

import (
	"encoding/json"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/labstack/echo/v4"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
	"recipe-api/gateway"
)

var _ = Describe("Recipe", func() {
	var c echo.Context

	BeforeEach(func() {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c = e.NewContext(req, rec)
	})

	It("should get types from database", func() {
		rows := sqlmock.NewRows([]string{
			"id",
			"name",
		}).AddRow(
			123,
			"Protein")

		db, mock, err := sqlmock.New()
		if err != nil {
			fmt.Println("an error was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		recipeDbGateway := gateway.NewRecipeDbGateway(c, db)

		mock.ExpectQuery("SELECT").WillReturnRows(rows)
		recipeTypes, getErr := recipeDbGateway.GetTypes()
		if getErr != nil {
			Fail("error getting types" + getErr.Error())
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			Fail("failed db expectation")
		}

		Expect(recipeTypes[0].Id).To(Equal(123))
		Expect(recipeTypes[0].Name).To(Equal("Protein"))
	})

	It("should get methods from database", func() {
		rows := sqlmock.NewRows([]string{
			"id",
			"name",
		}).AddRow(
			123,
			"Pan")

		db, mock, err := sqlmock.New()
		if err != nil {
			fmt.Println("an error was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		recipeDbGateway := gateway.NewRecipeDbGateway(c, db)

		mock.ExpectQuery("SELECT").WillReturnRows(rows)
		recipeMethods, getErr := recipeDbGateway.GetMethods()
		if getErr != nil {
			Fail("error getting types" + getErr.Error())
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			Fail("failed db expectation")
		}

		Expect(recipeMethods[0].Id).To(Equal(123))
		Expect(recipeMethods[0].Name).To(Equal("Pan"))
	})

	It("should get recipe by id from database", func() {
		directions := json.RawMessage(`{"0": "Direction 1"}`)
		rows := sqlmock.NewRows([]string{
			"id",
			"name",
			"prep_time",
			"cook_time",
			"servings",
			"method",
			"type",
			"description",
			"directions",
		}).AddRow(
			123,
			"test recipe name",
			20,
			15,
			4,
			1,
			2,
			"test description",
			directions)

		db, mock, err := sqlmock.New()
		if err != nil {
			fmt.Println("an error was not expected when opening a stub database connection", err)
		}
		defer db.Close()

		recipeDbGateway := gateway.NewRecipeDbGateway(c, db)
		mock.ExpectQuery("SELECT").WillReturnRows(rows)
		recipe, getErr := recipeDbGateway.GetRecipe(123)
		if getErr != nil {
			Fail("error getting types" + getErr.Error())
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			Fail("failed db expectation")
		}

		Expect(recipe.Id).To(Equal(123))
		Expect(recipe.Method).To(Equal("1"))
		Expect(recipe.Type).To(Equal("2"))
	})
})
