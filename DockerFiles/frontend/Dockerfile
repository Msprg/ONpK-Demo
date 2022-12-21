FROM node:12-alpine AS build

WORKDIR /app
RUN apk add --no-cache git
RUN git clone https://github.com/michaello1/zauni-zadanie-appfrontend.git
WORKDIR ./zauni-zadanie-appfrontend

RUN yarn install
RUN yarn run build

RUN cp env.sh /app/
RUN cp conf.d/default.conf /app/
RUN cp -r public /app/html

#Stage: 2
FROM nginx:stable-alpine
COPY --from=build /app/env.sh /dockerentrypoint.d
COPY --from=build /app/default.conf /etc/nginx/conf.d
COPY --from=build /app/html /usr/share/nginx/html
