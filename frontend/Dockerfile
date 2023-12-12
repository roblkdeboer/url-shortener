# Use a base image suitable for Node.js applications
FROM node:latest

# Set the working directory inside the container
WORKDIR /frontend

# Copy the React application source code
COPY . .

# Install dependencies and build the React app
RUN npm ci 
RUN npm run build

EXPOSE 3000

# Set the entry point for the container
CMD ["npm", "start"]