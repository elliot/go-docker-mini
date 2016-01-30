FROM centurylink/ca-certs

WORKDIR /app

COPY serve /app/

ENV HOME /app
ENV PORT 8080

EXPOSE 8080

CMD ["./serve"]
