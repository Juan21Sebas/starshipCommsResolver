package ports

type CommunicationServices interface {
	GetLocation(...float32) (x, y float32)
	GetMessage(...[]string) (msg string)
}
