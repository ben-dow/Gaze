FROM golang:1.21 as backend
WORKDIR /src
COPY go.mod /src
COPY go.sum /src
RUN go mod download
ADD . /src
RUN CGO_ENABLED=0 go build -o /src/.build/backend /src/cmd/gaze/

FROM node:18 as frontend
WORKDIR /src
ADD ./web /src/web
WORKDIR /src/web/gaze
RUN npm install
RUN npm run build
RUN echo $(ls /src/web/gaze)

FROM scratch
COPY --from=backend /src/.build/backend /bin/backend
COPY --from=frontend /src/web/gaze/build /opt/web
CMD ["/bin/backend"]
EXPOSE 3000