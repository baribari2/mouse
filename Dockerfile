FROM 1.18.6-bullseye

WORKDIR /

COPY . .

RUN go mod download && go mod verify 

RUN go run cmd/mouse/main.go