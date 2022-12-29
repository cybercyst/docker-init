package generators

import (
	"github.com/noirbizarre/gonja"
)

var templateContent = `
# syntax=docker/dockerfile:1

ARG IMAGE=golang:1.19-alpine

FROM $IMAGE AS development

WORKDIR /app

# Copy project files
# NOTE: For more efficient builds, copy only
# required files here.
# For example, for a Go project, you may only
# need go.mod, go.sum and the actual *.go files
# that make up your project
COPY . .

RUN go run .

FROM $IMAGE AS builder

WORKDIR /app

# Copy project files
# NOTE: For more efficient builds, copy only
# required files here.
# For example, for a Go project, you may only
# need go.mod, go.sum and the actual *.go files
# that make up your project
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o {{ binary }}

FROM scratch AS production

COPY --from=builder /app/{{ binary }} .

{%- if port %}EXPOSE {{ port }}{% endif %}

ENTRYPOINT [ "/{{ binary }}" ]
`

func GenerateGo(context gonja.Context) (string, error) {
	tmpl, err := gonja.FromString(templateContent)
	if err != nil {
		return "", err
	}

	output, err := tmpl.Execute(context)
	if err != nil {
		return "", err
	}

	return output, nil
}
