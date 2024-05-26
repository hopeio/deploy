# build
#GOOS=linux go build -trimpath -o notify plugin/drone/dingding_notify
#docker run --rm -v $GOPATH:/go -v $Code:/work -w /work/plugin/drone -e GOPROXY=$GOPROXY $GOIMAGE go build -o /work/plugin/drone/dingding_notify/notify /work/plugin/drone/dingding_notify
docker build -t jybl/notify plugin/drone/dingding_notify
docker push jybl/notify


# test notify
docker run --rm -e PLUGIN_DING_TOKEN=xxx -e PLUGIN_DING_SECRET=xxx jybl/notify
