# syntax=docker/dockerfile:1

ARG IMAGE=node:18-alpine3.16

FROM $IMAGE AS development

WORKDIR /app

# Copy project files
# NOTE: For more efficient builds, copy only
# required files here.
# For example, for a Go project, you may only
# need go.mod, go.sum and the actual *.go files
# that make up your project
COPY . .

RUN npm ci

EXPOSE 3000

CMD ["npm", "start"]

FROM $IMAGE AS builder

WORKDIR /app

# Copy project files
# NOTE: For more efficient builds, copy only
# required files here.
# For example, for a Go project, you may only
# need go.mod, go.sum and the actual *.go files
# that make up your project
COPY . .

RUN npm ci && \
    npm run build

FROM nginx:1.23 AS production

EXPOSE 80

COPY --from=builder /app/build /usr/share/nginx/html/
