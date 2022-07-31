FROM golang:1.18-bullseye

ARG shared_workspace=/opt/workspace

RUN mkdir -p ${shared_workspace} 

ENV SHARED_WORKSPACE=${shared_workspace}

WORKDIR ${SHARED_WORKSPACE}

VOLUME ${SHARED_WORKSPACE}
CMD ["bash"]