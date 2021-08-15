FROM golang:1.16-alpine

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod .
RUN go mod download

COPY *.go ./
COPY /nifs /app/nifs

# Build
RUN go build -o /nifstable

# Run
CMD [ "/nifstable" ]