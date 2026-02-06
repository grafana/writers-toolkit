FROM nginx:1.29.4-alpine@sha256:4870c12cd2ca986de501a804b4f506ad3875a0b1874940ba0a2c7f763f1855b2

COPY ./dist /usr/share/nginx/dist

RUN rm -rf /etc/nginx/sites-enabled && \
  rm -rf /etc/nginx/nginx.conf

COPY deploy-preview-files/deploy-preview/nginx.conf /etc/nginx/nginx.conf
COPY build.conf /etc/nginx/build.conf

RUN nginx -t