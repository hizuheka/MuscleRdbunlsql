package main

import (
  "os"

  "github.com/urfave/cli/v2"
)

func main() {
	var dbname string
	var sqlF string
	var sql string
	var input string
	var output string

	app := &cli.App{
		Name: "MuscleRdbunlsql",
		Version: Version,
		Usage: "力技でrdbunlsql",
		Flags: []cli.Flag {
			&cli.StringFlag{
				Name: "dbname",
				Aliases: []string{"d"},
				Usage: "データを取り出すデータベースの名前 `DBNAME` を指定します。",
				Destination: &dbname,
				Required: true,
			},
			&cli.StringFlag{
				Name: "sqlF",
				Aliases: []string{"v"},
				Usage: "SQLを記述したファイルのパス名(絶対パス) `SQL_FILE_PATH` を指定します。",
				Destination: &sqlF,
			},
			&cli.StringFlag{
				Name: "sql",
				Aliases: []string{"s"},
				Usage: "`SQL` を指定します。",
				Destination: &sql,
			},
			&cli.StringFlag{
				Name: "input",
				Aliases: []string{"i"},
				Usage: "データを入力するファイルのパス `INPUT_FILE_PATH` を指定します。",
				Destination: &input,
				Required: true,
			},
			&cli.StringFlag{
				Name: "output",
				Aliases: []string{"o"},
				Usage: "データを出力するファイルのパス `OUTPUT_FILE_PATH` を指定します。",
				Destination: &output,
				Required: true,
			},
		},
		Before: func(c *cli.Context) error {
			if !c.IsSet("sqlF") && !c.IsSet("sql") {
				ec := cli.Exit("sqlF か sql のどちらかを指定する必要があります。", 1)
				return ec
			}
		}
		Action: func(c *cli.Context) error {
			name := "Nefertiti"
			if c.NArg() > 0 {
				name = c.Args().Get(0)
			}
			if c.String("lang") == "spanish" {
				fmt.Println("Hola", name)
			} else {
				fmt.Println("Hello", name)
			}
			return nil
			// ec := cli.Exit("ohwell", 86)
			// fmt.Fprintf(c.App.Writer, "%d", ec.ExitCode())
			// fmt.Printf("made it!\n")
			// return ec
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

// const (
// 	conmap = "あ,Ａ\nい,Ｉ\nう,Ｕ\nえ,Ｅ\nお,О\n"
// 	target = "あaかさ1いiき2しうuくす3えeけせ4おoこそ5"
// )

// type Jef struct {
// 	fuj90     rune
// 	fuj2004Ch chan rune
// }

// func convert(r rune, conmap map[rune]rune) rune {
// 	//fmt.Printf("[%#U]:start\n", r)

// 	res := r
// 	if r > unicode.MaxASCII {
// 		if v, ok := conmap[r]; ok {
// 			res = v
// 			time.Sleep(time.Duration(rand.Intn(5)) * time.Second) // 5秒までランダムに待つ
// 		}
// 	}

// 	//fmt.Printf("[%#U]:end\n", r)
// 	return res
// }

// func genConvertMap(r io.Reader) (map[rune]rune, error) {
// 	m := make(map[rune]rune)

// 	s := bufio.NewScanner(r)
// 	for s.Scan() {
// 		ary := strings.Split(s.Text(), ",")
// 		//fmt.Printf("ary=%v, len=%d\n", ary, len(ary))
// 		if len(ary) != 2 {
// 			return nil, fmt.Errorf("変換定義ファイルのフォーマット不正. len=%d", len(ary))
// 		}
// 		fuj90 := []rune(ary[0])
// 		fuj2004 := []rune(ary[1])
// 		if len(fuj90) != 1 || len(fuj2004) != 1 {
// 			return nil, fmt.Errorf("変換定義ファイルのフォーマット不正. len(fuj90)=%d, len(fuj2004)=%d", len(fuj90), len(fuj2004))
// 		}

// 		m[fuj90[0]] = fuj2004[0]

// 	}
// 	if s.Err() != nil {
// 		// non-EOF error.
// 		return nil, s.Err()
// 	}

// 	return m, nil
// }

func gen(r io.Reader, sql string) (<-chan Unlsql) {
	out := make(chan string, 50)

	go func(br *bufio.Reader) {
		defer close(out)
		scanner := bufio.NewScanner(br)
		for scanner.Scan() {
			// sql 内のパラメタを置換する。パラメタは $1, $2 …の形式
			ary := strings.Split(s.Text(), ",")
			var s string
			for i, v := range ary {
				s = strings.Replace(sql, fmt.Sprintf("$%d", i), v, -1)
			}
		}
		select {
		case out <- Unlsql{s, make(chan []byte, 1)}:
		default:
			fmt.Println("入力ファイル用チャネルのバッファが不足しています.")
		}
		if scanner.Err() != nil {
			// non-EOF error.
			return nil, scanner.Err()
		}
		}
	}(bufio.NewReader(br))

	return out
}

// // ワーカー
// func worker(i int, src <-chan Jef, conmap map[rune]rune, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	// タスクがなくなってタスクのチェネルがcloseされるまで無限ループ
// 	for j := range src {
// 		fmt.Printf("goroutin #%d [%#U]: conv start\n", i, j.fuj90)
// 		cr := convert(j.fuj90, conmap)
// 		//fmt.Printf("%#U", cr)
// 		j.fuj2004Ch <- cr
// 		fmt.Printf("goroutin #%d [%#U]: conv end\n", i, j.fuj90)
// 	}
// }

// // 順番に従ってに書き出しをする(goroutineで実行される)
// func writeInOrder(convChs <-chan chan rune, w io.Writer, done chan<- struct{}) error {
// 	bw := bufio.NewWriter(w)
// 	defer bw.Flush()

// 	// 書き出し完了を表すチャネルをクローズする
// 	defer close(done)

// 	// 順番に取り出す
// 	for convCh := range convChs {
// 		// 選択された仕事が終わるまで待つ
// 		r := <-convCh
// 		fmt.Printf("writeInOrder:%#U\n", r)
// 		if _, err := bw.WriteRune(r); err != nil {
// 			return err
// 		}
// 		close(convCh)
// 	}

// 	return nil
// }

// // チャネル src から受信した変換対象文字を、変換用ハッシュマップ conmap を使用して変換する。
// // 変換処理はWorkersパターンを使用するが、チャネル src の順番を w に出力するために、
// // バッファありチャネル convChs を使用する。
// func processSession(src <-chan Jef, conmap map[rune]rune, w io.Writer) {
// 	// 変換結果を直列化して書き出すためのチャネル
// 	convChs := make(chan chan rune, 50)
// 	// 変換対象の文字をためるチャネル
// 	runes := make(chan Jef, 50)

// 	// 変換結果を直列化して書き出す専用のゴルーチン
// 	done := make(chan struct{})
// 	go writeInOrder(convChs, w, done)

// 	// CPU数に応じたワーカーを生成
// 	// ワーカーの数をCPU数の半分にしているのは、CPU数の半分の方が処理速度が早いと書いてあるサイトがあったから
// 	var wg sync.WaitGroup
// 	for i := 0; i < runtime.NumCPU()/2; i++ {
// 		wg.Add(1)
// 		go worker(i+1, runes, conmap, &wg)
// 	}

// 	// src から受信した変換対象文字を、ワーカーに渡すためのチャネル runes と convChs に格納していく。
// 	go func() {
// 		for r := range src {
// 			convChs <- r.fuj2004Ch
// 			runes <- r
// 		}
// 		close(convChs)
// 		close(runes)
// 	}()

// 	wg.Wait()

// 	<-done
// }

// func main() {
// 	fmt.Printf("runtime.NumCPU()=%d\n", runtime.NumCPU())

// 	m := bytes.NewBufferString(conmap)
// 	conmap, err := genConvertMap(m)

// 	//filename := os.Args[1]
// 	//f, err := os.Open(filename)
// 	//if err != nil {
// 	//	log.Fatalf("cannot open file %q: %v", filename, err)
// 	//}
// 	//defer f.Close()
// 	f := bytes.NewBufferString(target)
// 	jefs, err := gen(f)
// 	if err != nil {
// 		log.Fatalf("cannot generate convert map: %v", err)
// 	}

// 	stdout := new(bytes.Buffer)
// 	processSession(jefs, conmap, stdout)
// 	fmt.Printf("%s -> %s", target, stdout.String())
// }
