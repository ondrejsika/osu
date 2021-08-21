FROM debian:10-slim
COPY osu /
ENTRYPOINT [ "/osu" ]
