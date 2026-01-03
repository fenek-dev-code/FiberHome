package pages

import (
	"go-fiber/home/pkg/tadapter"
	"go-fiber/home/views"
	"go-fiber/home/views/component"
	"time"

	"github.com/gofiber/fiber/v2"
)

type handlerPage struct {
	router fiber.Router
}

func NewHanlerPage(router fiber.Router) {
	h := &handlerPage{router: router}
	h.router.Get("/", h.homePage)
	h.router.Get("/register", h.registerPage)
	h.router.Post("/register", h.registerPost)

}

func (h *handlerPage) homePage(c *fiber.Ctx) error {
	component := views.Main()
	return tadapter.Render(c, component)
}

func (h *handlerPage) registerPage(c *fiber.Ctx) error {
	component := views.Register()
	return tadapter.Render(c, component)
}

func (h *handlerPage) registerPost(c *fiber.Ctx) error {
	time.Sleep(2 * time.Second)
	form := NewRegisterFormModel(c)
	verrs := form.Validate()
	if verrs.HasAny() {
		resp := component.Notification(verrs.Error(), component.NotificationFail)
		return tadapter.Render(c, resp)
	}
	resp := component.Notification("Регистрация прошла успешно!", component.NotificationSuccess)
	return tadapter.Render(c, resp)
}
