package bookmark

import "gorm.io/gorm"

type BookmarkRepository struct {
	db *gorm.DB
}

func NewBookmarkRepository(db *gorm.DB) *BookmarkRepository {
	return &BookmarkRepository{db: db}
}

func (r *BookmarkRepository) Create(employeeId, vacancyId uint) error {
	err := r.db.Exec("INSERT INTO employeesvacancies VALUES(?, ?)", employeeId, vacancyId).Error
	return err
}

func (r *BookmarkRepository) Delete(employeeId, vacancyId uint) error {
	err := r.db.Exec("DELETE FROM employeesvacancies WHERE employee_id = ? AND vacancy_id = ?", employeeId, vacancyId).Error
	return err
}
