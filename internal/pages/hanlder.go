package pages

import (
	"go-fiber/home/pkg/tadapter"
	"go-fiber/home/views"

	"github.com/gofiber/fiber/v2"
)

type handlerPage struct {
	router fiber.Router
}

func NewHanlerPage(router fiber.Router) {
	h := &handlerPage{router: router}
	h.router.Get("/", h.homePage)

}

func (h *handlerPage) homePage(c *fiber.Ctx) error {
	component := views.Main()
	return tadapter.Render(c, component)
}
