package profile

import (
	"app/internal/models"
	"net/http"
	"app/internal/store"

	"github.com/labstack/echo/v4"
)
//settings.html мой но я пока хз как с ним что делать  
func UserSettingHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "settings.html", nil)
}


// Объявляем структуру хендлера с зависимостью
type Handler struct {
	Store store.ProfileStore // Поле Store должно ссылаться на интерфейс ProfileStore
}

// NewHandler создаёт новый хендлер с внедрённым интерфейсом
func NewUserSettingsHandler(store store.ProfileStore) *Handler {
	return &Handler{Store: store}
}
// UpdateUserSettings обновляет профиль пользователя
func (h *Handler) UpdateUserSettings(c echo.Context) error {
	UserID := c.Param("id")
	var updatedProfile models.FormData
	if err := c.Bind(&updatedProfile); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request payload",
		})
	}
	// Вызов метода интерфейса ProfileStore для обновления профиля
	if err := h.Store.UpdateProfileByID(UserID, &updatedProfile); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to update profile",
		})
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Profile updated successfully",
	})
}