FROM node:lts-alpine

WORKDIR /app

COPY package*.json /app

RUN npm install

#RUN npm install --save axios
#RUN npm install vue-router
#RUN npm install -g @vue/cli
#RUN vue add @vue/cli-plugin-eslint

COPY . .

#EXPOSE 8080

CMD [ "npm", "run", "serve" ]
