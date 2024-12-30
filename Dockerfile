# Etapa de construcción
ARG GO_VERSION=1
FROM golang:${GO_VERSION}-bookworm as builder

# Configura el directorio de trabajo
WORKDIR /usr/src/app

# Copia los archivos de configuración del proyecto
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Copia el resto del código fuente
COPY . .

# Compila el binario para linux/amd64
RUN GOOS=linux GOARCH=amd64 go build -v -o auth-service ./cmd/main.go

# Etapa final (imagen más ligera)
FROM debian:bookworm

# Establecer el directorio de trabajo
WORKDIR /usr/local/bin

# Copia el binario generado desde la etapa anterior
COPY --from=builder /usr/src/app/auth-service /usr/local/bin/auth-service

# Exponer el puerto utilizado por la aplicación
EXPOSE 8080

# Define el comando por defecto
CMD ["auth-service"]