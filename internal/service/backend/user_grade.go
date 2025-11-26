package backend

import (
	"context"
	"go-service/api/backendRoute"
)

type (
	IUserGrade interface {
		LIndex(ctx context.Context, req *backendRoute.UserGradesReq) (res interface{}, err error)
		LSaveUserGrades(ctx context.Context, req *backendRoute.SaveUserGradesReq) (err error)
		LDeleteUserGrade(ctx context.Context, req *backendRoute.DeleteUserGradeReq) (err error)
	}
)

var localUserGrade IUserGrade

func ServiceUserGrade() IUserGrade {
	return localUserGrade
}

func RegisterUserGrade(p IUserGrade) {
	localUserGrade = p
}
