#
# Multi-stage build for Angular app
#

# Use node image with ng cli
FROM node:12-alpine as ngbuilder
RUN npm install -g @angular/cli

# Install npm dependencies
FROM ngbuilder as ngbuilder2
RUN mkdir /src/
RUN mkdir /src/snaily-web
COPY package.json /src/snaily-web
WORKDIR /src/snaily-web
RUN npm i


FROM ngbuilder2 as ngbuilder3
COPY . /src/snaily-web
WORKDIR /src/snaily-web
RUN ng build

# Server
FROM nginx:alpine as ngserver
COPY nginx.conf /etc/nginx/nginx.conf

FROM ngserver as ngserver2
WORKDIR /usr/share/nginx/html
COPY --from=ngbuilder3 /src/snaily-web/dist/snaily-web .

