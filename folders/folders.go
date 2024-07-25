package folders

import (
	"github.com/gofrs/uuid"
)

// GetAllFolders processes the request to fetch all folders for a given organization ID.
// It returns a FetchFolderResponse containing a list of folders and the errors encountered.
func GetAllFolders(req *FetchFolderRequest, fetcher DataFetcher) (*FetchFolderResponse, error) {
	// Fetch all folders by organization ID
	folders, err := FetchAllFoldersByOrgID(req.OrgID, fetcher)
	if err != nil {
		return nil, err
	}

	// Construct the FetchFolderResponse
	res := &FetchFolderResponse{
		Folders: folders,
	}

	return res, nil
}

// FetchAllFoldersByOrgID fetches all folders for a given organization ID.
// It returns the pointers to the folders and any errors encountered.
func FetchAllFoldersByOrgID(orgID uuid.UUID, fetcher DataFetcher) ([]*Folder, error) {
	// Get sample data from the JSON file
	folders, err := fetcher.GetSampleData()
	if err != nil {
		return nil, err
	}

	// Filter the folders by organization ID
	resFolder := []*Folder{}
	for _, folder := range folders {
		if folder.OrgId == orgID {
			resFolder = append(resFolder, folder)
		}
	}

	return resFolder, nil
}
