package service

import (
	"fmt"
	"kis/internal/repository"
)

type BookmarkService struct {
	bookmarkRepo repository.Bookmark
}

func NewBookmarkService(bookmarkRepo repository.Bookmark) *BookmarkService {
	return &BookmarkService{bookmarkRepo: bookmarkRepo}
}

func (s *BookmarkService) CreateBookmark(employeeId, vacancyId uint) error {
	if err := s.bookmarkRepo.Create(employeeId, vacancyId); err != nil {
		return fmt.Errorf("cannot create bookmark: %s", err.Error())
	}
	return nil
}

func (s *BookmarkService) DeleteBookmark(employeeId, vacancyId uint) error {
	if err := s.bookmarkRepo.Delete(employeeId, vacancyId); err != nil {
		return fmt.Errorf("cannote delete bookmark: %s", err.Error())
	}
	return nil
}
