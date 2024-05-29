cd $(dirname $0) && pwd
# build
#GOOS=linux go build -trimpath -o deploy plugin/drone/deploy
#docker run --rm -v $GOPATH:/go -v $PWD:/work -w /work/plugin -e GOPROXY=$GOPROXY $GOIMAGE go build -o /work/build/deploy
docker build -t jybl/deploy .
docker push jybl/deploy


# test notify
#docker run --rm -e PLUGIN_DING_TOKEN=xxx -e PLUGIN_DING_SECRET=xxx jybl/deploy
