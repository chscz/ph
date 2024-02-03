FROM golang:1.21-alpine AS gobuilder
ENV CGO_ENABLED 0
COPY . /app
WORKDIR /app
RUN go build -o ph .

FROM scratch
COPY --from=gobuilder /app/ph /
COPY ./templates /app/templates
CMD ["/ph"]
EXPOSE 8080