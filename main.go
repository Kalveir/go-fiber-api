package main
import(
       "github.com/gofiber/fiber/v3"
       "go-fiber-api/models"
       "go-fiber-api/controllers"
)

func main(){
        models.ConnectDatabase()

        app := fiber.New()

        api := app.Group("/api")
        book := app.Group("/books")

        //Book
        book.Get("/", bookcontroller.Index)
        book.Get("/show/:id",bookcontroller.Show)
        book.Post("/add/",bookcontroller.Store)
        book.Put("/update/:id",bookcontroller.Update)
        book.Delete("delete/:id", bookcontroller.Delete)

        app.Listen("8000")
}

