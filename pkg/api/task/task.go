package task

type CreateRequest struct {
	Title       string `json:"title"`
	Description string `json:"description" binding:"required"`
	Type        string `json:"type" binding:"required"`
	Level       uint   `json:"level" binding:"required"`
	Timestamp   uint   `json:"timestamp" binding:"required"`
}
type CreateResponse struct {
	Id          uint
	Title       string
	Description string
	Type        string
	Level       string
	Timestamp   int64
}
type DeleteRequest struct {
	Id uint `uri:"id"`
}
type DeleteResponse struct {
	Id          uint
	Title       string
	Description string
	Type        string
	Level       string
	Timestamp   int64
}
type EditRequest struct {
	Id          uint   `uri:"id" binding:"required"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Type        string `json:"type"`
	Level       uint   `json:"level"`
}
type EditResponse struct {
	Id          uint
	Title       string
	Description string
	Type        string
	Level       string
	Timestamp   int64
}
type ReadRequest struct {
	Id    uint   `uri:"id"`
	Title string `form:"title,omitempty"`
	Type  string `form:"type,omitempty"`
	Level uint   `form:"level,omitempty"`
	Year  int    `form:"year,omitempty"`
	Month int    `form:"month,omitempty"`
	Day   int    `form:"day,omitempty"`
}

type ReadResponse struct {
	Id          uint
	Title       string
	Description string
	Type        string
	Level       string
	Timestamp   int64
}
