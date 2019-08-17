FROM node:lts-alpine

WORKDIR /app

COPY . .

RUN npm install --save axios
RUN npm install vue-router

EXPOSE 8080

CMD [ "npm", "run", "serve" ]