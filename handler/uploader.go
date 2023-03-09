package handler

import (
	common "api-ecommerce/common/uploader"
	"api-ecommerce/config"
	"api-ecommerce/helper"
	"api-ecommerce/user"
	"fmt"
	"github.com/gabriel-vasile/mimetype"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"path/filepath"
	"strconv"
	"time"
)

type uploaderHandler struct {
	attachmentService common.ServiceAttachment
}

func NewUploaderhandler(attachmentService common.ServiceAttachment) *uploaderHandler {
	return &uploaderHandler{attachmentService}
}

func (h *uploaderHandler) Save(c *gin.Context) {

	name := c.PostForm("name")
	module := c.PostForm("module")
	email := c.PostForm("email")

	log.Info("name : ", name)
	log.Info("email : ", email)

	// Source
	file, err := c.FormFile("file")
	if err != nil {

		c.JSON(http.StatusBadRequest, err.Error())
		return

	}

	dataFile, err := file.Open()
	if err != nil {

		c.JSON(http.StatusBadRequest, err.Error())
		return

	}

	currentUser := c.MustGet("currentUser").(user.User)

	m, err := validateUploader(currentUser.ID, module, dataFile)
	if err != nil {

		c.JSON(http.StatusBadRequest, err.Error())
		return

	}

	filename := filepath.Base(file.Filename)
	buildFileName := fmt.Sprintf("%d-%s", time.Now().UnixNano(), filename)
	pathDestination := "upload-files/images/" + module + "/" + buildFileName

	if err := c.SaveUploadedFile(file, pathDestination); err != nil {

		c.JSON(http.StatusBadRequest, err.Error())
		return

	}

	var input common.AttachmentInput
	input.FileName = buildFileName
	input.Module = module
	input.UserID = currentUser.ID
	input.FileExtension = m.Extension()
	input.FileLocation = "images/" + module + "/" + buildFileName

	attachment, err := h.attachmentService.Save(input)
	if err != nil {

		c.JSON(http.StatusBadRequest, err.Error())
		return

	}

	uploadFileURL := config.BASE_URL + "/images/" + module + "/" + buildFileName

	fileResponse := common.FormatFileUploader(attachment, uploadFileURL, m.Extension(), file.Header.Get("Content-Type"), int(file.Size))

	response := helper.APIResponse("success send picture", http.StatusOK, "success", fileResponse)

	c.JSON(http.StatusOK, response)

}

func (h *uploaderHandler) Deleted(c *gin.Context) {

	// formatter := transaction.FormatTransaction(newTransaction, newTransaction.TransactionDetails)

	// response := helper.APIResponse("transaction has been created", http.StatusOK, "success", formatter)

	// c.JSON(http.StatusOK, response)

	paramUplaoderFIleID := c.Param("uploaderID")
	attachmentID, _ := strconv.Atoi(paramUplaoderFIleID)

	err := h.attachmentService.DeleteByID(attachmentID)
	if err != nil {

		response := helper.APIResponse("delete attachment suuccess", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return

	}

	response := helper.APIResponse("product has been deleted", http.StatusOK, "success", nil)

	c.JSON(http.StatusOK, response)

}

func validateUploader(userID int, module string, file io.Reader) (*mimetype.MIME, error) {
	mType, err := mimetype.DetectReader(file)
	if err != nil {
		log.Error(err)
		return mType, err
	}

	log.Debug("got module: ", module)
	log.Debug("actual mime type: ", mType)
	log.Debug("passing arguments to validate:", userID, " ", module)

	exampleModuleReaady := []string{"user", "product", "category"}

	if !helper.IsIn(module, exampleModuleReaady) {
		err := errors.New("Error module tidak tersedia")
		return mType, errors.Wrap(err, fmt.Sprintf("got: %s. expected: %v", module, exampleModuleReaady))
	}

	if err != nil {
		log.Error(err, "got: ", mType)
		return mType, errors.Wrap(err, fmt.Sprintf("got %s", mType.String()))
	}

	return mType, nil

}
