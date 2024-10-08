# todo-app

### ディレクトリ構造
```
├── api/                             バックエンド用のディレクトリ（Go関連ファイル）
│   ├── domain/
│   ├── infrastructure/
│   ├── interface/
│   ├── migrations/
│   ├── main.go
│   ├── go.mod
│   ├── go.sum
│   └── readme.md
├── front/                           フロントエンド用ディレクトリ（ReactやVueなど）
│   ├── src/                         フロントエンドのコード
│   │   ├── components/              コンポーネント
│   │   ├── pages/                   ページコンポーネント
│   │   ├── services/                API呼び出しや外部通信のロジック
│   │   ├── App.js                   アプリケーションのエントリポイント
│   │   └── index.js                 メインエントリポイント
│   ├── public/                      静的ファイル
│   ├── package.json                 npmパッケージの管理
│   ├── package-lock.json
│   └── webpack.config.js            WebpackやViteなどの設定ファイル
├── .env                             環境変数（共通で使う場合）
└── readme.md                        プロジェクト全体の説明
```

#### バックエンド
```
├── domain                           ドメイン層
│   ├── application                  Application Business Rules
│   │   ├── repository               リポジトリのインターフェース
│   │   │   ├── user.go
│   │   │   └── todo.go
│   │   └── usecase                  Usecase
│   │       ├── user_input_port.go   Input PortのインターフェースとInput Dataの構造定義。ドメインが望む入力形式の定義
│   │       ├── user_interactor.go   Interactor。Input DataとRepositoryを使ってデータを取得、Output Portの実装であるPresenterに渡す
│   │       ├── user_output_port.go  Output PortのインターフェースとOutput Dataの構造定義。ドメインが望む出力形式の定義
│   │       ├── todo_input_port.go   Input PortのインターフェースとInput Dataの構造定義。ドメインが望む入力形式の定義
│   │       ├── todo_interactor.go   Interactor。Input DataとRepositoryを使ってデータを取得、Output Portの実装であるPresenterに渡す
│   │       └── todo_output_port.go  Output PortのインターフェースとOutput Dataの構造定義。。ドメインが望む出力形式の定義
│   └── entity                       Enterprise Business Rules
│       ├── user.go                  ユーザーのデータ構造。DDDにおける値オブジェクトやドメインサービスまで定義していいと思う。
│       ├── admin.go                 管理者のデータ構造。DDDにおける値オブジェクトやドメインサービスまで定義していいと思う。
│       └── todo.go                  データ構造。DDDにおける値オブジェクトやドメインサービスまで定義していいと思う。
├── go.mod
├── go.sum
├── infrastructure                   Frameworks & Drivers
│   ├── middleware                   Middlewareの実装。認証や認可のミドルウェア
│   │   └── auth.go                  JWT認証の実装
│   └── router                       Ginの実装。Controllerと繋げる
│       ├── api
│       │   ├── user.go
│       │   └── todo.go
│       └── router.go
├── interface                        Interface Adapters
│   ├── controller                   Controller。データをInteractorが望むInputDataに整形して渡し、Presenterが整形したデータを表示する
│   │   ├── context.go
│   │   ├── error.go
│   │   ├── user.go
│   │   └── todo.go
│   ├── presenter                    Presenter。Output Portの実装。OutDataを受け取ってレスポンスを整形する
│   │   └── user.go
│   │   └── todo.go
│   └── repository
│       ├── user.go                  リポジトリのインターフェース
│       └── todo.go                  リポジトリの実装。Databaseを使ったデータの取得など
├── main.go
├── .env                             auth0の環境変数
└── migrations
    ├── 000001_create_todos.down.sql
    └── 000001_create_todos.up.sql
```

# front
```
src
|
+-- app
|   |
|   +-- routes
|   |   |
|   |   +-- index.tsx        # ホームページ（タイトルとログイン/サインアップボタン）
|   |   +-- todos.tsx        # Todoリストのページ
|   |
|   +-- app.tsx              # メインアプリケーションコンポーネント
|   +-- provider.tsx         # グローバルプロバイダ（Redux、Auth0など）
|   +-- router.tsx           # ルーティング設定
|
+-- assets                   # 画像やフォントなどの静的ファイル
|
+-- components               # 共通コンポーネント（ボタン、ヘッダーなど）
|
+-- config
|   |
|   +-- auth0.ts             # Auth0の設定ファイル
|
+-- features
|   |
|   +-- auth
|   |   +-- authSlice.ts     # 認証用のReduxスライス
|   |
|   +-- todos
|       +-- todosSlice.ts    # Todo用のReduxスライス
|
+-- hooks
|   |
|   +-- useAuth.ts           # 認証状態を取得するカスタムフック
|   +-- useTodos.ts          # Todoリストを取得するカスタムフック
|
+-- lib
|   |
|   +-- apiClient.ts         # APIクライアントの設定（openapi-typescriptで生成）
|
+-- stores
|   |
|   +-- index.ts             # Reduxストアの設定
|
+-- types                    # 型定義ファイル
|
+-- utils                    # ユーティリティ関数

```

### reference
https://github.com/arakawamoriyuki/go-clean-handson/blob/main/clean-architecture/readme.md  
https://github.com/bxcodec/go-clean-arch/blob/master/internal/repository/mysql/article.go  
https://zenn.dev/mstn_/articles/75667657fa5aed  
