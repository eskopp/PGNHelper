del *.exe
go test -v
go build -o PGNHelper.exe
.\PGNHelper.exe -parsePGN
