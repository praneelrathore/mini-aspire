########################################
## Build Stage
########################################
FROM golang:1.22-bullseye

# setup the working directory
WORKDIR /go/src

# install dependencies
COPY ./go.sum ./go.sum
COPY ./go.mod ./go.mod

ENV GO111MODULE=on
ENV TZ=Asia/Kolkata

RUN go mod download

# add source code
COPY . .

# build the source
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o mini-aspire-linux-amd64

#COPY --from=builder /go/src/mini-aspire-linux-amd64 ./mini-aspire-linux-amd64
# add required files from host
#COPY ./configs/ ./configs/

# Run
ENTRYPOINT ["./mini-aspire-linux-amd64"]