build:
	go build -o dn src/main.go
clean:
	rm -rf dn
first run:
	touch movie.db && sqlite3 movie.db -init init.sql
