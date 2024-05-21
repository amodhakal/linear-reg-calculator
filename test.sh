go build main.go
./main input.csv display > output.txt
diff output.txt expected.txt