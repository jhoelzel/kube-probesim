FROM gcr.io/distroless/static:nonroot
WORKDIR /
ADD ./bin/infrastructure-k8s-annotation-toggler /infrastructure-k8s-annotation-toggler
EXPOSE 8080

USER 65532:65532
ENTRYPOINT ["/infrastructure-k8s-annotation-toggler"]
