FROM golang:1.18 as builder

COPY . /njson
WORKDIR /njson/mayhem

RUN go install github.com/dvyukov/go-fuzz/go-fuzz@latest github.com/dvyukov/go-fuzz/go-fuzz-build@latest
RUN go get github.com/dvyukov/go-fuzz/go-fuzz-dep
RUN apt update && apt install -y clang


RUN go-fuzz-build -libfuzzer -o fuzznjson.a && \
    clang -fsanitize=fuzzer fuzznjson.a -o fuzznjson.libfuzzer

FROM debian:bookworm-slim
COPY --from=builder /njson/mayhem/fuzznjson.libfuzzer /