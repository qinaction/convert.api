package request

type ConvertPdf struct {
	Url     string `json:"url" binding:"required"`
	FileType string  `json:"fileType" binding:"required"`
}
