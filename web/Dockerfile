FROM node:lts-alpine

WORKDIR /app

RUN npm install vue
#RUN npm install --save axios
#RUN npm install vue-router
RUN npm install -g @vue/cli
RUN vue add @vue/cli-plugin-eslint

COPY . .

EXPOSE 8080

CMD [ "npm", "run", "serve" ]
