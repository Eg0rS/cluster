package microservices

import (
	"auth/microservices"
	"fmt"
)

func NewFakePasswordHasher() microservices.PasswordHasherService {
	return &FakePasswordHasherClient{}
}

type FakePasswordHasherClient struct{}

func (c *FakePasswordHasherClient) Get(password string) (string, error) {
	return fmt.Sprintf("HASHED_%s", password), nil
}

func (c *FakePasswordHasherClient) IsValid(hashedPassword string, password string) (bool, error) {
	hp, _ := c.Get(password)
	return hashedPassword == hp, nil
}
