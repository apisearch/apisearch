FROM scratch
ENV DOCKER=true
ADD apisearch /
ENTRYPOINT ["/apisearch"]
EXPOSE 8080
