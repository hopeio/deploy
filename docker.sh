cd $(dirname $0) && pwd
# build
docker run --rm -v $GOPATH:/go -v $PWD:/work -w /work/plugin -e GOPROXY=$GOPROXY $GOIMAGE go build -o /work/deploy/deploy
docker build -t jybl/deploy .
docker push jybl/deploy


# test notify
#docker run --rm -e PLUGIN_DING_TOKEN=xxx -e PLUGIN_DING_SECRET=xxx jybl/deploy
