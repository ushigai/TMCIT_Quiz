# テーマ
- yayyay
- [GitHub](https://github.com/ushigai/TMCIT_Quiz)
- [スライド](https://ushigai.github.io/TMCIT_Quiz/)

# ヒアリングの内容
# 要件定義

# 機能
## 概要
- クイズができるよ
    - クイズ作成機能（curlリクエスト）
    - `docker entrypoint`
- クイズが解けるよ
    - 4択クイズのみ
    - [正解|不正解]が表示される
    - 選択までの秒数が表示される
- クイズ
- 懺悔
    - Web上での`Quiz`とか`FaQs`とかサボってしまった

## 画面遷移図
![](https://github.com/ushigai/TMCIT_Quiz/tree/main/doc/img/ushigai/home.png)
![](https://github.com/ushigai/TMCIT_Quiz/tree/main/doc/img/ushigai/lobby.png)
![](https://github.com/ushigai/TMCIT_Quiz/tree/main/doc/img/ushigai/room.png)
![](https://github.com/ushigai/TMCIT_Quiz/tree/main/doc/img/ushigai/quiz1.png)
![](https://github.com/ushigai/TMCIT_Quiz/tree/main/doc/img/ushigai/quiz_correct.png)
![](https://github.com/ushigai/TMCIT_Quiz/tree/main/doc/img/ushigai/quiz_wrong.png)

## データベース
- `rooms`では各`quiz`にアクセスするための部屋情報が格納されている
- `rooms`には`title`と`subtitle`とあり、`英語検定`,`5級`のように設定するとフロントでいい感じになる
- hoge
- ER図
![](https://github.com/ushigai/TMCIT_Quiz/tree/main/doc/img/er.png)

# 技術スタック
## 使用ソフトウェア/プログラミング言語

| カテゴリー | ソフトウェア/プログラミング言語 |
| -------- | -------- |
| Webフロント | HTML,CSS,JavaScript |
| バックエンド | Go(echo) |
| O/Rマッパ | gorm |
| データベース | MySQL |
| 環境構築 | Docker |
| コンテナオーケストレーションツール | Kubernetes |
| コード管理 | GitHub |

## インフラアーキテクチャ図
![](https://github.com/ushigai/TMCIT_Quiz/tree/main/doc/img/infra_architecture.png)
