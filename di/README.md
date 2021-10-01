# DI

- io.Writerを引数の型にすることで利用側から自由に出力先を選択できる
- つまり利用側から出力したい場所を注入できる（Http、Stdout...）
- これが依存関係の注入
- 依存関係を注入するには汎用的なInterfaceが必要
- Goの場合はio.Writerがある