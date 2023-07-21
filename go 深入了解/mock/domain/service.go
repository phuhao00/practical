package domain

import "mocktest/entity"

type Service struct {
	person entity.Person
}

func (s *Service) GetPersonName() string {
	return s.person.GetName()
}
