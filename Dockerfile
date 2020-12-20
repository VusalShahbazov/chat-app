# Start from golang base image
FROM golang:alpine

# Set the current working directory inside the container
WORKDIR /app

# Copy go mod and sum files
#COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed
#RUN go mod download

# Copy the source from the current directory to the working Directory inside the container
COPY ./app .

# Build the Go app
RUN go build  -o main .

RUN mkdir  front

COPY ./front   ../front

RUN apk add --update nodejs npm

RUN cd ../front && npm install -y && npm run build
# Expose port 8080 to the outside world
EXPOSE 3000

#Command to run the executable
CMD ["./main"]