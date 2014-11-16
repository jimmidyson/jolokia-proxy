FROM progrium/busybox
MAINTAINER jimmidyson@gmail.com

# Grab jolokia-proxyfrom the staging directory.
ADD jolokia-proxy /usr/bin/jolokia-proxy

EXPOSE 8080
CMD ["/usr/bin/jolokia-proxy"]
