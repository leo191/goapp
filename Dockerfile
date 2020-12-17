FROM golang:alpine AS build-env
RUN apk --no-cache add build-base git mercurial gcc
ADD pokemon /src/
RUN cd /src && go build -o gomon

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /src/gomon /app/
COPY --from=build-env /src/templates /app/templates 
ENTRYPOINT ./gomon