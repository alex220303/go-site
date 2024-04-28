FROM node:20-alpine3.19

WORKDIR /app

COPY . .

RUN yarn

RUN yarn build

CMD ["yarn", "serve"]
