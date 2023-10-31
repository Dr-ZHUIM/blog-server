BINARY_NAME=./build/server.out

compiler: 
	go build -o ${BINARY_NAME} main.go

dev:
	go run main.go

run: 
	make compiler
	./${BINARY_NAME}

test:
	curl -X POST http://localhost:8080/GetArticleList

	curl http://localhost:8080/addArticle \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"id": "4","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}'

	curl -X POST http://localhost:8080/getArticle/2

clean:
	go clean
	rm ${BINARY_NAME}