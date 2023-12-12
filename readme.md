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
3. Run `docker compose up -d` and visit [localhost:3000](localhost:3000) to interact with the frontend
4. If any changes are made to either the front or backend, you can run `docker compose up --build -d` to rebuild the images

Note: In the docker-compose.yml file, you will need to configure an absolute path for data persistence. For me, I set it to my documents on my windows PC.
