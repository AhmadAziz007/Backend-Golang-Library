package model

import (
	"library-synapsis/model/domain"
	"library-synapsis/model/web/response"
)

func ToUserResponse(user domain.UserManagement) response.UserManagementResponse {
	return response.UserManagementResponse{
		UserId:   user.UserId,
		RoleId:   user.RoleId,
		UserName: user.UserName,
		Email:    user.Email,
		Password: user.Password,
	}
}

func ToUserResponses(users []domain.UserManagement) []response.UserManagementResponse {
	var userResponses []response.UserManagementResponse
	for _, user := range users {
		userResponses = append(userResponses, ToUserResponse(user))
	}
	return userResponses
}
