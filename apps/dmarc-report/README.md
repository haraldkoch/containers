# sonarr

## Custom environment configuration

This container support setting certain custom enviroment variables with the use of [drone/envsubst](https://github.com/drone/envsubst).

| Name                        | Example                              |
| --------------------------- | ------------------------------------ |
| REPORT_DB_TYPE              | "pgsql"                              |
| REPORT_DB_HOST              | "postgres-rw.dbms.svc.cluster.local" |
| REPORT_DB_PORT              | "5432"                               |
| REPORT_DB_NAME              | "dmarc-report"                       |
| REPORT_DB_USER              | "dmarc"                              |
| REPORT_DB_PASS              | "password"                           |
| PARSER_IMAP_SERVER          | "imap.example.com"                   |
| PARSER_IMAP_PORT            | "143"                                |
| PARSER_IMAP_USER            | "<user@example.com>"                 |
| PARSER_IMAP_PASS            | "imap_password"                      |
| PARSER_IMAP_READ_FOLDER     | "Inbox"                              |
| PARSER_IMAP_MOVE_FOLDER     | "processed"                          |
| PARSER_IMAP_MOVE_FOLDER_ERR | "error"                              |
