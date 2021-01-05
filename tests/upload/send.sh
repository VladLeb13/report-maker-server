#!/bin/bash

cat structs.json | curl -X POST --data @- -H "Content-Type: application/json" -H "Authorization: Basic Y2xpZW50OmNsaWVudHBhc3M=" http://localhost:8080/upload
#cat structs.json | curl -X POST --data @- -H "Content-Type: application/json" -H "Authorization: Basic Y2xpZW50czpjbGllbnRwYXNz" http://localhost:8080/upload
#cat structs.json | curl -X POST --data @- -H "Content-Type: application/json"  http://localhost:8080/upload
