package articleServer

import (
	Controller "article-server/internal/app/controllers"
	"fmt"
)

func Launch() {
	fmt.Println("---------------------------------------------------")
	fmt.Println("server started, listenning to port 8080...")
	Controller.GetServices()
	fmt.Println("---------------------------------------------------")
}
