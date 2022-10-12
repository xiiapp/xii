ARG NGINX_VERSION
FROM ${NGINX_VERSION}

ARG TZ
ARG NGINX_VERSION
ARG CONTAINER_PACKAGE_URL
ARG NGINX_INSTALL_APPS


RUN if [ "${CONTAINER_PACKAGE_URL}" != "" ]; then \
        sed -i "s/dl-cdn.alpinelinux.org/${CONTAINER_PACKAGE_URL}/g" /etc/apk/repositories; \
    fi

# Install acme.sh
RUN apk add openssl && curl https://get.acme.sh | sh -s email=xii@exmaple.com && alias acme.sh=~/.acme.sh/acme.sh


WORKDIR /www
