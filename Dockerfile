FROM golang:1.22.5 as build
WORKDIR /app
COPY . .
RUN go build -o /cmd/web/main .

FROM scratch
COPY --from=build /server /server
EXPOSE 3000
CMD ["/server"]