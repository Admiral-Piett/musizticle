package daos

import (
	"github.com/Admiral-Piett/musizticle/app/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GetUser_success(t *testing.T) {
	dao := resetDao()
	seededId := seedUser("HanSolo", "MillenniumFalconsRock", dao)

	user, err := dao.GetUser("HanSolo", "MillenniumFalconsRock")

	assert.Nil(t, err)
	assert.Equal(t, int(seededId), user.Id)
	assert.Equal(t, "HanSolo", user.Username)
}

func Test_GetUser_no_user_returns_error(t *testing.T) {
	dao := resetDao()

	user, err := dao.GetUser("HanSolo", "MillenniumFalconsRock")

	assert.Error(t, err)
	assert.Equal(t, models.User{}, user)
}
