// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package sqlc

import (
	"context"
)

type Querier interface {
	DeclarationGetAll(ctx context.Context, arg DeclarationGetAllParams) ([]DeclarationGetAllRow, error)
	DeclarationGetAllCount(ctx context.Context, organizationID string) (int64, error)
	DeclarationGetOne(ctx context.Context, id string) (DeclarationGetOneRow, error)
	DeclarationInsertOne(ctx context.Context, arg DeclarationInsertOneParams) (string, error)
	LicenseGetAll(ctx context.Context, arg LicenseGetAllParams) ([]License, error)
	LicenseGetAllCount(ctx context.Context, arg LicenseGetAllCountParams) (int64, error)
	LicenseGetOne(ctx context.Context, id string) (License, error)
	LicenseInsertOne(ctx context.Context, arg LicenseInsertOneParams) (string, error)
	LicenseTypesGetAll(ctx context.Context) ([]LicenseType, error)
	OrganizationGetAll(ctx context.Context, arg OrganizationGetAllParams) ([]Organization, error)
	OrganizationGetAllCount(ctx context.Context, arg OrganizationGetAllCountParams) (int64, error)
	OrganizationGetOne(ctx context.Context, id string) (Organization, error)
	OrganizationInsertOne(ctx context.Context, arg OrganizationInsertOneParams) (string, error)
	PaymentGetAll(ctx context.Context, arg PaymentGetAllParams) ([]Payment, error)
	PaymentGetAllCount(ctx context.Context, arg PaymentGetAllCountParams) (int64, error)
	PaymentGetOne(ctx context.Context, id string) (Payment, error)
	PaymentInsertOne(ctx context.Context, arg PaymentInsertOneParams) (string, error)
	PaymentTypesGetAll(ctx context.Context) ([]PaymentType, error)
	PermissionGetOneByRoleAndPath(ctx context.Context, arg PermissionGetOneByRoleAndPathParams) (PermissionGetOneByRoleAndPathRow, error)
	StatusesGetAll(ctx context.Context) ([]Status, error)
	UserDeleteOne(ctx context.Context, id string) error
	UserGetAll(ctx context.Context, arg UserGetAllParams) ([]User, error)
	UserGetAllCount(ctx context.Context, arg UserGetAllCountParams) (int64, error)
	UserGetOne(ctx context.Context, id string) (User, error)
	UserGetOneByEmail(ctx context.Context, email string) (User, error)
	UserInsertOne(ctx context.Context, arg UserInsertOneParams) (string, error)
	UserUpdateOne(ctx context.Context, arg UserUpdateOneParams) error
}

var _ Querier = (*Queries)(nil)
