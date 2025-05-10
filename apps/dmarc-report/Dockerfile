ARG VERSION

FROM gutmensch/dmarc-report:${VERSION}

RUN ln -s /config/env.conf /etc/php81/php-fpm.d/env.conf

COPY ./resources/cron-root /etc/cron.d/root
COPY ./resources/entrypoint.sh /entrypoint.sh

VOLUME ["/config"]

CMD ["/entrypoint.sh"]
