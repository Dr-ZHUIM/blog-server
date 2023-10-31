package articleServer

import (
	articleController "article-server/internal/app/controllers"
	"fmt"
)

func Launch() {
	fmt.Println("server started, listenning to port 8080...")
	articleController.ArticleServices()
}
