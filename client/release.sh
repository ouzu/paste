
#!/bin/sh

env GOOS=linux GOARCH=amd64 go build -o ./release/linux-amd64/
env GOOS=linux GOARCH=arm GOARM=5 go build -o ./release/linux-arm5/
env GOOS=linux GOARCH=arm64 go build -o ./release/linux-arm64/
env GOOS=windows GOARCH=amd64 go build -o ./release/windows-amd64/
env GOOS=darwin GOARCH=amd64 go build -o ./release/darwin-amd64/
env GOOS=darwin GOARCH=arm64 go build -o ./release/darwin-arm64/
env GOOS=android GOARCH=arm64 go build -o ./release/android-arm64/
