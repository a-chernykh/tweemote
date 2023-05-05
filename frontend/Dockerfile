FROM ubuntu:16.04

RUN apt-get update && \
      apt-get install -y curl git apt-transport-https && \
      curl -sL https://deb.nodesource.com/setup_6.x | bash - && \
      curl -sS https://dl.yarnpkg.com/debian/pubkey.gpg | apt-key add - && \
      echo "deb https://dl.yarnpkg.com/debian/ stable main" | tee /etc/apt/sources.list.d/yarn.list && \
      apt-get update && \
      apt-get install -y nodejs \
                         yarn=1.0.2-1 && \
      rm -rf /var/lib/apt/lists/*

ARG CONFIG

ENV APP_PATH=/app \
    CONFIG=${CONFIG:-docker}

RUN mkdir -p ${APP_PATH}
WORKDIR ${APP_PATH}

COPY package.json .
COPY yarn.lock .
RUN yarn install --non-interactive

COPY . ${APP_PATH}/

ENV NODE_ENV=production
RUN yarn bundle



FROM nginx:1.13.5

ENV APP_PATH=/app

COPY nginx.conf /etc/nginx/conf.d/default.conf
COPY --from=0 ${APP_PATH}/dist /usr/share/nginx/html
