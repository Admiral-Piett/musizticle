FROM golang:1.19.3 as base

WORKDIR "/opt/svc"

COPY app/ ./app
COPY main.go ./
COPY go.mod ./
COPY go.sum ./

RUN go build -o musizticle

# TODO - Figure out how to run off a smaller image
#FROM scratch
#COPY --from=base /opt/svc/musizticle ./

# TODO - add test build

ARG git_sha="local"
ENV GIT_SHA=${git_sha}

EXPOSE 9000

CMD ./musizticle > /var/log/musizticle.log 2>&1
