package service

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLocation(t *testing.T) {
	s := NewService(nil)

	// // Caso de prueba con distancias válidas
	x, y := s.GetLocation(100, 142.14, 125)
	fmt.Printf("Calculated: x = %f, y = %f\n", x, y)

	// Caso de prueba con menos de 3 distancias (debería devolver 0, 0)
	x, y = s.GetLocation(100, 142.14)
	assert.Equal(t, float32(0), x)
	assert.Equal(t, float32(0), y)
}

func TestGetMessage(t *testing.T) {
	s := NewService(nil)

	// Caso de prueba con mensajes válidos
	message := s.GetMessage([]string{"este", "", "", "mensaje", ""}, []string{"", "es", "", "", "secreto"})
	assert.Equal(t, "este es  mensaje secreto", message)

	// Caso de prueba con mensajes vacíos
	message = s.GetMessage([]string{"", "", "", ""}, []string{"", "", "", ""})
	assert.Equal(t, "    ", message)
}
