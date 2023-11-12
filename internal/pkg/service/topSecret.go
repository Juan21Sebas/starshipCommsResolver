package service

import (
	"strings"

	"starshipCommsResolver/internal/pkg/ports"
)

type Service struct {
}

func NewService(port ports.CommunicationServices) *Service {
	return &Service{}
}

func (s *Service) GetLocation(distances ...float32) (x, y float32) {
	if len(distances) < 3 {
		return 0, 0
	}

	// Coordenadas de los satélites
	kenobiCoords := [2]float32{-500.0, -200.0}
	skywalkerCoords := [2]float32{100.0, -100.0}
	satoCoords := [2]float32{500.0, 100.0}

	// Distancias a los satélites
	distanceKenobi := float64(distances[0])
	distanceSkywalker := float64(distances[1])
	distanceSato := float64(distances[2])

	// Cálculos auxiliares
	emisorX := float32(0) // Inicializar con un valor temporal
	emisorY := float32(0) // Inicializar con un valor temporal

	dx1 := kenobiCoords[0] - emisorX
	dy1 := kenobiCoords[1] - emisorY
	dx2 := skywalkerCoords[0] - emisorX
	dy2 := skywalkerCoords[1] - emisorY
	dx3 := satoCoords[0] - emisorX
	dy3 := satoCoords[1] - emisorY

	// Fórmulas de trilateración
	A := float64(2 * dx2)
	B := float64(2 * dy2)
	C := float64(distanceKenobi*distanceKenobi - distanceSkywalker*distanceSkywalker - float64(dx1*dx1) - float64(dy1*dy1) + float64(dx2*dx2) + float64(dy2*dy2))
	D := float64(2 * dx3)
	E := float64(2 * dy3)
	F := float64(distanceKenobi*distanceKenobi - distanceSato*distanceSato - float64(dx1*dx1) - float64(dy1*dy1) + float64(dx3*dx3) + float64(dy3*dy3))

	// Calcular las coordenadas del emisor del mensaje
	denominator := B*D - A*E
	if denominator == 0 {
		return 0, 0
	}

	emisorX = float32((C*E - F*B) / denominator)
	emisorY = float32((C*D - A*F) / denominator)

	return emisorX, emisorY
}

func (s *Service) GetMessage(messages ...[]string) string {

	adjustedMessage := make([]string, 5)
	// Itera sobre cada palabra en los mensajes y agrega al mensaje ajustado
	// si la palabra no está en blanco, teniendo en cuenta el desfase
	for _, m := range messages {
		for i, world := range m {
			if adjustedMessage[i] == "" {
				adjustedMessage[i] = world
			}
		}
	}

	// Convierte el mensaje ajustado en una cadena
	finalMessage := strings.Join(adjustedMessage, " ")
	return finalMessage
}
