> [!WARNING]  
> The project is still in progress. There is no realease and no stable version yet. The current use is at your own risk.

# PGN Helper

[![Windows](https://github.com/eskopp/PGNHelper/actions/workflows/windows.yml/badge.svg)](https://github.com/eskopp/PGNHelper/actions/workflows/windows.yml) [![Linux](https://github.com/eskopp/PGNHelper/actions/workflows/linux.yml/badge.svg)](https://github.com/eskopp/PGNHelper/actions/workflows/linux.yml) [![Mac](https://github.com/eskopp/PGNHelper/actions/workflows/Mac.yml/badge.svg)](https://github.com/eskopp/PGNHelper/actions/workflows/Mac.yml)

## Build

```bash
git clone https://github.com/eskopp/PGNHelper.git
cd PGNHelper
git checkout main
git pull 
go mod tidy
go test -v
go build -o PGNHelper
```

## Install
```bash
sudo cp PGNHelper /usr/local/bin/pgn
```

## Usage

### Help
The ``help`` flag allows you to view the help view. There you will find all information and hints.
```bash
pgn -help
```


### Parse PGN 
The ``-parsePGN`` flag creates a JSON file from the PGN file. JSON files cannot be read by normal chess programs but are much less error-prone and more efficient to process and store. PGN and JSON files can be converted into each other at any time.
```bash
pgn -parsePGN test.pgn output.json
```
If no output json is specified, the name of the PGN is adopted. 



### Parse JSON
The ``-parseJSON`` flag creates a PGN file from the JSON file. JSON files cannot be read by normal chess programs but are much less error-prone and more efficient to process and store. PGN and JSON files can be converted into each other at any time.
```bash
pgn -parseJSON test.json output.pgn
```
If no output pgn is specified, the name of the json is adopted. 


### Delete Event Date
> [!CAUTION]  
> The flag is not yet working. Errors occur when reading the input file. This function is not available in the current release.

The function removes the Event Date attribute from the PGN files. The attribute generates errors in ChessBase and ChessX for files that are connected to DGT boards and the Lichess API.
```bash
pgn -EventDate test.pgn
```
or
```bash
pgn -EventDate test.json
```
The flag accepts both pgn and json files. 
