package user

import "belscourrsego/internal/domain/common"

type ServicePasswordPolicy struct {
	hasher common.PasswordHasher
}

func NewServicePasswordPolicy(hasher common.PasswordHasher) *ServicePasswordPolicy {
	return &ServicePasswordPolicy{hasher: hasher}
}

func (s *ServicePasswordPolicy) HashService(password string, cost int) (string, error) {
	value, err := s.hasher.Hash(password, cost)
	if err != nil {
		return "", err
	}

	return value, nil
}

func (s *ServicePasswordPolicy) CompareService(hash, password string) error {
	return s.hasher.Compare(hash, password)
}
