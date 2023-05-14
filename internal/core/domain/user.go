package domain

import "gitlab.com/greatsoft/xif-backend/internal/core/repository/psql/sqlc"

type UserCreateParams struct {
	Email          string `json:"email"`
	Inn            string `json:"inn"`
	HashedPassword string `json:"hashed_password"`
	PhoneNumber    string `json:"phone_number"`
	RegionID       int32  `json:"region_id"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	SecondName     string `json:"second_name"`
	ID             string `json:"id" swaggerignore:"true"`
}

type Users struct {
	Objects []sqlc.User `json:"objects"`
	Count   int32       `json:"count"`
}

type GetAllParams struct {
	Search         string `form:"search" json:"search"`
	RegionID       int32  `form:"region_id" json:"region_id"`
	Offset         int32  `form:"offset" json:"offset"`
	Limit          int32  `form:"limit" json:"limit" default:"10"`
	OrganizationID string `form:"organization_id" json:"organization_id"`
}
