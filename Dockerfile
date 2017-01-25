FROM scratch
ADD apisearch /
ENTRYPOINT ["/apisearch"]
EXPOSE 8080
