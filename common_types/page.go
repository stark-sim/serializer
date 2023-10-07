package common_types

type PaginateReq struct {
	PageIndex int `form:"page_index" json:"page_index" binding:"required"`
	PageSize  int `form:"page_size" json:"page_size" binding:"required"`
}

type PaginateResp struct {
	PageIndex int         `json:"page_index"`
	PageSize  int         `json:"page_size"`
	Total     int         `json:"total"`
	List      interface{} `json:"list"`
}
