package starshipcommsresolver

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/dig"
)

func TestTopSecretRoutes(t *testing.T) {
	// Crea una instancia de tu servidor
	server := &Server{
		container: dig.New(),
		// Puedes proporcionar otras configuraciones necesarias para el servidor en tu prueba.
	}

	// Crea un enrutador Gin para probar
	router := gin.New()

	// Configura la inyección de dependencias para topsecretHandler
	assert.NoError(t, server.container.Provide(func() *topsecretHandler {
		return &topsecretHandler{}
	}))

	// Llama a la función que estás probando
	err := server.TopSecretRoutes(router)

	// Verifica que no haya error al configurar las rutas
	assert.NoError(t, err)

	// Simula una solicitud POST a "/topsecret/"
	req := httptest.NewRequest(http.MethodPost, "/topsecret/", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Verifica que la respuesta sea un código 400 Bad Request
	assert.Equal(t, http.StatusBadRequest, w.Code)
}
