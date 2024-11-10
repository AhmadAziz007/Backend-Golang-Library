package model

import (
	"library-synapsis/model/domain"
	"library-synapsis/model/web/response"
)

func ToCategoryResponse(category domain.Category) response.CategoryResponse {
	return response.CategoryResponse{
		CategoryId:   category.CategoryId,
		CategoryName: category.CategoryName,
	}
}

func ToCategoryResponses(categories []domain.Category) []response.CategoryResponse {
	var categoryResponses []response.CategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, ToCategoryResponse(category))
	}
	return categoryResponses
}
