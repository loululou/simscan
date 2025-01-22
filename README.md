This is for practice.  

# Installation (Linux)
```
git clone https://github.com/loululou/simscan
cd simscan
go build -o simscan cmd/main.go
```

## Optional
```
GOOS=linux GOARCH=amd64 go build -o simscan-linux cmd/main.go
GOOS=linux GOARCH=arm64 go build -o simscan-linux cmd/main.go
GOOS=windows GOARCH=amd64 go build -o simscan.exe cmd/main.go
```