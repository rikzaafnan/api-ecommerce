package uploader

import "errors"

type ServiceAttachment interface {
	Save(input AttachmentInput) (Attachment, error)
	Update(attachmentID int, input AttachmentInput) (Attachment, error)
	DeleteByID(attachmentID int) error
}

type serviceAttachmentImpl struct {
	repositoryAttachment RepositoryAttachment
}

func NewServiceAttachment(repositoryAttachment RepositoryAttachment) *serviceAttachmentImpl {
	return &serviceAttachmentImpl{repositoryAttachment}
}

func (s serviceAttachmentImpl) Save(input AttachmentInput) (Attachment, error) {
	// TODO implement me

	var attachment Attachment
	attachment.FileLocation = input.FileLocation
	attachment.Module = input.Module
	attachment.FileExtension = input.FileExtension
	attachment.FileName = input.FileName
	attachment.UserID = input.UserID

	attachment, err := s.repositoryAttachment.Save(attachment)
	if err != nil {
		return attachment, errors.New("failed create attachmetn")
	}

	return attachment, nil

}

func (s serviceAttachmentImpl) Update(attachmentID int, input AttachmentInput) (Attachment, error) {
	// TODO implement me
	panic("implement me")
}

func (s serviceAttachmentImpl) DeleteByID(attachmentID int) error {
	err := s.DeleteByID(attachmentID)

	if err != nil {
		return errors.New("failed delete attachment")

	}

	return nil
}
