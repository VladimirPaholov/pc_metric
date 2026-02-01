package service

import (
	"errors"
	"time"
)

type MetricRepository interface {
	AddMetricDB(timestamp time.Time, message string) error
}
type UserService struct {
	repo MetricRepository
}

func NewUserService(repo MetricRepository) *UserService {
	return &UserService{repo: repo}
}
func (u *UserService) AddMetricDB(createdAt time.Time, message string) error {
	if message == "" {
		return errors.New("message if empty")
	}
	return u.repo.AddMetricDB(time.Now(), message)
}
