#! /bin/bash

BIN="App_linux" 
DIR="./$BIN"
DATA="./$DIR/data"
CONF="./configs.json"
INPUT="./data/airserv-adr.txt"

# clean up
rm -r $DIR

# make build directory
mkdir $DIR 
mkdir $DATA

# compile binary
env GOOS=linux GOARCH=amd64 go build -o ./$DIR/$BIN

# copy dependancies
cp $INPUT $DATA
cp $CONF $DIR

