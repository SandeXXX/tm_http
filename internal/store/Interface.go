package store
	import(
  "app/internal/models"
)

type UserStore interface {
  UserExists(email string) (bool, error)
  CreateUser(user *models.SignUpParams) error
}
type ProfileStore interface {
	UpdateProfileByID(id string, updatedProfile *models.FormData) error
}