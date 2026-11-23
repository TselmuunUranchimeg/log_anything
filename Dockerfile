FROM golang:1.27.3-alpine AS build
WORKDIR /src

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /out/server ./cmd

FROM gcr.io/distroless/static-debian12:nonroot
COPY --from=build /out/server /server

EXPOSE 8000
USER nonroot:nonroot
ENTRYPOINT ["/server"]
