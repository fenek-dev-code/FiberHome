package pages

import "github.com/gofiber/fiber/v2"

type handlerPage struct {
	router fiber.Router
}

var newsTags = []string{
	"#Еда", "#Спорт", "#Машины", "#Животные", "#Музыка", "#Остальное", "#Технологии",
}

func NewHanlerPage(router fiber.Router) {
	h := &handlerPage{router: router}
	api := h.router.Group("/page")
	api.Get("/", h.homePage)

}

func (h *handlerPage) homePage(c *fiber.Ctx) error {

	return c.Render("page", fiber.Map{
		"newsTags": newsTags,
	})
}
