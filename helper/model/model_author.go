package model

import (
	"library-synapsis/model/domain"
	"library-synapsis/model/web/response"
)

func ToAuthorManagementResponse(authorManagement domain.AuthorManagement) response.AuthorManagementResponse {
	return response.AuthorManagementResponse{
		AuthorId:   authorManagement.AuthorId,
		AuthorName: authorManagement.AuthorName,
	}
}

func ToAuthorManagementResponses(authorManagement []domain.AuthorManagement) []response.AuthorManagementResponse {
	var authorManagementResponses []response.AuthorManagementResponse
	for _, authorManagement := range authorManagement {
		authorManagementResponses = append(authorManagementResponses, ToAuthorManagementResponse(authorManagement))
	}
	return authorManagementResponses
}
