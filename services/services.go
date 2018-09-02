package services

import "github.com/boniface/gonosqltests/domain"

type Services interface {
	create() domain.Results
	read() domain.Results
	readAll() domain.Results
	delete() domain.Results
	update() domain.Results
}
