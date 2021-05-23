# docker build -t array-client -f Dockerfile.client .
# docker run -it -p 3000:3000 array-client 

FROM node:12.16.1-alpine3.9 as build
WORKDIR /app
ENV PATH /app/node_modules/.bin:$PATH
COPY /package.json /app
COPY /package-lock.json /app
COPY / /app
RUN npm install --silent
RUN npm run build


FROM nginx:1.17.8-alpine
COPY --from=build /app/build /usr/share/nginx/html
RUN rm /etc/nginx/conf.d/default.conf
COPY nginx/cert.pem /etc/ssl
COPY nginx/key.pem /etc/ssl
COPY nginx/nginx.conf /etc/nginx/conf.d
CMD ["nginx", "-g", "daemon off;"]