# PGN Helper

[![Windows](https://github.com/eskopp/PGNHelper/actions/workflows/windows.yml/badge.svg)](https://github.com/eskopp/PGNHelper/actions/workflows/windows.yml) [![Linux](https://github.com/eskopp/PGNHelper/actions/workflows/linux.yml/badge.svg)](https://github.com/eskopp/PGNHelper/actions/workflows/linux.yml) [![Mac](https://github.com/eskopp/PGNHelper/actions/workflows/Mac.yml/badge.svg)](https://github.com/eskopp/PGNHelper/actions/workflows/Mac.yml)
[![wakatime](https://wakatime.com/badge/user/018d28f6-efdf-4097-b633-b84af841166c/project/018d3b8e-3298-4faa-87c7-b494ad230448.svg)](https://wakatime.com/badge/user/018d28f6-efdf-4097-b633-b84af841166c/project/018d3b8e-3298-4faa-87c7-b494ad230448)

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
The function removes the Event Date attribute from the PGN files. The attribute generates errors in ChessBase and ChessX for files that are connected to DGT boards and the Lichess API.
```bash
pgn -EventDate test.pgn
pgn -EventDate test.json
```
The flag accepts both pgn and json files.  The basic problem is that the EventDate attribute overwrites the other Date attribute. As a result, the wrong date is displayed in PGN files.


### Create / Delete PGN or JSON file
> [!NOTE]  
> The function is not yet included in the current release.

> [!TIP]  
> New Syntax.

The flag creates or deletes pgn or json files.
```bash
pgn -create dummy.pgn 
pgn -delete dummy.pgn 
```


### Remove Chessbase Files
> [!NOTE]  
> The function is not yet included in the current release.

> [!TIP]  
> Chessbase does not need these files, if you open the PGN file again Chessbase will create the files again. You can only delete the files if you have closed Chessbase.

The flag deletes all pgi and ini files of Chessbase from the folder. This leaves the PGN files. Chessbase archives and CBH archives remain unaffected.
```bash
pgn -removecb
```