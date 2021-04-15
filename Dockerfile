FROM ubuntu:16.04

ARG UNAME=user
ARG UID=1000
ARG GID=1000

RUN apt-get update -y && apt-get install -y curl gcc

RUN groupadd -g $GID -o $UNAME
RUN useradd -m -u $UID -g $GID -o -s /bin/bash $UNAME

USER $UNAME

RUN curl -sSL https://git.io/g-install | /bin/bash -s -- -y

ENV BASH_ENV="/home/${UNAME}/.bashrc"
ENV APP_HOME="/home/${UNAME}/app"

RUN mkdir -p "${APP_HOME}"

WORKDIR "${APP_HOME}"


CMD [ "/bin/bash", "-i", "-c", "go mod download && go build -buildmode c-shared -o \"${APP_HOME}/lib/dql/parse.so\" \"${APP_HOME}/src/dql-parse.go\"" ]
