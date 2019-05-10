FROM nginx:1.15.5-alpine

ENV APP_PATH /app

RUN mkdir $APP_PATH && rm -rf /etc/nginx/conf.d/*
WORKDIR $APP_PATH

RUN mkdir /app/dist
COPY dist/ /app/dist/

COPY default.conf /etc/nginx/conf.d/

RUN rm -rf /usr/share/nginx/html/* \
  && mv ./dist/* /usr/share/nginx/html/ \
  && rm -rf ./*

CMD ["nginx", "-g", "daemon off;"]
