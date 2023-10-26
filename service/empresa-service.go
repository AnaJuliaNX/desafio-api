package service

import (
	"desafio1/repositorio"
	"desafio1/tipos"
)

type EmpService interface {
	SaveEmp(tipos.Empresa) tipos.Empresa
	UpdateEmp(emp tipos.Empresa)
	DeleteEmp(emp tipos.Empresa)
}

type empService struct {
	emprepositorio repositorio.EmpRepositorio
}

func NeW(rep repositorio.EmpRepositorio) EmpService {
	return &empService{
		emprepositorio: rep,
	}
}

func (service *empService) SaveEmp(emp tipos.Empresa) tipos.Empresa {
	service.emprepositorio.SaveEmp(emp)
	return emp
}

func (service *empService) UpdateEmp(emp tipos.Empresa) {
	service.emprepositorio.UpdateEmp(emp)
}

func (service *empService) DeleteEmp(emp tipos.Empresa) {
	service.emprepositorio.DeleteEmp(emp)
}
