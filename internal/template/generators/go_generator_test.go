package generators

import (
	"testing"

	"github.com/noirbizarre/gonja"
)

func TestGenerateGo(t *testing.T) {
	got, err := GenerateGo(gonja.Context{"binary": "test-binary", "port": 8080})
	if err != nil {
		t.Fatalf("generating go template should not error")
	}

	want := `
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

RUN CGO_ENABLED=0 GOOS=linux go build -o test-binary

FROM scratch AS production

COPY --from=builder /app/test-binary .

EXPOSE 8080

ENTRYPOINT [ "/test-binary" ]`

	if got != want {
		t.Fatalf("generated go dockerfile did not match expectation")
	}
}
