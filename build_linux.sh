rm -rf build/
mkdir -p build/
cp -r config build/
GOOS=linux GOARCH=amd64 go build -o build/main-cicd
