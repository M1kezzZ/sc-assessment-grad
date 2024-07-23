package folders

import (
	"github.com/gofrs/uuid"
)

// returns all folders for a given orgID.
func GetAllFolders(req *FetchFolderRequest) (*FetchFolderResponse, error) {

	r, err := FetchAllFoldersByOrgID(req.OrgID) // fetch all folders by the given orgID

	if err != nil { // check if there is an error
		return nil, err // return nil and the error
	}

	var ffr = &FetchFolderResponse{Folders: r} // create a FetchFolderResponse with the folder pointers
	return ffr, nil                             // return the FetchFolderResponse and nil for error
}

// This func fetches all folders matching the given orgID.
func FetchAllFoldersByOrgID(orgID uuid.UUID) ([]*Folder, error) {
	folders := GetSampleData() // get sample folder data

	resFolder := []*Folder{}         // empty slice to store result folders
	for _, folder := range folders { // iterate through the sample data
		if folder.OrgId == orgID { // check if the folder belongs to the given orgID
			resFolder = append(resFolder, folder) // append matching folders to the result slice
		}
	}
	return resFolder, nil // return the result slice and nil for error
}
