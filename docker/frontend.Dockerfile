FROM node:24-alpine

WORKDIR /nft-music

COPY . .
COPY frontend/package.json frontend/yarn.lock frontend/.yarnrc.yml ./
COPY frontend/.yarn ./.yarn

RUN apk update && apk upgrade && \
    apk add python3 py3-pip cmake clang bash git alpine-sdk libressl-dev libc6-compat dumb-init ca-certificates

RUN npm install -g prettier @wagmi/cli
RUN npm install -g vscode-solidity-server

RUN corepack enable && yarn install --immutable
