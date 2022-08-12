FROM rockylinux:8

RUN yum update -y && \
    yum install -y bind-utils make git iproute python3-pip

RUN dnf install -y https://download.postgresql.org/pub/repos/yum/reporpms/EL-8-x86_64/pgdg-redhat-repo-latest.noarch.rpm && \
    dnf -qy module disable postgresql && \
    dnf install -y postgresql14-server

COPY scripts/* /scripts/
CMD /scripts/start.sh
