################################################################################
FROM --platform=${BUILDPLATFORM} golang:1.25.2-alpine@sha256:2389ebfa5b7f43eeafbd6be0c3700cc46690ef842ad962f6c5bd6be49ed82039 AS builder
ARG TARGETOS TARGETARCH

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS="${TARGETOS}" GOARCH="${TARGETARCH}" go build -o /out/broken-links ./cmd/link-checker

################################################################################
FROM nginx:1.29.6-alpine@sha256:f46cb72c7df02710e693e863a983ac42f6a9579058a59a35f1ae36c9958e4ce0

# Git supports comment generation in environments that infer changed files from repo history.
RUN apk add --no-cache ca-certificates git

COPY --from=builder /out/broken-links /usr/local/bin/broken-links

ENTRYPOINT ["broken-links"]
CMD ["help"]
