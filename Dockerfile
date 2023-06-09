# Build web
FROM golang:1.17-alpine3.15 as web
RUN mkdir /build
WORKDIR /build
ADD . ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -o web .


FROM scratch
WORKDIR /app
COPY --from=web /build/web /app/
#COPY --from=web /build/.env /app/.env
COPY --from=web /build/assets /app/assets
COPY --from=web /build/views /app/views
ENTRYPOINT ["./web"]
