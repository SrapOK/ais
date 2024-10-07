package model

type QueryDTO struct {
	Employee_id int    `form:"employee_id"`
	Counry      string `form:"country"`
	Region      string `form:"region"`
	City        string `form:"city"`
	Role        string `form:"role"`
	Status      bool   `form:"status"`
	Grade       string `form:"grade"`
	SalaryFrom  int    `form:"salary_from"`
	SalaryTo    int    `form:"salary_to"`
	Skills      string `form:"skills"`
	Mine        bool   `form:"mine"`
	Favs        bool   `form:"favs"`
	Page        int    `form:"page"`
	NewFirst    bool   `form:"new_first"`
}

type VacancyDTO struct {
	Title           string `json:"title" validate:"required"`
	Role            string `json:"role" validate:"required"`
	Grade           string `json:"grade" validate:"required"`
	Status          bool   `json:"status" validate:"required"`
	CustomerName    string `json:"customer_name" validate:"required"`
	CandidatesCount string `json:"candidates_count" validate:"required"`
	InFav           bool   `json:"in_fav" validate:"required"`
}

func (v *VacancyDTO) FromVacancy(m *Vacancy) {
	v.Title = m.Title
	v.Role = m.RoleName
	v.Status = m.IsActive
}
