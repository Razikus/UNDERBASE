FROM node:20 as build-stage
ARG AUTHORIZER
ARG BACKEND

WORKDIR /app

COPY package*.json ./

RUN yarn

COPY . .
RUN echo "AUTHORIZER=${AUTHORIZER}" > .env
RUN echo "BACKEND=${BACKEND}" >> .env
RUN yarn run build

FROM nginx:alpine

COPY --from=build-stage /app/dist/spa /usr/share/nginx/html

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]