# Create the container from the alpine linux image
FROM node:lts-alpine as build-stage

WORKDIR /app

# Copy our source code into the container
COPY . .
# Install the dependencies, can be commented out if you're running the same node version
RUN yarn install

# run webpack and the vue-loader
RUN yarn build

FROM nginx:stable-alpine as production-stage

# Create the directories we will need
RUN mkdir -p /tmp/nginx/simple-db-monitor
RUN mkdir -p /var/log/nginx
RUN mkdir -p /var/www/html

COPY --from=build-stage /app /tmp/nginx/simple-db-monitor

WORKDIR /tmp/nginx/simple-db-monitor

# Copy the respective nginx configuration files
COPY nginx.conf /etc/nginx/nginx.conf
COPY default.conf /etc/nginx/conf.d/default.conf

# copy the built app to our served directory
RUN cp -r dist/* /var/www/html

# make all files belong to the nginx user
RUN chown nginx:nginx /var/www/html

# start nginx and keep the process from backgrounding and the container from quitting
CMD ["nginx", "-g", "daemon off;"]
