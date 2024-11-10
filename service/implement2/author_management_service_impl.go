package implement2

import (
	"context"
	"database/sql"
	"library-synapsis/exception"
	"library-synapsis/helper"
	"library-synapsis/helper/model"
	"library-synapsis/model/domain"
	"library-synapsis/model/web/create"
	"library-synapsis/model/web/response"
	"library-synapsis/model/web/update"
	"library-synapsis/repository"
	"library-synapsis/service"

	"github.com/go-playground/validator/v10"
)

type AuthorManagementServiceImpl struct {
	AuthorManagementRepository repository.AuthorManagementRepository
	DB                         *sql.DB
	Validate                   *validator.Validate
}

func NewAuthorManagementService(authorManagementRepository repository.AuthorManagementRepository, DB *sql.DB, validate *validator.Validate) service.AuthorManagementService {
	return &AuthorManagementServiceImpl{
		AuthorManagementRepository: authorManagementRepository,
		DB:                         DB,
		Validate:                   validate}
}

func (service *AuthorManagementServiceImpl) CreateAuthor(ctx context.Context, request create.AuthorManagementCreateRequest) response.AuthorManagementResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	author := domain.AuthorManagement{
		AuthorName: request.AuthorName,
	}

	author = service.AuthorManagementRepository.CreateAuthor(ctx, tx, author)

	return model.ToAuthorManagementResponse(author)
}

func (service *AuthorManagementServiceImpl) UpdateAuthor(ctx context.Context, request update.AuthorManagementUpdateRequest) response.AuthorManagementResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	author, err := service.AuthorManagementRepository.FindByAuthorId(ctx, service.DB, request.AuthorId)
	helper.PanicIfError(err)

	author.AuthorName = request.AuthorName

	author = service.AuthorManagementRepository.UpdateAuthor(ctx, tx, author)

	return model.ToAuthorManagementResponse(author)
}

func (service *AuthorManagementServiceImpl) DeleteAuthor(ctx context.Context, authorId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	author, err := service.AuthorManagementRepository.FindByAuthorId(ctx, service.DB, authorId)
	helper.PanicIfError(err)

	service.AuthorManagementRepository.DeleteAuthor(ctx, tx, author)
}

func (service *AuthorManagementServiceImpl) FindByAuthorId(ctx context.Context, authorId int) response.AuthorManagementResponse {

	author, err := service.AuthorManagementRepository.FindByAuthorId(ctx, service.DB, authorId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return model.ToAuthorManagementResponse(author)
}

func (service *AuthorManagementServiceImpl) FindAllAuthor(ctx context.Context) []response.AuthorManagementResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	authors := service.AuthorManagementRepository.FindAllAuthor(ctx, tx)
	return model.ToAuthorManagementResponses(authors)
}
