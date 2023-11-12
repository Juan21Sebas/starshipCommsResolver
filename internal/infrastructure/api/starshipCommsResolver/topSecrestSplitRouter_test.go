package starshipcommsresolver

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/dig"
)

func TestTopSecretSplitRoutes(t *testing.T) {
	// Crea una instancia de tu servidor
	server := &Server{
		container: dig.New(),
		// Puedes proporcionar otras configuraciones necesarias para el servidor en tu prueba.
	}

	// Proporciona manualmente una instancia de topsecretSplitHandler al contenedor
	server.container.Provide(func() *topsecretSplitHandler {
		return &topsecretSplitHandler{} // Aquí inicializa la estructura según tus necesidades
	})

	// Crea un enrutador Gin para probar
	router := gin.New()

	// Llama a la función que estás probando
	err := server.TopSecretSplitRoutes(router)

	// Verifica que no haya error al configurar las rutas
	assert.NoError(t, err)

	req := httptest.NewRequest(http.MethodPost, "/topsecret_split/satellite1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Verifica que la respuesta sea un código 400 Bad Request
	assert.Equal(t, http.StatusBadRequest, w.Code)

	req = httptest.NewRequest(http.MethodGet, "/topsecret_split/satellite1", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Verifica que la respuesta sea un código 404 Not Found
	assert.Equal(t, http.StatusNotFound, w.Code)
}
