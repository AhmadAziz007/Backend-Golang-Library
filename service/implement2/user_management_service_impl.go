package implement2

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"library-synapsis/exception"
	"library-synapsis/helper"
	"library-synapsis/helper/model"
	"library-synapsis/model/domain"
	"library-synapsis/model/web/create"
	"library-synapsis/model/web/response"
	"library-synapsis/model/web/update"
	"library-synapsis/repository"
	"library-synapsis/service"
)

type UserManagementServiceImpl struct {
	UserRepository repository.UserManagementRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewUserManagementService(userRepository repository.UserManagementRepository, db *sql.DB, validate *validator.Validate) service.UserManagementService {
	return &UserManagementServiceImpl{
		UserRepository: userRepository,
		DB:             db,
		Validate:       validate,
	}
}

func (service *UserManagementServiceImpl) CreateUser(ctx context.Context, request create.UserManagementCreateRequest) response.UserManagementResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user := domain.UserManagement{
		RoleId:   request.RoleId,
		UserName: request.UserName,
		Email:    request.Email,
		Password: request.Password,
	}

	user = service.UserRepository.CreateUser(ctx, tx, user)

	return model.ToUserResponse(user)
}

func (service UserManagementServiceImpl) UpdateUser(ctx context.Context, request update.UserManagementUpdateRequest) response.UserManagementResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindByUserId(ctx, tx, request.UserId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	user.UserName = request.UserName
	user.RoleId = request.RoleId
	user.Email = request.Email
	user.Password = request.Password

	user = service.UserRepository.UpdateUser(ctx, tx, user)
	return model.ToUserResponse(user)
}

func (service UserManagementServiceImpl) DeleteUser(ctx context.Context, userId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindByUserId(ctx, tx, userId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.UserRepository.DeleteUser(ctx, tx, user)
}

func (service UserManagementServiceImpl) FindByUserId(ctx context.Context, userId int) response.UserManagementResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindByUserId(ctx, tx, userId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	return model.ToUserResponse(user)
}

func (service UserManagementServiceImpl) FindAllUser(ctx context.Context) []response.UserManagementResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user := service.UserRepository.FindAllUser(ctx, tx)
	return model.ToUserResponses(user)
}
