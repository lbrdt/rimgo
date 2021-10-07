FROM node:16-slim as builder

RUN mkdir /app
COPY package.json /app/
COPY package-lock.json /app/
WORKDIR /app
RUN npm ci
COPY . /app
RUN npm run build

FROM node:16-slim

COPY --from=builder /app/dist/ /app/

WORKDIR /app

RUN npm install --production

CMD ["/usr/local/bin/node", "index.js"]
