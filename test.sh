#!/usr/bin/env bash
blue='\033[1;34m'
NC='\033[0m'
printf "${blue}Run Example00 \n\n${NC}"
go run . example00.txt | cat -e
echo "----------------------" 
sleep 2
printf "${blue}Run Example01 \n\n${NC}"
go run . example01.txt | cat -e
echo "----------------------"
sleep 2
printf "${blue}Run Example02 \n\n${NC}"
go run . example02.txt | cat -e
echo "----------------------"
sleep 2
printf "${blue}Run Example03 \n\n${NC}"
go run . example03.txt | cat -e
echo "----------------------"
sleep 2
printf "${blue}Run Example04 \n\n${NC}"
go run . example04.txt | cat -e
echo "----------------------"
sleep 2
printf "${blue}Run Example05 \n\n${NC}"
go run . example05.txt | cat -e
echo "----------------------"
sleep 2
printf "${blue}Run Badexample00 \n\n${NC}"
go run . badexample00.txt | cat -e
echo "----------------------"
sleep 2
printf "${blue}Run Badexample01 \n\n${NC}"
go run . badexample01.txt | cat -e
echo "----------------------"
sleep 2
printf "${blue}Run Example06 with 100 ants \n\n${NC}"
go run . example06.txt | cat -e
echo "----------------------"
sleep 2
printf "${blue}Run Example07 with 1000 ants \n\n${NC}"
go run . example07.txt | cat -e