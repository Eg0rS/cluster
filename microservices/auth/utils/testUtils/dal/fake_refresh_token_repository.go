package dal

import "auth/dal"

func NewFakeRefreshTokenRepository() *FakeRefreshTokenRepository {
	return &FakeRefreshTokenRepository{make(map[string]string, 0)}
}

type FakeRefreshTokenRepository struct {
	store map[string]string
}

func (r *FakeRefreshTokenRepository) Save(token *dal.RefreshToken) error {
	r.store[token.Token] = token.UserId
	return nil
}

func (r *FakeRefreshTokenRepository) Get(token string, userId string) (*dal.RefreshToken, error) {
	userId, ok := r.store[token]
	if !ok {
		return nil, nil
	}
	return &dal.RefreshToken{
		UserId: userId,
		Token:  token,
	}, nil
}

func (r *FakeRefreshTokenRepository) Delete(token string, userId string) error {
	_, ok := r.store[token]
	if ok {
		delete(r.store, token)
	}
	return nil
}

func (r *FakeRefreshTokenRepository) DeleteByUserId(userId string) error {
	return nil
}

func (r *FakeRefreshTokenRepository) TokenExists(token string) bool {
	return false
}
