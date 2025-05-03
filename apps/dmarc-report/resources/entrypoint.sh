#!/usr/bin/env bash

PHP_ENV_FILE="/config/env.conf"

echo '[www]' > "$PHP_ENV_FILE"
echo 'user = nginx' >> "$PHP_ENV_FILE"
echo 'group = www-data' >> "$PHP_ENV_FILE"
echo 'listen.owner = nginx' >> "$PHP_ENV_FILE"
echo 'listen.group = www-data' >> "$PHP_ENV_FILE"
env | grep -e 'REPORT_DB_TYPE' -e 'REPORT_DB_HOST' -e 'REPORT_DB_PORT' -e 'REPORT_DB_NAME' -e 'REPORT_DB_USER' -e 'REPORT_DB_PASS' | sed "s/\(.*\)=\(.*\)/env[\1] = '\2'/" >> "$PHP_ENV_FILE"

# Get and parse dmarc reports once at startup to avoid PHP errors with a new database
if /usr/bin/dmarcts-report-parser.pl -i -d -r ; then
    echo 'INFO: Dmarc reports parsed successfully'
else
    echo 'CRIT: Dmarc reports could not be parsed. Check your IMAP and MYSQL Settings.'
    echo -e "DEBUG: Parsing failed with the following output:\n"
    cat /config/dmarc-reports.log
    exit 1
fi


# Start supervisord and services
/usr/bin/supervisord -n -c /etc/supervisord.conf
