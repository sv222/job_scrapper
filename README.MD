# Job scrapper service

1. Scrapping fresh jobs from profesia.sk
2. Parser works concurrently and adds new raws to list every 8 hours .
3. Filtering only job positions with Go in title.

## Endpoints:
list of data: /

check service status: /health-check

### Docker

For build docker image look at Dockerfile.

Default apps port is 7070