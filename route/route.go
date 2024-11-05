package route

import(
	"github.com/gofiber/fiber/v3"
	"github.com/Kalveir/go-fiber-api/handler"
)

func Setup(app *fiber.App){
	app.Get("/book", handler.GetBooks)
	app.Get("/book/:id", handler.FindBooks)
	app.Post("/book/store", handler.StoreBooks)
	app.Put("/book/:id", handler.UpdateBooks)
	app.Delete("/book/:id", handler.DeleteBooks)
}
