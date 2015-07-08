package main

type LangPkg map[string]string

var chinesePkg = LangPkg{
	"snake":        "贪 食 蛇",
	"move_tip":     "使 用  h j k l 进 行 移 动",
	"quit_tip":     "使 用  q 强 制 退 出 游 戏",
	"start_game":   "开 始 游 戏",
	"end_game":     "退 出 游 戏",
	"you_lose":     "你 输 了 ！ ",
	"back_to_menu": "使 用 空 格 键 返 回 菜 单 ",
}

var englishPkg = LangPkg{
	"snake":        "Snake",
	"move_tip":     "Use h, j, k, l to move",
	"quit_tip":     "Use  q to quit game forcibly",
	"start_game":   "start game",
	"end_game":     "quit game",
	"you_lose":     "You lose!",
	"back_to_menu": "Use SPACE to come back to main menu",
}

func SetGloablLang(lang string) {
	switch lang {
	case "chinese":
		gLangPkg = chinesePkg

	case "english":
		gLangPkg = englishPkg

	default:
		panic("no this language package")
	}
}

var gLangPkg LangPkg
