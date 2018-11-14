# build
FROM node:alpine as builder
ARG VUE_APP_API
ENV VUE_APP_API ${VUE_APP_API}
RUN mkdir /app
WORKDIR /app
COPY package.json /app/package.json  
RUN npm install
COPY . /app
RUN npm run build

# deployment
FROM nginx:latest
COPY --from=builder /app/dist /usr/share/nginx/html
# COPY heroku-nginx.conf /app/nginx.conf
# CMD /bin/bash -c "envsubst '\$PORT' < /app/nginx.conf > /etc/nginx/conf.d/default.conf" && nginx -g 'daemon off;'
CMD ["nginx", "-g", "daemon off;"]