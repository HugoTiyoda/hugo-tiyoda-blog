package dtos

import "blog/application/domain"

type RegisterAuthorRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Bio      string `json:"bio"`
	Password string `json:"password"`
}

func (d *RegisterAuthorRequest) ToAuthor() *domain.Author {
	return &domain.Author{
		Name:  d.Name,
		Email: d.Email,
		Bio:   d.Bio,
	}
}

type RegisterAuthorResponse struct {
	Author *domain.Author `json:"author"`
	Token  string         `json:"token"`
}

func ToRegisterAuthorResponse(a *domain.Author, token string) *RegisterAuthorResponse {
	return &RegisterAuthorResponse{
		Author: a,
		Token:  token,
	}
}
