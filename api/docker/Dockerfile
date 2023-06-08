# Entorno: Desarrollo

# Utiliza una imagen base adecuada para tu aplicación en el entorno de desarrollo
FROM golang:alpine

# Instalar git
RUN apk update && apk add --no-cache git

# Establece el directorio de trabajo
WORKDIR /app

# Copia los archivos necesarios de tu aplicación al contenedor
COPY ./ /app

COPY go.mod go.sum ./

# Instala las dependencias de tu aplicación
RUN go mod download

# Instala Air para hot reload
RUN go get -u github.com/cosmtrek/air

CMD ["go", "run", "main.go"]