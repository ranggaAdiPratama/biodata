package api

import (
	"testing"

	db "github.com/ranggaAdiPratama/go_biodata/db/sqlc"
	"github.com/ranggaAdiPratama/go_biodata/util"
	"github.com/stretchr/testify/require"
)

func randomUser(t *testing.T) (user db.User, password string) {
	password = util.RandomString(6)

	hashedPassword, err := util.HashPassword(password)

	require.NoError(t, err)

	user = db.User{
		Name:     util.RandomString(10),
		Password: hashedPassword,
		Username: util.RandomUsername(),
		Email:    util.RandomEmail(),
	}

	return
}
