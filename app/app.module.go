package app

import (
	order "yoojinyoung/gin/app/modules/order"

	"github.com/gin-gonic/gin"
)

type AppModule struct {
	engine      *gin.Engine
	orderModule *order.OrderModule
}

func NewAppModule() *AppModule {
	return new(AppModule)
}

func (app *AppModule) Init() {
	app.engine = gin.Default()
	app.orderModule = order.NewOrderModule(app.engine)

	(*app.engine).Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
