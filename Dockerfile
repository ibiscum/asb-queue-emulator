FROM mcr.microsoft.com/cbl-mariner/base/core:2.0 as base
RUN tdnf -y update
RUN tdnf -y install ca-certificates

FROM scratch
COPY --from=base /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

WORKDIR /ASBEMULATOR

# Copy the source code
COPY asbemulator /ASBEMULATOR
COPY /config/default_config.json /ASBEMULATOR/config/
EXPOSE 4444

ENTRYPOINT [ "./asbemulator" ]