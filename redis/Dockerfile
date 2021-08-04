FROM redis:5-alpine

ARG REDIS_PASSWORD
ENV REDIS_PASSWORD ${REDIS_PASSWORD}

COPY redis.conf /etc/redis.conf
RUN echo "requirepass ${REDIS_PASSWORD}" >> /etc/redis.conf

CMD ["redis-server", "/etc/redis.conf"]
