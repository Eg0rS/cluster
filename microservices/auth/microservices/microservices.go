package microservices

type PasswordHasherService interface {
	Get(password string) (string, error)
	IsValid(hashedPassword string, password string) (bool, error)
}
