package route

import (
	"github.com/SitanayaIvan/testing-deploy/handler"
	"github.com/SitanayaIvan/testing-deploy/user"
	"gorm.io/gorm"
)

func IntegrateFile(db *gorm.DB) RouteBody {
	var routeBody RouteBody

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	routeBody.UserHandler = handler.NewHandlerUser(userService)

	return routeBody
}
