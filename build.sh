

rm -f go.sum
go get -u github.com/Patrick-ring-motive/utils
go get -u github.com/Patrick-ring-motive/traigo
go build -ldflags "-g" -gcflags="-B -v -std"  -o ione ione.go