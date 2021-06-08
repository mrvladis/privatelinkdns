#!/bin/zsh
   # Make sure we're in the project directory within our GOPATH
    cd "hello-world"
      # Fetch all dependencies
    go get -t ./...
      # Ensure code passes all lint tests
    golint -set_exit_status
      # Check the Go code for common problems with 'go vet'
    go vet .
      # Run all tests included with our application
    go test .
    cd ..
      # Function 2
      # Make sure we're in the project directory within our GOPATH
    cd "hello-mr"
      # Fetch all dependencies
    go get -t ./...
      # Ensure code passes all lint tests
    golint -set_exit_status
      # Check the Go code for common problems with 'go vet'
    go vet .
      # Run all tests included with our application
    go test .
    cd ..
    # Make sure we're in the project directory within our GOPATH
    cd "books"
      # Fetch all dependencies
    go get -t ./...
      # Ensure code passes all lint tests
    golint -set_exit_status
      # Check the Go code for common problems with 'go vet'
    go vet .
      # Run all tests included with our application
    go test .
    cd ..
