package store

import (
	"app/internal/models"
	"fmt"

	sq "github.com/Masterminds/squirrel"
)
func (s *Store) UpdateProfileByID(id string, updatedProfile *models.FormData) error {
	// Используем squirrel для построения запроса
	queryBuilder := sq.Update("profiles").
		Set("email", updatedProfile.Email).
		Set("name", updatedProfile.Name).
		Set("fullname", updatedProfile.FullName).
		Set("contacts", updatedProfile.Contacts).
		Where(sq.Eq{"id": id})

	// Генерируем SQL-запрос и параметры
	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return fmt.Errorf("error building query: %v", err)
	}

	// Выполняем запрос с использованием DB.Exec
	_, err = s.db.Exec(query, args...)
	if err != nil {
		return fmt.Errorf("error updating profile: %v", err)
	}

	return nil
}
