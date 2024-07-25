package folders_test

import (
	"github.com/georgechieng-sc/interns-2022/folders"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	// "github.com/georgechieng-sc/interns-2022/folders"
	// "github.com/stretchr/testify/assert"
)

// MockDataFetcher is a mock implementation of DataFetcher for testing purposes.
type MockDataFetcher struct{}

// GetSampleData returns mock sample data for testing purposes.
func (MockDataFetcher) GetSampleData() ([]*folders.Folder, error) {
	return []*folders.Folder{
		{Id: uuid.Must(uuid.NewV4()), Name: "Folder1", OrgId: uuid.Must(uuid.FromString("c1556e17-b7c0-45a3-a6ae-9546248fb17a")), Deleted: false},
		{Id: uuid.Must(uuid.NewV4()), Name: "Folder2", OrgId: uuid.Must(uuid.FromString("c1556e17-b7c0-45a3-a6ae-9546248fb17a")), Deleted: true},
	}, nil
}

func Test_GetAllFolders(t *testing.T) {
	mockFetcher := MockDataFetcher{}

	t.Run("positive case", func(t *testing.T) {
		orgID := uuid.Must(uuid.FromString("c1556e17-b7c0-45a3-a6ae-9546248fb17a"))
		req := &folders.FetchFolderRequest{
			OrgID: orgID,
		}

		res, err := folders.GetAllFolders(req, mockFetcher)

		assert.NoError(t, err)
		assert.Len(t, res.Folders, 2)
	})

	t.Run("negative case - non-existent orgID", func(t *testing.T) {
		nonExistentOrgID := uuid.Must(uuid.NewV4())
		req := &folders.FetchFolderRequest{
			OrgID: nonExistentOrgID,
		}

		res, err := folders.GetAllFolders(req, mockFetcher)

		assert.NoError(t, err)
		assert.Len(t, res.Folders, 0)
	})

	t.Run("negative case - invalid orgID", func(t *testing.T) {
		invalidOrgID := uuid.Nil
		req := &folders.FetchFolderRequest{
			OrgID: invalidOrgID,
		}

		res, err := folders.GetAllFolders(req, mockFetcher)

		assert.NoError(t, err)
		assert.Len(t, res.Folders, 0)
	})
}
