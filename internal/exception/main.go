package exception

type Exception struct {
	Message string
	Status  int
}

func New(message string, status int) *Exception {
	return &Exception{
		message,
		status,
	}
}
