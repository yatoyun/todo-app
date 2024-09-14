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
│   │   │   └── todo.go
│   │   └── usecase                  Usecase
│   │       ├── todo_input_port.go   Input PortのインターフェースとInput Dataの構造定義。ドメインが望む入力形式の定義
│   │       ├── todo_interactor.go   Interactor。Input DataとRepositoryを使ってデータを取得、Output Portの実装であるPresenterに渡す
│   │       └── todo_output_port.go  Output PortのインターフェースとOutput Dataの構造定義。。ドメインが望む出力形式の定義
│   └── entity                       Enterprise Business Rules
│       └── todo.go                  データ構造。DDDにおける値オブジェクトやドメインサービスまで定義していいと思う。
├── go.mod
├── go.sum
├── infrastructure                   Frameworks & Drivers
│   └── router                       Ginの実装。Controllerと繋げる
│       ├── api
│       │   └── todo.go
│       └── router.go
├── interface                        Interface Adapters
│   ├── controller                   Controller。データをInteractorが望むInputDataに整形して渡し、Presenterが整形したデータを表示する
│   │   ├── context.go
│   │   ├── error.go
│   │   └── todo.go
│   ├── presenter                    Presenter。Output Portの実装。OutDataを受け取ってレスポンスを整形する
│   │   └── todo.go
│   └── repository
│       └── todo.go                  リポジトリの実装。Databaseを使ったデータの取得など
├── main.go
└── migrations
    ├── 000001_create_todos.down.sql
    └── 000001_create_todos.up.sql
```


### reference
https://github.com/arakawamoriyuki/go-clean-handson/blob/main/clean-architecture/readme.md