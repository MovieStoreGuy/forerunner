# This is an example dockerfile to use with forerunner
# This also allows for me to ensure what I am passing to forerunner
# is actually reflected inside the running docker container
FROM busybox:latest

ENTRYPOINT ["/bin/sh"]
CMD ["env","&&","cat","/etc/hosts"]
