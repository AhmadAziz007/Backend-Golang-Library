package controller

import (
	"database/sql"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"library-synapsis/helper"
	"library-synapsis/model/domain"
	"library-synapsis/repository"
	"net/http"
)

type LoginController struct {
	UserRepo repository.UserManagementRepository
	RoleRepo repository.RoleManagementRepository
	DB       *sql.DB // Menambahkan db ke dalam struct
}

// Konstruktor yang menerima parameter *sql.DB
func NewLoginController(userRepo repository.UserManagementRepository, db *sql.DB) *LoginController {
	return &LoginController{
		UserRepo: userRepo,
		DB:       db, // Menyimpan koneksi database
	}
}

func (controller *LoginController) Login(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	var user domain.UserManagement
	ctx := request.Context() // Mengambil context dari request

	// Mulai transaksi baru dengan menggunakan db yang ada di struct LoginController
	tx, err := controller.DB.Begin() // Gunakan controller.DB di sini
	if err != nil {
		http.Error(writer, "Failed to start transaction", http.StatusInternalServerError)
		return
	}
	defer tx.Commit() // Jangan lupa untuk commit setelah selesai

	// Decode input request body
	err = json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		http.Error(writer, "Invalid input", http.StatusBadRequest)
		return
	}

	// Mendapatkan user berdasarkan email
	storedUser, err := controller.UserRepo.FindByEmail(ctx, tx, user.Email)
	if err != nil || storedUser.Password != user.Password {
		http.Error(writer, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	// Mendapatkan role berdasarkan RoleId
	storedRole, err := controller.RoleRepo.FindByRoleName(ctx, tx, storedUser.RoleName)
	if err != nil {
		http.Error(writer, "Role not found", http.StatusInternalServerError)
		return
	}

	// Menghasilkan JWT token
	token, err := helper.GenerateJWT(storedUser.UserId, storedRole.RoleName)
	if err != nil {
		http.Error(writer, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	// Mengirimkan response dengan token
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(map[string]string{"token": token})
}
