FROM public.ecr.aws/lambda/provided:al2 as build
RUN yum update -y \
  && yum install -y \
  golang \
  git \
  bash \
  && yum clean all
RUN go env -w GOPROXY=direct
# ADD go.mod go.sum ./
ADD go.mod ./
RUN go mod download
ADD . .
RUN go build -o /main

FROM public.ecr.aws/lambda/provided:al2
COPY --from=build /main /main
ENTRYPOINT [ "/main" ]      