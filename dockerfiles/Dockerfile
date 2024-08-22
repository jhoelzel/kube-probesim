FROM gcr.io/distroless/static:nonroot
WORKDIR /
ADD ./bin/kube-probesim /kube-probesim
EXPOSE 8080

USER 65532:65532
ENTRYPOINT ["/kube-probesim"]
