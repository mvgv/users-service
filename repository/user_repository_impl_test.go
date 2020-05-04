package repository

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*TestGetUserByIDUserNotFound teste unitario para um usuario nao encontrado*/
func TestGetUserByIDUserNotFound(t *testing.T) {
	user, err := GetUserByID(0)
	assert.Nil(t, user, "Id nao esperado")
	assert.NotNil(t, err)

	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "User not found", err.Message)
}

/*TestGetUserByIDUserMatches teste unitario para um usuario encontrado*/
func TestGetUserByIDUserMatches(t *testing.T) {
	user, err := GetUserByID(123)
	assert.Nil(t, err)
	assert.NotNil(t, user)

	assert.EqualValues(t, 123, user.ID)
	assert.EqualValues(t, "Marcus", user.Name)
}

/*BenchmarkGetUserByIDUserNotFound teste de benchmark para um usuario nao encontrado*/
func BenchmarkGetUserByIDUserNotFound(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetUserByID(0)
	}
}
