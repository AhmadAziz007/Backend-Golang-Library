package implement2

import (
	"context"
	"database/sql"
	"errors"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"library-synapsis/helper"
	"library-synapsis/model/web/create"
	"library-synapsis/model/web/response"
	"library-synapsis/repository"
)

type LoginServiceImpl struct {
	LoginRepository repository.LoginRepository
	DB              *sql.DB
	Validate        *validator.Validate
}

func NewLoginService(loginRepository repository.LoginRepository, db *sql.DB, validate *validator.Validate) *LoginServiceImpl {
	return &LoginServiceImpl{
		LoginRepository: loginRepository,
		DB:              db,
		Validate:        validate,
	}
}

// Update metode Login sesuai dengan signature interface LoginService
func (service *LoginServiceImpl) Login(ctx context.Context, loginRequest create.LoginCreateRequest) (response.LoginResponse, error) {
	var loginResponse response.LoginResponse

	// Validasi input request
	err := service.Validate.Struct(loginRequest)
	if err != nil {
		return loginResponse, err
	}

	// Mulai transaksi
	tx, err := service.DB.Begin()
	if err != nil {
		return loginResponse, err
	}
	defer tx.Rollback() // Pastikan rollback jika terjadi kesalahan

	// Cari user berdasarkan email
	storedUser, err := service.LoginRepository.FindByEmail(ctx, tx, loginRequest.Email, loginRequest.Password)
	if err != nil {
		return loginResponse, err
	}

	hashedPassword, err := helper.HashPassword(loginRequest.Password)
	if err != nil {
		// Handle error
	}
	loginRequest.Password = hashedPassword

	// Bandingkan password yang diberikan dengan password yang tersimpan
	err = bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(loginRequest.Password))
	if err != nil {
		return loginResponse, errors.New("invalid credentials")
	}

	// Ambil role berdasarkan RoleName
	storedRole, err := service.LoginRepository.FindByRoleName(ctx, tx, storedUser.RoleId)
	if err != nil {
		return loginResponse, err
	}

	// Generate JWT token
	token, err := helper.GenerateJWT(storedUser.UserId, storedRole.RoleName)
	if err != nil {
		return loginResponse, err
	}

	// Commit transaksi
	err = tx.Commit()
	if err != nil {
		return loginResponse, err
	}

	// Return login response dengan token
	loginResponse.Token = token
	return loginResponse, nil
}
