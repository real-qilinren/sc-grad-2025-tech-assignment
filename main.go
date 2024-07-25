package main

import (
	"fmt"
	"github.com/georgechieng-sc/interns-2022/folders"
	"github.com/gofrs/uuid"
)

func main() {
	// Component 1
	//req := &folders.FetchFolderRequest{
	//	OrgID: uuid.FromStringOrNil(folders.DefaultOrgID),
	//}
	//
	//// Use the real data fetcher in production environment
	//res, err := folders.GetAllFolders(req, folders.RealDataFetcher{})
	//if err != nil {
	//	fmt.Printf("%v", err)
	//	return
	//}
	//
	//folders.PrettyPrint(res)

	// Component 2
	reqPag := &folders.FetchFolderRequestPag{
		OrgID:    uuid.FromStringOrNil(folders.DefaultOrgID),
		Token:    "",
		PageSize: 3,
	}

	for {
		resPag, err := folders.GetAllFoldersPag(reqPag, folders.RealDataFetcher{})
		if err != nil {
			fmt.Printf("%v", err)
			return
		}

		folders.PrettyPrint(resPag)

		// Update the token for the next iteration
		if resPag.NextToken == "" {
			break
		}

		reqPag.Token = resPag.NextToken
	}
}
