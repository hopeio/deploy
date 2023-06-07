GOOS=linux go build -trimpath -o out main.go

docker run --rm  --privileged=true -u root -v $PWD:/work -w /work node:16-alpine3.16 npm run build