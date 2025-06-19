# Stage 1: Build Vue.js site
#FROM node:lts-alpine as builder-vuejs

#WORKDIR /app

#COPY frontend/package*.json ./
#RUN npm install
#COPY frontend/. .
#RUN npm run build

# Stage 2: Build Golang binary
FROM golang:1.24.4-alpine3.22 as builder-golang

WORKDIR /app

COPY backend/ ./
RUN go build

# Stage 3: Serve with Nginx
FROM nginx:stable-alpine

#COPY --from=builder-vuejs /app/dist /usr/share/nginx/html
COPY static /usr/share/nginx/html
COPY --from=builder-golang /app/word-to-abbreviation /usr/word-to-abbreviation
COPY docker/webapp/default.conf /etc/nginx/conf.d/default.conf

COPY docker/webapp/entrypoint.sh /docker-entrypoint.d/entrypoint.sh
RUN ["chmod", "+x", "/docker-entrypoint.d/entrypoint.sh"]

EXPOSE 80