package internal

type Options struct {
	BotID          string `long:"bot_id" env:"BOT_ID"`
	NewsChatID     int64  `long:"news_chat_id" env:"NEWS_CHAT_ID"`
	NotifierChatID int64  `long:"notifier_chat_id" env:"NOTIFIER_CHAT_ID"`
	SongApiKey     string `long:"song_api_key" env:"SONG_API_KEY"`
	DBUrl          string `long:"db_url" env:"DB_URL"`
	OnlyNotifier   bool   `long:"only_notifier" env:"ONLY_NOTIFIER"`
}
