# URL shortener

### What it does
- It takes in yaml input of a shortened path and a url path
- The yaml input is converted to a map of the shortened path to url path
- When the server is run and a shortened path exists in the converted map, the server redirects to the url path
- If the shortened path does not exist, the server runs the mapHandler function to check if the url exists in the declared pathsToURLs
- If the shortened path does not exist in the declared pathsToURLs, it then returns a default response
