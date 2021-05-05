FROM nginx:1.15-alpine

RUN apk update
RUN apk add git

WORKDIR /usr/share/nginx/html/

# Clean the default public folder
RUN rm -fr * .??*

# Build website
WORKDIR /var/tmp
# Copy source
COPY ./source .
COPY ./merge.sh .
RUN ./merge.sh

WORKDIR /usr/share/nginx/html/

# This inserts a line in the default config file, including our file "expires.inc"
RUN sed -i '9i\        include /etc/nginx/conf.d/expires.inc;\n' /etc/nginx/conf.d/default.conf

# The file "expires.inc" is copied into the image
COPY ./expires.inc /etc/nginx/conf.d/expires.inc
RUN chmod 0644 /etc/nginx/conf.d/expires.inc
