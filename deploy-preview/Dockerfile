FROM nginx:1.21.6-alpine

COPY ./dist /usr/share/nginx/dist

RUN rm -rf /etc/nginx/sites-enabled && \
  rm -rf /etc/nginx/nginx.conf

COPY ./deploy-preview/nginx.conf /etc/nginx/nginx.conf

RUN nginx -t