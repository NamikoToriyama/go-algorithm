SRC = main.go llist/ stack/ 
OUT = a.exe

main: $(SRC)
	go build -o $(OUT)
	./$(OUT)
	
build: 
	go build -o $(OUT)

run:
	go run $(SRC)
	
clean:
	rm $(OUT)
