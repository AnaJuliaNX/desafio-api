package service

import (
	"desafio1/repositorio"
	"desafio1/tipos"
)

type DestService interface {
	SaveDest(tipos.Destinatario) tipos.Destinatario
	UpdateDest(dest tipos.Destinatario)
	FindAllDest() []tipos.Destinatario
	DeleteDest(dest tipos.Destinatario)
}

type destService struct {
	destrepositorio repositorio.DestRepositorio
}

func News(repos repositorio.DestRepositorio) DestService {
	return &destService{
		destrepositorio: repos,
	}
}

func (service *destService) SaveDest(dest tipos.Destinatario) tipos.Destinatario {
	service.destrepositorio.SaveDest(dest)
	return dest
}

func (service *destService) UpdateDest(dest tipos.Destinatario) {
	service.destrepositorio.UpdateDest(dest)
}

func (service *destService) FindAllDest() []tipos.Destinatario {
	return service.destrepositorio.FindAllDest()
}

func (service *destService) DeleteDest(dest tipos.Destinatario) {
	service.destrepositorio.DeleteDest(dest)
}
