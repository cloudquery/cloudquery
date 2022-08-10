FROM lambci/lambda:build-python3.8 as build

LABEL maintainer="Betajob AS" \
      description="Patched AWS Lambda build container"

COPY automake-1.13-to-1.16-spec.patch /root

RUN \
    cd /root \
 && yum install -y yum-utils spectool deltarpm \
 && yum-builddep -y automake

RUN \
    cd /root \
 && yumdownloader --source automake \
 && rpm -i automake-1.13.*.amzn2.src.rpm \
 && cd /root/rpmbuild \
 && patch -p0 < ../automake-1.13-to-1.16-spec.patch \
 && spectool -g -R SPECS/automake.spec \
 && rpmbuild -ba SPECS/automake.spec --nocheck \
 && yum install -y RPMS/noarch/*

FROM lambci/lambda:build-python3.8
COPY --from=build /root/rpmbuild/RPMS/noarch/*.rpm .
RUN yum install -y *.rpm \
 && rm *.rpm
