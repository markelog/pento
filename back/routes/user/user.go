package user

import (
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris/v12"
	"github.com/sirupsen/logrus"

	controller "github.com/markelog/pento/back/controllers/user"
)

// Up user route
func Up(app *iris.Application, db *gorm.DB, log *logrus.Logger) {
	app.Get("/user/status/{email}", func(ctx iris.Context) {
		email := ctx.Params().Get("email")
		ctrl := controller.New(db)
		user, err := ctrl.Status(email)

		if err != nil {
			errorString := err.Error()
			log.Error(errorString)

			ctx.StatusCode(iris.StatusBadRequest)
			ctx.JSON(iris.Map{
				"status":  "failed",
				"message": "Something went wrong",
				"payload": iris.Map{},
			})

			return
		}

		if user == nil {
			log.Errorf("Can't find user %s", email)

			ctx.StatusCode(iris.StatusNotFound)
			ctx.JSON(iris.Map{
				"status":  "failed",
				"message": "Not found",
				"payload": iris.Map{},
			})

			return
		}

		log.Infof("User data for %s returned", email)

		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(iris.Map{
			"status":  "success",
			"message": "There you go",
			"payload": user,
		})

		return
	})

}
