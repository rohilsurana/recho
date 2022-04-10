FROM alpine:3.13

COPY recho /usr/bin/recho

RUN apk --no-cache add ca-certificates bash

CMD ["recho"]
