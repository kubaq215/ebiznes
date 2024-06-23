# Use the official lightweight Node.js 14 image.
# https://hub.docker.com/_/node
FROM node:14-slim AS build

# Create and change to the app directory.
WORKDIR /frontend

# Copy application dependency manifests to the container image.
# A wildcard is used to ensure both package.json AND package-lock.json are copied.
COPY package*.json ./

# Install production dependencies.
RUN npm install

# Copy local code to the container image.
COPY . ./

# Build the app for production.
RUN npm run build

# Use the official nginx image for a production build.
# https://hub.docker.com/_/nginx
FROM nginx:stable

# Copy the build output to replace the default nginx contents.
COPY --from=build /frontend/build /usr/share/nginx/html

# Document that the service listens on port 80.
EXPOSE 80

# Run nginx with global directives and daemon off
ENTRYPOINT ["nginx", "-g", "daemon off;"]
