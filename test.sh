go build main.go
./main input.csv > output.txt
sdiff output.txt expected.txt