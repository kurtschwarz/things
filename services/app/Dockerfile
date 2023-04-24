FROM node:18.7.0-buster@sha256:5a72c3a0c03daba722426247af63565f5fb932b93a22733288c63d79bd29819e AS base

# ---

FROM base AS dependencies

WORKDIR /things/services/app

COPY ./services/app/package*.json /things/services/app/

RUN npm install

# ---

FROM base AS dev

WORKDIR /things/services/app

COPY --from=dependencies /things/services/app/node_modules/ /things/node_modules/
COPY --chown=node:node ./services/api/graph/ ./services/api/graph/
COPY --chown=node:node ./services/app/ /things/services/app/

ENTRYPOINT [ "npm" ]
CMD [ "run", "dev" ]
