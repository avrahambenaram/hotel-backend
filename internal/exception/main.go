package exception

type Exception struct {
	Message string
	Status  uint
}

func New(message string, status uint) *Exception {
	return &Exception{
		message,
		status,
	}
}
