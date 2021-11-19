FROM golang:1.17

WORKDIR /go/src
ENV PATH="/go/bin/${PATH}"

RUN go install github.com/spf13/cobra/cobra@latest && \
    go install github.com/golang/mock/mockgen@v1.5.0

CMD ["tail", "-f", "/dev/null"]