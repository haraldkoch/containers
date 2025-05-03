ARG VERSION

FROM gutmensch/dmarc-report:${VERSION}


RUN ln -s /config/env.conf /etc/php81/php-fpm.d/env.conf

COPY ./apps/dmarc-report/cron-root /etc/cron.d/root
COPY ./apps/dmarc-report/entrypoint.sh /entrypoint.sh
CMD ["/entrypoint.sh"]

LABEL org.opencontainers.image.source="https://github.com/gutmensch/docker-dmarc-report"
