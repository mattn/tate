package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"strings"

	"golang.org/x/text/width"

	"github.com/mattn/go-isatty"
	"github.com/mattn/go-runewidth"
)

var replacerWin = strings.NewReplacer(
	` `, `　`,
	`↑`, `→`,
	`↓`, `←`,
	`←`, `↑`,
	`→`, `↓`,
	`ー`, `｜`,
	`─`, `｜`,
	`−`, `｜`,
	`－`, `｜`,
	`—`, `︱`,
	`〜`, `∫`,
	`～`, `∫`,
	`／`, `＼`,
	`…`, `:`,
	`‥`, `:`,
	`：`, `--`,
	`:`, `--`,
	`；`, `--`,
	`;`, `--`,
	`＝`, `||`,
	`=`, `||`,
	`-`, `|`,
	`ｰ`, `|`,
	`_`, `|`,
	`,`, "`",
	`､`, "`",
)

var replacerUtf8 = strings.NewReplacer(
	` `, `　`,
	`↑`, `→`,
	`↓`, `←`,
	`←`, `↑`,
	`→`, `↓`,
	`。`, `︒`,
	`、`, `︑`,
	`ー`, `｜`,
	`─`, `｜`,
	`−`, `｜`,
	`－`, `｜`,
	`—`, `︱`,
	`〜`, `∫`,
	`～`, `∫`,
	`／`, `＼`,
	`…`, `︙`,
	`‥`, `︰`,
	`：`, `︓`,
	`:`, `︓`,
	`；`, `︔`,
	`;`, `︔`,
	`＝`, `॥`,
	`=`, `॥`,
	`（`, `︵`,
	`(`, `︵`,
	`）`, `︶`,
	`)`, `︶`,
	`［`, `﹇`,
	`[`, `﹇`,
	`］`, `﹈`,
	`]`, `﹈`,
	`｛`, `︷`,
	`{`, `︷`,
	`＜`, `︿`,
	`<`, `︿`,
	`＞`, `﹀`,
	`>`, `﹀`,
	`｝`, `︸`,
	`}`, `︸`,
	`「`, `﹁`,
	`」`, `﹂`,
	`『`, `﹃`,
	`』`, `﹄`,
	`【`, `︻`,
	`】`, `︼`,
	`〖`, `︗`,
	`〗`, `︘`,
	`｢`, `﹁`,
	`｣`, `﹂`,
	`-`, `| `,
	`ｰ`, `| `,
	`_`, `| `,
	`,`, `︐`,
	`､`, `︑`,
)

func main() {
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	isCmd := runtime.GOOS == "windows" && isatty.IsTerminal(os.Stdout.Fd()) && !isatty.IsCygwinTerminal(os.Stdin.Fd())

	s := strings.TrimSpace(strings.Replace(string(b), "\r", "", -1))
	if isCmd {
		s = replacerWin.Replace(s)
	} else {
		s = replacerUtf8.Replace(s)
	}
	lines := strings.Split(s, "\n")

	max := 0
	for _, l := range lines {
		w := len([]rune(l))
		if w > max {
			max = w
		}
	}

	for i := 0; i < max; i++ {
		for j := len(lines) - 1; j >= 0; j-- {
			rs := []rune(lines[j])
			if i < len(rs) {
				r := width.LookupRune(rs[i]).Wide()
				if r == 0 {
					r = rs[i]
				}
				if runewidth.RuneWidth(r) > 1 {
					fmt.Print(string(r))
				} else {
					fmt.Print(" " + string(r))
				}
			} else {
				fmt.Print("　")
			}
		}
		fmt.Println()
	}
}
