FROM nginx:1.21.6-alpine

COPY ./dist /usr/share/nginx/dist

RUN rm -rf /etc/nginx/sites-enabled && \
  rm -rf /etc/nginx/nginx.conf

COPY deploy-preview-files/deploy-preview/nginx.conf /etc/nginx/nginx.conf
COPY build.conf /etc/nginx/build.conf

RUN nginx -t