https://search-blogs.aws.mckee-tech.com/?q=[任意の文字列]

使用技術
Go + Docker (マルチステージビルド)

インフラ
AWS lambda + CloudFront + ECR


ECRリポジトリにアップロードされたコンテナイメージがlambdaのFucntion URLsへのアクセスによって実行されます。
またCloudFrontによって独自ドメインからFucntion URLsによる関数の実行を可能にしています。
