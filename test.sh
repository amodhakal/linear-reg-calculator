go build main.go
./main input.csv > output.txt
diff output.txt expected.txt