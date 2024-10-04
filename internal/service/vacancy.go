package service

import (
	model "kis/internal/models"
	"kis/internal/repository"
	"strconv"
	"strings"
)

type VacancyService struct {
	vacRepo repository.Vacancy
	cusRepo repository.Customer
}

func NewVacancyService(vacRepo repository.Vacancy, cusRepo repository.Customer) *VacancyService {
	return &VacancyService{vacRepo: vacRepo, cusRepo: cusRepo}
}

func (s *VacancyService) GetVacancies(params model.QueryDTO) ([]model.VacancyDTO, error) {
	searchTerms := model.VacancySearchTermsDTO{}
	searchTerms.FromQueryDTO(&params, 10)

	models, err := s.vacRepo.GetBySearchTerms(searchTerms)
	if err != nil {
		return nil, err
	}

	vacancyDtos := make([]model.VacancyDTO, len(models))
	for i, m := range models {
		v := &vacancyDtos[i]

		v.FromVacancy(&m)
		//helper
		v.CandidatesCount = strconv.Itoa(len(m.Candidates))

		cus, err := s.cusRepo.GetById(m.CustomerID)
		if err != nil {
			return vacancyDtos, err
		}
		v.CustomerName = cus.Name
		//HELPER
		var grades []string
		for _, g := range m.Grades {
			if len(g.Name) > 0 {
				grades = append(grades, g.Name)
			}
		}
		v.Grade = strings.Join(grades, ", ")
	}

	return vacancyDtos, nil
}

func (s *VacancyService) GetVacancyById(id uint) (model.Vacancy, error) {
	return s.vacRepo.GetById(id)
}

func (s *VacancyService) CreateVacancy(vacancy model.Vacancy) (uint, error) {
	return s.vacRepo.Create(vacancy)
}
