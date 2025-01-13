FROM nginx:1.21.6-alpine
ARG ACTION_PATH
ENV ACTION_PATH="${ACTION_PATH:-deploy-preview-files/deploy-preview/}"

RUN rm -rf /etc/nginx/sites-enabled && \
  rm -rf /etc/nginx/nginx.conf

COPY "${ACTION_PATH}/deploy-preview-files/deploy-preview/nginx.conf" /etc/nginx/nginx.conf
COPY build.conf /etc/nginx/build.conf

COPY ./dist /usr/share/nginx/dist

RUN nginx -t
