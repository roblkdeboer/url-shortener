# Background

This is a fullstack URL shortener app:

- React frontend
- Golang backend
- MongoDB database

# Functionality

You can:

- Shorten a URL
- Expand a shortened URL and get redirected to it
- Delete an existing shortened URL from the database

# How to run locally

1. Clone this repository
2. Ensure Docker is installed on your machine
3. Run `docker compose up -d` and visit localhost:3000 to interact with the frontend

Note: In the Dockerfile for the backend, you will need to configure an absolute path for data persistence. For me, I set it to my documents on my windows PC.
