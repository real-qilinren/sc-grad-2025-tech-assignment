package folders

import "github.com/gofrs/uuid"

// FetchFolderRequest represents the request to fetch all folders for a given organization ID.
type FetchFolderRequest struct {
	OrgID uuid.UUID
}

// FetchFolderRequestPag represents the request to fetch all folders for a given organization ID with pagination.
type FetchFolderRequestPag struct {
	OrgID    uuid.UUID
	Token    string
	PageSize int
}

// FetchFolderResponse represents the response that containing the list of folders.
type FetchFolderResponse struct {
	Folders []*Folder
}

// FetchFolderResponsePag represents the response that containing the list of folders with pagination.
type FetchFolderResponsePag struct {
	Folders   []*Folder
	NextToken string
}

// DataFetcher is an interface that provides a method to get sample data based on the purpose of the implementation (production/testing).
type DataFetcher interface {
	GetSampleData() ([]*Folder, error)
}

// RealDataFetcher is an implementation of the DataFetcher interface.
type RealDataFetcher struct{}
