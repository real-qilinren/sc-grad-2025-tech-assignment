package folders

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"os"
	"path/filepath"
	"runtime"

	"github.com/gofrs/uuid"
	"github.com/lucasepe/codename"
)

// These are all helper methods and fixed types.
// There's no real need for you to be editting these, but feel free to tweak it to suit your needs.
// If you do make changes here, be ready to discuss why these changes were made.

const dataSetSize = 1000
const DefaultOrgID = "c1556e17-b7c0-45a3-a6ae-9546248fb17a"

// The Folder struct represents a folder with its attributes.
type Folder struct {
	// An unique identifier for the folder, must be a valid UUID.
	// For example: '00001d65-d336-485a-8331-7b53f37e8f51'
	Id uuid.UUID `json:"id"`
	// Name associated with folder.
	Name string `json:"name"`
	// The organisation that the folder belongs to.
	OrgId uuid.UUID `json:"org_id"`
	// Whether a folder has been marked as deleted or not.
	Deleted bool `json:"deleted"`
}

// GenerateData generates a list of sample folders
// It returns the pointers to the Folder struct.
func GenerateData() []*Folder {
	rng, _ := codename.DefaultRNG()
	sampleData := []*Folder{} // Create an empty slice of Folder pointers

	for i := 1; i < dataSetSize; i++ { // Generate folders
		orgId := uuid.FromStringOrNil(DefaultOrgID)

		if i%3 == 0 { // Randomly change the orgId
			orgId = uuid.Must(uuid.NewV4())
		}

		deleted := rand.Int() % 2 // Randomly set the deleted flag

		sampleData = append(sampleData, &Folder{
			Id:      uuid.Must(uuid.NewV4()),
			Name:    codename.Generate(rng, 0),
			OrgId:   orgId,
			Deleted: deleted != 0,
		})
	}

	return sampleData
}

// PrettyPrint prints the data in a human-readable JSON format.
// It takes any data type and prints the JSON string.
func PrettyPrint(b interface{}) {
	s, _ := json.MarshalIndent(b, "", "\t")
	fmt.Print(string(s))
}

// GetSampleData reads the sample data from a JSON file
// RealDataFetcher is an implementation of the DataFetcher interface.
// It returns the pointers to the Folder struct.
// THIS METHOD HAS BEEN MODIFIED - simply adding error handling
func (RealDataFetcher) GetSampleData() ([]*Folder, error) {
	_, filename, _, _ := runtime.Caller(0) // Get the current file path
	fmt.Println(filename)
	basePath := filepath.Dir(filename)
	filePath := filepath.Join(basePath, "sample.json") // Get the sample data file path

	fmt.Println(filePath)

	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	jsonByte, _ := io.ReadAll(file) // Read the file content

	//folders := []*Folder{}
	var folders []*Folder                    // Change to var to ensure clarity and no there is no memory allocation
	err = json.Unmarshal(jsonByte, &folders) // Unmarshal the JSON content to the folders slice
	if err != nil {
		return nil, err
	}

	return folders, nil
}
