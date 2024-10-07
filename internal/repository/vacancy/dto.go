package vacancy

import model "kis/internal/models"

type VacancySearchTermsDTO struct {
	Country    string `form:"country"`
	Region     string `form:"region"`
	City       string `form:"city"`
	Role       string `form:"role"`
	Status     bool   `form:"status"`
	SalaryFrom int    `form:"salary_from"`
	SalaryTo   int    `form:"salary_to"`
	Page       int    `form:"page"`
	PageSize   int    `form:"page_size"`
	NewFirst   bool   `form:"new_first"`
}

func (v *VacancySearchTermsDTO) FromQueryDTO(q *model.QueryDTO, pageSize int) {
	v.Country = q.Counry
	v.Region = q.Region
	v.City = q.City
	v.Role = q.Role
	v.Status = q.Status
	v.SalaryFrom = q.SalaryFrom
	v.SalaryTo = q.SalaryTo
	v.Page = q.Page
	v.NewFirst = q.NewFirst
	v.PageSize = pageSize
}
