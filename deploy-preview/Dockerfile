FROM nginx:1.29.6-alpine@sha256:f46cb72c7df02710e693e863a983ac42f6a9579058a59a35f1ae36c9958e4ce0

COPY ./dist /usr/share/nginx/dist

RUN rm -rf /etc/nginx/sites-enabled && \
  rm -rf /etc/nginx/nginx.conf

COPY deploy-preview-files/deploy-preview/nginx.conf /etc/nginx/nginx.conf
COPY deploy-preview-files/deploy-preview/locations.conf /etc/nginx/locations.conf
COPY deploy-preview-files/deploy-preview/redirects.conf /etc/nginx/redirects.conf
COPY build.conf /etc/nginx/build.conf

RUN nginx -t
