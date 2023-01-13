FROM golang:1.19 as build
WORKDIR /go/src/github.com/drewart/hen
COPY . .
RUN go install -v ./...

FROM debian:bookworm-slim
EXPOSE 3000 
ENV SERVICE=hen
COPY --from=build /go/bin/hen /bin/hen
ENTRYPOINT ["/bin/hen", "server"]
