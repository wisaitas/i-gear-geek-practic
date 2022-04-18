FROM node:17-alpine
WORKDIR /nuxt-web-app
RUN apk update && apk upgrade
RUN apk add git

COPY . /nuxt-web-app/
RUN npm install