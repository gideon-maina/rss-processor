# Use the offical golang image
FROM golang

# Create directory for the app
RUN mkdir /app

# Copy the application files
ADD . /app

# Set the working directory
WORKDIR /app
