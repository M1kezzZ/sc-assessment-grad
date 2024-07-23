package folders

import (
	"github.com/gofrs/uuid"
)

// returns all folders for a given orgID.
func PagGetAllFolders(req *FetchFolderRequest) (*FetchFolderResponse, error) {

	r, err := PagFetchAllFoldersByOrgID(req.OrgID)

	if err != nil { // check if there is an error
		return nil, err // return nil and the error
	}
	// pagination
	start := req.Page * req.PageSize
	end := start + req.PageSize

	if start > len(r) {
		start = len(r)
	}
	if end > len(r) {
		end = len(r)
	}

	paginationFolder := r[start:end]

	var ffr = &FetchFolderResponse{Folders: paginationFolder}
	return ffr, nil
}

// This func fetches all folders matching the given orgID.
func PagFetchAllFoldersByOrgID(orgID uuid.UUID) ([]*Folder, error) {
	folders := GetSampleData()

	resFolder := []*Folder{}
	for _, folder := range folders {
		if folder.OrgId == orgID {
			resFolder = append(resFolder, folder)
		}
	}
	return resFolder, nil
}
