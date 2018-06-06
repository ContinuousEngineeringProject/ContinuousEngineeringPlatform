FROM golang:1.10 AS BUILD

WORKDIR /cePlatform

#TODO: need to build the go app
# Now just add the binary
ADD myOminus /Ominus/
ENTRYPOINT ["./myOminus"]




# WORKDIR /go/src/app
# COPY . .

# RUN go get -d -v ./...
# RUN go install -v ./...

# CMD ["app"]