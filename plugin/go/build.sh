
if [ -z "$1" ]; then
  echo "编译目录参数为空"
  exit 1
else
  output=$(basename "$1")
  echo "output: $output"
fi


cd $(dirname $0) && pwd

# build
GOOS=linux go build -trimpath -o build/$output ./$1
#docker run --rm -v $GOPATH:/go -v $PWD:/work -w /work -e GOPROXY=$GOPROXY $GOIMAGE go build  -trimpath -o /work/build/$output /work/$1
docker build -t jybl/$output -f $1/Dockerfile ./build
docker push jybl/$output
