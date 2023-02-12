package main

import (
	"flag"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/datsukan/datsukan-blog-comment-core/usecase"
	"github.com/datsukan/datsukan-blog-comment-post/controller"
	"github.com/joho/godotenv"
)

func main() {
	t := flag.Bool("local", false, "ローカル実行か否か")
	articleID := flag.String("article-id", "", "ローカル実行用の記事ID")
	parentID := flag.String("parent-id", "", "ローカル実行用の返信元コメントID")
	userName := flag.String("user-name", "", "ローカル実行用の表示名")
	content := flag.String("content", "", "ローカル実行用のコメント内容")
	flag.Parse()

	isLocal, err := isLocal(t, articleID, parentID, userName, content)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if isLocal {
		fmt.Println("local")
		localController(articleID, parentID, userName, content)
		return
	}

	fmt.Println("production")
	lambda.Start(controller.Post)
}

// isLocal はローカル環境の実行であるかを判定する。
func isLocal(t *bool, articleID *string, parentID *string, userName *string, content *string) (bool, error) {
	if !*t {
		return false, nil
	}

	if *articleID == "" {
		fmt.Println("no exec")
		return false, fmt.Errorf("ローカル実行だが記事ID指定が無いので処理不可能")
	}

	if *userName == "" {
		fmt.Println("no exec")
		return false, fmt.Errorf("ローカル実行だが表示名指定が無いので処理不可能")
	}

	if *content == "" {
		fmt.Println("no exec")
		return false, fmt.Errorf("ローカル実行だがコメント内容指定が無いので処理不可能")
	}

	return true, nil
}

// localController はローカル環境での実行処理を行う。
func localController(articleID *string, parentID *string, userName *string, content *string) {
	if err := godotenv.Load(); err != nil {
		fmt.Printf("環境変数を読み込み出来ませんでした: %v\n", err)
		return
	}

	c, err := usecase.Post(*articleID, *parentID, *userName, *content)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(c)
}
