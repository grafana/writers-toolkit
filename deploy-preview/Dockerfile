FROM nginx:1.21.6-alpine@sha256:a74534e76ee1121d418fa7394ca930eb67440deda413848bc67c68138535b989

COPY ./dist /usr/share/nginx/dist

RUN rm -rf /etc/nginx/sites-enabled && \
  rm -rf /etc/nginx/nginx.conf

COPY deploy-preview-files/deploy-preview/nginx.conf /etc/nginx/nginx.conf
COPY build.conf /etc/nginx/build.conf

RUN nginx -t