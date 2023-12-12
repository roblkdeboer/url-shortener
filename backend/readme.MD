# Background

This is the backend to a URL shortener. It will:

- Accept a URL posted to it and encode it using base52
- Store the encoded URL and it's corresponding URL in mongodb
- Return the encoded URL to be used on the frontend
- Retrieve the full URL and redirec the user when they try to visit the shortened URL
- Give the user the ability to delete a stored encoded URL

# How to run

- Clone this repoistory
- Ensure atleast go version 1.20.6 is installed
- Run go mod download to get any required packages
- Ensure all third party modules (like mongodb are installed)
- Run `go run main.go` in the root folder
