FROM golang:1.9.2-alpine3.6 AS BUILD

WORKDIR /legion

#TODO: need to build the go app as
# Now just add the binary
ADD mylegion /legion/
ENTRYPOINT ["./mylegion"]
