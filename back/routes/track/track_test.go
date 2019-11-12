package track_test

import (
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"

	"github.com/markelog/pento/back/database"
	"github.com/markelog/pento/back/logger"
	"github.com/markelog/pento/back/routes/track"
	"github.com/markelog/pento/back/test/env"
	"github.com/markelog/pento/back/test/request"
	"github.com/markelog/pento/back/test/routes"
	"github.com/markelog/pento/back/test/schema"
)

var (
	app *iris.Application
	db  *gorm.DB
)

func teardown() {
	db.Raw("TRUNCATE users CASCADE;").Row()
	db.Raw("TRUNCATE tracks CASCADE;").Row()
}

func TestMain(m *testing.M) {
	env.Up()

	app = routes.Up()
	db = database.Up()
	log := logger.Up()
	log.Out = ioutil.Discard

	track.Up(app, db, log)

	app.Build()

	os.Exit(m.Run())
}

func TestBothStartStopPresent(t *testing.T) {
	defer teardown()
	req := request.Up(app, t)

	data := map[string]interface{}{
		"email":  "markelog@gmail.com",
		"name":   "test",
		"active": true,
		"start":  "2014-01-08T08:54:44+01:00",
		"stop":   "2014-01-08T08:54:44+01:00",
	}

	test := req.POST("/tracks").
		WithHeader("Content-Type", "application/json").
		WithJSON(data).
		Expect().
		Status(http.StatusBadRequest)

	json := test.JSON()

	json.Schema(schema.Response)

	json.Object().
		Value("payload").Object().
		Value("errors").Array().
		Elements("(root): Must validate one and only one schema (oneOf)")
}

func TestActivePresent(t *testing.T) {
	defer teardown()
	req := request.Up(app, t)

	data := map[string]interface{}{
		"email": "markelog@gmail.com",
		"name":  "test",
		"start": "2014-01-08T08:54:44+01:00",
	}

	test := req.POST("/tracks").
		WithHeader("Content-Type", "application/json").
		WithJSON(data).
		Expect().
		Status(http.StatusBadRequest)

	json := test.JSON()

	json.Schema(schema.Response)

	json.Object().
		Value("payload").Object().
		Value("errors").Array().
		Elements("active: active is required")
}

func TestSuccess(t *testing.T) {
	defer teardown()
	req := request.Up(app, t)

	data := map[string]interface{}{
		"email":  "markelog@gmail.com",
		"start":  "2014-01-08T08:54:44+01:00",
		"active": true,
	}

	project := req.POST("/tracks").
		WithHeader("Content-Type", "application/json").
		WithJSON(data).
		Expect().
		Status(http.StatusOK)

	project.JSON().Schema(schema.Response)
}

func TestList(t *testing.T) {
	defer teardown()
	teardown()
	req := request.Up(app, t)

	data := map[string]interface{}{
		"email":  "markelog@gmail.com",
		"name":   "sup",
		"start":  "2014-01-08T08:54:44+01:00",
		"active": true,
	}

	req.POST("/tracks").
		WithHeader("Content-Type", "application/json").
		WithJSON(data).
		Expect().
		Status(http.StatusOK)

	result := req.GET("/tracks/markelog@gmail.com").
		Expect().
		Status(http.StatusOK).
		JSON()

	result.Schema(schema.Response)

	element := result.Object().Value("payload").Array().
		Element(0).Object()

	element.Value("name").Equal("sup")
}
