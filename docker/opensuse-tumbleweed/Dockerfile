FROM opensuse:tumbleweed
MAINTAINER Pacur <contact@pacur.org>

RUN zypper -y up
RUN zypper -y in -t pattern devel_basis devel_C_C++ devel_rpm_build
RUN zypper -y in tar expect gpg2 openssl createrepo rsync make automake gcc wget zip git bzr mercurial

RUN wget https://storage.googleapis.com/golang/go1.9.2.linux-amd64.tar.gz
RUN tar -C /usr/local -xf go1.9.2.linux-amd64.tar.gz
RUN rm -f go1.9.2.linux-amd64.tar.gz

ENV GOPATH /go
ENV PATH /usr/local/go/bin:$PATH:/go/bin

RUN go get github.com/pacur/pacur

ENTRYPOINT ["pacur"]
CMD ["build", "opensuse-tumbleweed"]
