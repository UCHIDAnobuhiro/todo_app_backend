package domain

// Todo はアプリケーションのドメインモデルの1つで、
// ユーザーが管理するタスクを表します。
// Clean ArchitectureにおけるEntityであり、
// ビジネスルールを保持する構造体です。
// DBやフレームワークの依存情報は持ちません。
type Todo struct {
	// ID はタスクを一意に識別する番号です。
	ID uint `json:"id"`
	// Title はタスクの内容や名前を表します
	Title string `json:"title"`
	// Completed はタスクが完了しているかどうかを示します。
	Completed bool `json:"completed"`
}
