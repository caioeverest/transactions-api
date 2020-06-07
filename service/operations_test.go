package service

import (
	"testing"

	"github.com/caioeverest/transactions-api/model"
	"github.com/caioeverest/transactions-api/util"
	"github.com/stretchr/testify/assert"
)

func TestOperationsService_RecoverDescriptionShouldReturnErrorIfNotFound(t *testing.T) {
	beforeEach()
	notfound := util.Error("not found")
	repoMock.On("FindById", COMPRA_A_VISTA, &model.Operation{}).Return(notfound)
	desc, err := opTarget.RecoverDescription(COMPRA_A_VISTA)
	assert.Equal(t, "", desc)
	assert.Equal(t, notfound, err)
}

func TestOperationsService_RecoverDescriptionShouldReturnSuccessIfOperationExits(t *testing.T) {
	beforeEach()
	repoMock.On("FindById", COMPRA_A_VISTA, &model.Operation{}).Return(nil)
	desc, err := opTarget.RecoverDescription(COMPRA_A_VISTA)
	assert.Equal(t, "", desc)
	assert.Nil(t, err)
}
