############################
# STEP 1 build executable binary
############################
FROM golang:alpine as builder

# Install git + SSL ca certificates.
# Git is required for fetching the dependencies.
# Ca-certificates is required to call HTTPS endpoints.
RUN apk update && apk add --no-cache git ca-certificates tzdata && update-ca-certificates

# Create appuser
RUN adduser -D -g '' appuser
RUN mkdir /app
COPY . /app
WORKDIR /app

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod=vendor -ldflags="-w -s" -o /server
############################
#STEP 2 build a small image
############################
FROM scratch
ENV LISTEN_PORT=3000
# Import from builder.
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
ENV TZ=Europe/Moscow

# Copy our static executable
COPY --from=builder /server /server

# Use an unprivileged user.
USER appuser

# Port on which the service will be exposed.
EXPOSE ${LISTEN_PORT}

# Run the hello binary.
ENTRYPOINT ["/server"]