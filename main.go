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
