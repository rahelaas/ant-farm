# Ant-farm 
(original name Lem-in)

## Objectives
This project is a digital version of an ant farm. It was part of the projects in the Golang module.

The program lem-in reads from a file (describing the ants and the colony) given in the arguments.

Upon successfully finding the quickest path, lem-in will display the content of the file passed as argument and each move the ants make from room to room.

[Full description of the task](https://github.com/01-edu/public/tree/master/subjects/lem-in)


**How does it work?**

The ant farm is made with tunnels and rooms.
You place the ants on one side and look at how they find the exit.


It finds the best possible way for ants to get from start to end. 
## How to run
To run the program it is necessary to first download Golang.

**Command line commands:**

>`./test.sh` - to test out all files

>`go run . "filename.txt"` - to test out one file at a time

> `EX: go run . example00.txt`


## Example00 output
```
4
##start
0 0 3
2 2 5
3 4 0
##end
1 8 3
0-2
2-3
3-1

L1-2
L1-3 L2-2
L1-1 L2-3 L3-2
L2-1 L3-3 L4-2
L3-1 L4-3
L4-1
```

### Authors

* @rahela
* @kaspars
* @kerlz 