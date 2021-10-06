FROM node:16-slim as builder

COPY . /app
WORKDIR /app
RUN npm ci && npm run build

FROM node:16-slim

COPY --from=builder /app/dist/ /app/
COPY --from=builder /app/package.json /app/
COPY --from=builder /app/package-lock.json /app/
COPY --from=builder /app/static/ /app/static/
COPY --from=builder /app/templates/ /app/templates/

WORKDIR /app

RUN npm install --production

CMD ["/usr/local/bin/node", "index.js"]
