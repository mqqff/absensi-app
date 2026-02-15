package dto

type PaginationResponse struct {
	TotalData int64 `json:"total_data"`
	TotalPage int   `json:"total_page"`
	Page      int   `json:"page"`
	Limit     int   `json:"limit"`
}

type PaginationRequest struct {
	Page   int
	Limit  int
	Offset int
}

func NewPagination(page, limit int) PaginationRequest {
	if limit < 1 {
		limit = 10
	}

	if page < 1 {
		page = 1
	}

	return PaginationRequest{
		Page:   page,
		Limit:  limit,
		Offset: (page - 1) * limit,
	}
}

func NewPaginationResponse(totalData int64, page, limit int) PaginationResponse {
	totalPage := int(totalData) / limit
	if int(totalData)%limit != 0 {
		totalPage++
	}

	return PaginationResponse{
		TotalData: totalData,
		TotalPage: totalPage,
		Page:      page,
		Limit:     limit,
	}
}
