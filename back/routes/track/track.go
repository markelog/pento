package track

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"
	"github.com/sirupsen/logrus"
	"github.com/xeipuuv/gojsonschema"

	controller "github.com/markelog/pento/back/controllers/track"
)

type postTrack struct {
	Email  string     `json:"email,omitempty"`
	Active bool       `json:"active,omitempty"`
	Name   string     `json:"name,omitempty"`
	Start  *time.Time `json:"start,omitempty"`
	Stop   *time.Time `json:"stop,omitempty"`
}

var schema = gojsonschema.NewStringLoader(`{
	"type": "object",
	"properties": {
		"email": {"type": "string", "format": "email"},
		"active": {"type": "boolean"},
		"name": {"type": "string"},
		"start": {"type": ["string", "null"], "format": "date-time"},
		"stop": {"type": ["string", "null"], "format": "date-time"}
	},
	"required": ["email", "active"],
	"oneOf": [
		{
			"required": ["start"]
		},
		{
			"required": ["stop"]
		}
	]
}`)

func validate(params *postTrack) (*gojsonschema.Result, *iris.Map) {
	var (
		paramsLoader = gojsonschema.NewGoLoader(params)
		check, _     = gojsonschema.Validate(schema, paramsLoader)

		errors  []string
		payload *iris.Map
	)

	if check.Valid() == false {
		for _, desc := range check.Errors() {
			errors = append(errors, desc.String())
		}

		payload = &iris.Map{"errors": errors}

		return check, payload
	}

	return check, nil
}

// Up track route
func Up(app *iris.Application, db *gorm.DB, log *logrus.Logger) {
	app.Post("/tracks", func(ctx iris.Context) {
		var params postTrack
		ctx.ReadJSON(&params)

		validation, errors := validate(&params)

		if validation.Valid() == false {
			log.WithFields(logrus.Fields{
				"email":  params.Email,
				"active": params.Active,
			}).Error("Params are not valid")

			ctx.StatusCode(iris.StatusBadRequest)
			ctx.JSON(iris.Map{
				"active":  "failed",
				"message": "Params are not valid",
				"payload": errors,
			})

			return
		}

		err := controller.New(db).Create(&controller.CreateArgs{
			Email:  params.Email,
			Active: params.Active,
			Name:   params.Name,
			Start:  params.Start,
			Stop:   params.Stop,
		})

		if err != nil {
			log.WithFields(logrus.Fields{
				"email":  params.Email,
				"active": params.Active,
			}).Error("Can't track ya")

			ctx.StatusCode(iris.StatusBadRequest)
			ctx.JSON(iris.Map{
				"active":  "failed",
				"message": "Can't track ya",
				"payload": iris.Map{},
			})

			return
		}

		log.WithFields(logrus.Fields{
			"email":  params.Email,
			"active": params.Active,
		}).Info("Tracking")

		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(iris.Map{
			"active":  "created",
			"message": "Yey!",
			"payload": iris.Map{},
		})
	})

	app.Get("/tracks/{email}", func(ctx iris.Context) {
		ctrl := controller.New(db)
		tracks, err := ctrl.List(ctx.Params().Get("email"))

		if err != nil {
			errorString := err.Error()
			log.Error(errorString)

			ctx.StatusCode(iris.StatusBadRequest)
			ctx.JSON(iris.Map{
				"active":  "failed",
				"message": "Something went wrong",
				"payload": iris.Map{},
			})

			return
		}

		if len(tracks) == 0 {
			log.Warn("Can't find any tracks")

			ctx.StatusCode(iris.StatusNotFound)
			ctx.JSON(iris.Map{
				"active":  "failed",
				"message": "Not found",
				"payload": []string{},
			})

			return
		}

		log.Info("Tracks returned")

		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(iris.Map{
			"active":  "success",
			"message": "There you go",
			"payload": tracks,
		})

		return
	})

}
