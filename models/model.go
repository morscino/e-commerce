package models

// ResponseObject structure for the response object
type ResponseObject struct {
	Code    int         `json:"-"`
	Status  string      `json:"status,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

// PagingInfo is the pagination info structure
type PagingInfo struct {
	TotalCount  int  `json:"totalCount"`
	Page        int  `json:"page"`
	HasNextPage bool `json:"hasNextPage"`
	Count       int  `json:"count"`
}

// APIPagingDto is the pagination data transfer object
type APIPagingDto struct {
	Limit     int      `json:"limit,omitempty"`
	Sort      string   `json:"sort,omitempty"`
	Direction string   `json:"direction,omitempty"`
	Select    []string `json:"select,omitempty"`
	Filter    string   `json:"filter,omitempty"`
	Page      int      `json:"page,omitempty"`
}
