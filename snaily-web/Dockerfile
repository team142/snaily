
# Build in a node container
FROM node:12-alpine as builder
RUN npm install -g @angular/cli
RUN mkdir /src
WORKDIR /src
COPY . .
RUN rm -rf /src/dist/*
RUN npm i
RUN ng build


# Release in an Nginx container
FROM nginx:1.16.0 as final
RUN rm -rf /etc/nginx/conf.d/*
RUN rm -rf /usr/share/nginx/html/*
COPY --from=builder /src/dist/snaily-web /usr/share/nginx/html/
COPY default.conf /etc/nginx/conf.d/
CMD ["nginx", "-g", "daemon off;"]
