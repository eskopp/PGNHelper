#!/bin/bash

git checkout main
git pull origin main
git checkout work
git merge main
git pull origin work
git checkout main
git merge work
