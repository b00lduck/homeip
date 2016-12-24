FROM scratch
ENTRYPOINT ["/homeip"]
ADD homeip /
USER 999:999
