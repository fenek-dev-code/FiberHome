package pages

import "github.com/gofiber/fiber/v2"

type handlerPage struct {
	router fiber.Router
}

func NewHanlerPage(router fiber.Router) {
	h := &handlerPage{router: router}
	api := h.router.Group("/page")
	api.Get("/", h.homePage)

}

func (h *handlerPage) homePage(c *fiber.Ctx) error {
	return c.SendString("Welcome to the Home Page!")
}
