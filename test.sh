docker build -t hbutmc/web-backend . &&
docker-compose -f /srv/hbutmc-web/docker-compose.yml up -d &&
docker logs -f hbutmc-web-backend
