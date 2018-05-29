package checkers

import "github.com/kiali/kiali/services/models"

type Checker interface {
	Check() ([]*models.IstioCheck, bool)
}
