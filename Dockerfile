FROM scratch

ENTRYPOINT ["homeip"]

ENV PORT=8080

ADD homeip homeip

USER 999:999
