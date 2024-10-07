package service

import (
	"fmt"
	"kis/internal/consts"
	model "kis/internal/models"
	"kis/internal/repository"
	"kis/internal/repository/vacancy"
	"reflect"
	"strconv"
	"strings"
)

type VacancyService struct {
	vacRepo    repository.Vacancy
	cusRepo    repository.Customer
	employRepo repository.Employee
}

func NewVacancyService(
	vacRepo repository.Vacancy,
	cusRepo repository.Customer,
	employRepo repository.Employee,
) *VacancyService {
	return &VacancyService{
		vacRepo:    vacRepo,
		cusRepo:    cusRepo,
		employRepo: employRepo,
	}
}

func (s *VacancyService) GetVacancies(params model.QueryDTO) ([]model.VacancyDTO, error) {
	searchTerms := vacancy.VacancySearchTermsDTO{}
	searchTerms.FromQueryDTO(&params, consts.PAGE_SIZE)

	models, err := s.vacRepo.GetBySearchTerms(searchTerms)
	if err != nil {
		return nil, err
	}

	//TODO microtask | refactor repository
	vacancyDtos := make([]model.VacancyDTO, len(models))
	for i, m := range models {
		v := &vacancyDtos[i]
		v.FromVacancy(&m)

		v.CandidatesCount = strconv.Itoa(len(m.Candidates))

		cus, err := s.cusRepo.GetById(m.CustomerID)
		if err != nil {
			return vacancyDtos, err
		}
		v.CustomerName = cus.Name

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

func (s *VacancyService) SearchVacancies(employeeId uint, query string) ([]model.VacancyDTO, int, error) {
	models, page, err := s.vacRepo.SearchVacancies(employeeId, query)
	vacancyDtos := make([]model.VacancyDTO, len(models))
	if err != nil {
		return nil, 0, fmt.Errorf("cannot search vacancies: %s", err.Error())
	}

	//TODO microtask | refactor repository
	for i, m := range models {
		v := &vacancyDtos[i]
		v.FromVacancy(&m)

		v.CandidatesCount = strconv.Itoa(len(m.Candidates))

		cus, err := s.cusRepo.GetById(m.CustomerID)
		if err != nil {
			return vacancyDtos, 1, err
		}
		v.CustomerName = cus.Name

		var grades []string
		for _, g := range m.Grades {
			if len(g.Name) > 0 {
				grades = append(grades, g.Name)
			}
		}
		v.Grade = strings.Join(grades, ", ")
	}
	return vacancyDtos, page, nil
}

// dto
// enum op
// generic value
func (s *VacancyService) UpdateVacancyField(id uint, op string, path string, value any) error {
	model, err := s.vacRepo.GetById(id)
	if err != nil {
		return fmt.Errorf("cannot update vacancy field: %s", err.Error())
	}

	switch v := value.(type) {
	case int:
		f := reflect.ValueOf(&model).Elem().FieldByName(path)
		if !f.CanSet() {
			return fmt.Errorf("cannot update vacancy field " + f.String())
		}
		f.SetInt(int64(v))
	case string:
		f := reflect.ValueOf(&model).Elem().FieldByName(path)
		if !f.CanSet() {
			return fmt.Errorf("cannot update vacancy field " + f.String())
		}
		if path == "Status" {
			tmp := v == "true"
			f.SetBool(tmp)

		} else {
			f.SetString(v)
		}
	case uint:
		f := reflect.ValueOf(&model).Elem().FieldByName(path)
		if !f.CanSet() {
			return fmt.Errorf("cannot update vacancy field " + f.String())
		}
		f.SetUint(uint64(v))
	default:
		return fmt.Errorf("cannot apply value")
	}

	err = s.vacRepo.Update(model)
	return err
}

func (s *VacancyService) GetVacancyById(id uint) (model.Vacancy, error) {
	return s.vacRepo.GetById(id)
}

func (s *VacancyService) CreateVacancy(vacancy model.Vacancy) (uint, error) {
	return s.vacRepo.Create(vacancy)
}
