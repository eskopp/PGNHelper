# git checkout main (Debug)
git pull
go mod tidy
go test -v
go build -o PGNHelper
sudo mv PGNHelper /usr/local/bin/pgn
zsh
