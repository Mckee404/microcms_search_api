FROM golang:1.24.2 AS build

# 作業ディレクトリを設定
WORKDIR /workspace/go

# 依存関係のインストールとビルド
ARG ARCH="arm64"

# --mount=type=bind でホストのgoディレクトリをコンテナ内にマウント
RUN --mount=type=bind,source=./go,target=/workspace/go go mod download

# Go バイナリのビルド
RUN --mount=type=bind,source=./go,target=/workspace/go GOARCH=${ARCH} CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o /functions/hello ./main.go

# ローカル実行用
FROM public.ecr.aws/lambda/provided:al2 AS local
COPY --from=build /functions /functions
ENTRYPOINT ["/usr/local/bin/aws-lambda-rie"]

# 本番環境用
FROM gcr.io/distroless/static-debian11 AS production
COPY --from=build /functions /functions
ENTRYPOINT ["/functions/hello"]