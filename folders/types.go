package folders

import "github.com/gofrs/uuid"

type FetchFolderRequest struct {
	OrgID      uuid.UUID
	Page       int // Page number
	PageSize   int // Number of items per page
}

type FetchFolderResponse struct {
	Folders []*Folder
}