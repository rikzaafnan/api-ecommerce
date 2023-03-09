package uploader

import "gorm.io/gorm"

type RepositoryAttachment interface {
	Save(attachment Attachment) (Attachment, error)
	Update(attachmentID int, attachment Attachment) (Attachment, error)
	DeleteByID(attachmentID int) error
}

type repositoryAttachmentImpl struct {
	db *gorm.DB
}

func NewRepositoryAttachment(db *gorm.DB) *repositoryAttachmentImpl {
	return &repositoryAttachmentImpl{db}
}

func (r repositoryAttachmentImpl) Save(attachment Attachment) (Attachment, error) {
	// TODO implement me

	err := r.db.Create(&attachment).Error

	if err != nil {
		return attachment, err
	}

	return attachment, nil

}

func (r repositoryAttachmentImpl) Update(attachmentID int, attachment Attachment) (Attachment, error) {

	var attachmentModel Attachment

	err := r.db.Where("id = ?", attachmentID).First(&attachmentModel).Error
	if err != nil {

		return attachment, err
	}

	err = r.db.Save(&attachment).Error
	if err != nil {
		return attachment, err
	}

	return attachment, nil
}

func (r repositoryAttachmentImpl) DeleteByID(attachmentID int) error {
	var attachment Attachment

	err := r.db.Where("id = ?", attachmentID).Delete(&attachment).Error
	if err != nil {

		return err
	}

	return nil
}
