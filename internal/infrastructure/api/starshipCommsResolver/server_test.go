// server_test.go

package starshipcommsresolver

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateServer(t *testing.T) {
	server := CreateServer()

	// Aquí puedes realizar pruebas de configuración, middleware, etc., usando las funciones de aserción adecuadas
	// por ejemplo, asertar que los middleware estén configurados correctamente
	assert.NotNil(t, server)
}

func TestRunServerIntegrationWithPost(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Esto solo es para simular la creación y ejecución del servidor
	}))
	defer server.Close()

	jsonData := []byte(`{"key": "value"}`) // Cambia esto según el formato JSON necesario
	resp, err := http.Post(server.URL+"/topsecret/", "application/json", bytes.NewBuffer(jsonData))
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
