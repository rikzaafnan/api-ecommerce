package uploader

type FileResponse struct {
	ID           int    `json:"id"`
	FileName     string `json:"fileName"`
	FileSize     int    `json:"fileSize"`
	MIME         string `json:"MIME"`
	ActualMIME   string `json:"actualMIME"`
	UploadedFile string `json:"uploadedFile"`
	UploadedURL  string `json:"uploadedURL"`
}

func FormatFileUploader(fileAttachment Attachment, uploadURL string, actualMime, mime string, fileSize int) FileResponse {
	formatter := FileResponse{
		ID:           fileAttachment.ID,
		FileName:     fileAttachment.FileName,
		FileSize:     fileSize,
		MIME:         mime,
		ActualMIME:   actualMime,
		UploadedFile: fileAttachment.FileLocation,
		UploadedURL:  uploadURL,
	}

	return formatter
}
