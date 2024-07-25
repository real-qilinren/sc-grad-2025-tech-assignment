# Explanation

To implement the Pagination (Component 2) for our code, I did the following modifications: 

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