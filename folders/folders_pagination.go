package folders

import (
	"encoding/base64"
	"github.com/gofrs/uuid"
	"strconv"
)

// EncodeToken encodes the page number into a token.
func EncodeToken(pageNum int) string {
	token := strconv.Itoa(pageNum)
	return base64.StdEncoding.EncodeToString([]byte(token))
}

// DecodeToken decodes the token into a page number.
func DecodeToken(token string) (int, error) {
	if token == "" {
		return 0, nil
	}

	decoded, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return 0, err
	}

	return strconv.Atoi(string(decoded))
}

func GetAllFoldersPag(req *FetchFolderRequestPag, fetcher DataFetcher) (*FetchFolderResponsePag, error) {
	// Decode the token to get the page number
	pageNum, err := DecodeToken(req.Token)

	// Fetch all folders by organization ID
	folders, token, err := FetchAllFoldersByOrgIDPag(req.OrgID, fetcher, pageNum, req.PageSize)
	if err != nil {
		return nil, err
	}

	// Construct the FetchFolderResponsePag
	res := &FetchFolderResponsePag{
		Folders:   folders,
		NextToken: token,
	}

	return res, nil
}

func FetchAllFoldersByOrgIDPag(orgID uuid.UUID, fetcher DataFetcher, pageNum int, pageSize int) ([]*Folder, string, error) {
	// Get sample data from the JSON file
	folders, err := fetcher.GetSampleData()
	if err != nil {
		return nil, "", err
	}

	// Filter the folders by organization ID
	resFolder := []*Folder{}
	for _, folder := range folders {
		if folder.OrgId == orgID {
			resFolder = append(resFolder, folder)
		}
	}

	// Pagination Process
	// Step1: Calculate the page number and page size
	startIdx := pageNum * pageSize
	endIdx := startIdx + pageSize

	if startIdx >= len(resFolder) {
		return nil, "", nil
	}

	if endIdx > len(resFolder) {
		endIdx = len(resFolder)
	}

	// Step2: Get the folders for the current page
	currentPage := resFolder[startIdx:endIdx]

	// Step3: Generate the token for the next page
	nextToken := ""
	if endIdx < len(resFolder) {
		nextToken = EncodeToken(pageNum + 1)
	}

	return currentPage, nextToken, nil
}
