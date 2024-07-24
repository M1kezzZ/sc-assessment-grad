package folders_test

import (
	"testing"

	"github.com/georgechieng-sc/interns-2022/folders"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_GetAllFolders(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		// Prepare a sample request
		req := &folders.FetchFolderRequest{
			OrgID: uuid.FromStringOrNil(folders.DefaultOrgID),
		}

		// Call the GetAllFolders function
		resp, err := folders.GetAllFolders(req)
		assert.NoError(t, err)

		// Check if the response is not nil
		assert.NotNil(t, resp)

		// Check if the response contains folders
		assert.NotEmpty(t, resp.Folders)

		// Verify if the folders belong to the given OrgID
		for _, folder := range resp.Folders {
			assert.Equal(t, req.OrgID, folder.OrgId)
		}
	})

	// pagination test
	t.Run("test", func(t *testing.T) {
		// Prepare a sample request with pagination
		req := &folders.FetchFolderRequest{
			OrgID:    uuid.FromStringOrNil(folders.DefaultOrgID),
			Page:     0,
			PageSize: 3,
		}

		// Call the GetAllFolders function
		resp, err := folders.PagGetAllFolders(req)
		assert.NoError(t, err)

		// Check if the response is not nil
		assert.NotNil(t, resp)

		// Check if the response contains the correct number of folders
		assert.Len(t, resp.Folders, 3) // Adjust based on sample data

		// Verify if the folders belong to the given OrgID
		for _, folder := range resp.Folders {
			assert.Equal(t, req.OrgID, folder.OrgId)
		}
	})
}
