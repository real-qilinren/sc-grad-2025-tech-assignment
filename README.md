# sc-grad-2025

The technical assessment for SafetyCulture graduate program of 2025.

**Please note the explanation for component 2 is at the end of the README.md!**

## Getting started

Requires `Go` >= `1.20`

follow the official install instruction: [Golang Installation](https://go.dev/doc/install)

To run the code on your local machine
```
  go run main.go
```

## Folder structure

```
| go.mod
| README.md
| sample.json
| main.go
| folders
    | folders.go
    | folders_test.go
    | static.go
```

## Instructions

- This technical assessment consists of 2 components:
- Component 1:
  - within `folders.go`.
    - We would like you to read through, and run, the code.
    - Write some comments on what you think the code does.
    - Suggest some improvements that can be made to the code.
    - Implement any suggested improvements.
    - Write up some unit tests in `folders_test.go` for your new `GetAllFolders` method

- Component 2:
  - Extend your improved code to now facilitate pagination.
  - You can copy over the existing methods into `folders_pagination.go` to get started.
  - Write a short explanation of your chosen solution.

## What is pagination?
  - Pagination helps break down a large dataset into smaller chunks.
  - Those smaller chunks can then be served to the client, and are usually accompanied by a token pointing to the next chunk.
  - The end result could potentially look like this:
```
  original data: [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]

  This might result in the following payload to the client:
  { data: [1, 2, 3, ..., 10] }

  However, with pagination implemented, the payload might instead look like this:
  request() -> { data: [1, 2], token: "nQsjz" }

  The token could then be used to fetch more results:

  request("nQsjz") -> { data : [3, 4], token: "uJsnQ" }

  .
  .
  .

  And more results until there's no data left:

  { data: [9, 10], token: "" }
```

## Explanation for Component 2

To implement the Pagination for our code, I did the following modifications:

1. **Changed the request and response structures**
  * Ensure the resquest contains the token (represents the page number) so that our code knows the start index (where to begin scanning the folders) and also the page size so that our code knows how many folders should one single page contains.
  * Ensure the response has the folder list and the token for the next scan so that our code knows what page number will be used in the next iteration.
2. **Token Encryption/Decryption**
  * Simply encode the page numbers in integer format to base64 string to create unique tokens.
  * And decode the tokens to get the current page number.
3. **Data Extraction for the current page**
  * Caculate the start and end indices based on the page number (token) and the page size to know which subset of folders will be extracted.
4. **Token Generation**
  * Then generate the next token if there are more folders (pages) to be retrieved, the user uses the next token to request the subsequent page, continuing until no more data is available.