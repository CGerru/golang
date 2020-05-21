# Description
This is not a project *per se*. Instead is a series of examples using 'structs' and pointers

Structs are easy, very similar to java classes (kinda) but the tricky part comes with pointers.
There are a lot of comments in code, but here are more notes. This will truly come handy in future.

One of the main rules of Go when playing with pointers and functions is that whenever a type is 
passed as a receiver Go makes a copy of it. So structs properties cannot be altered inside a function, 
unless a pointer is required as receiver on it. 

But here comes the pain train!!!!!

When using (i.e) slices, modifications can be made inside a function, even without using pointers. 
Why is this? Well, this is because life hates you, very deeply. And also because there are some types 
in GO which behavior is different internally than the primitive types.

| Value Types |   Reference Types |
| int | slices |
| float | maps |
| string | channels |
| bool | pointers |
| struts | functions |

Explaining it quickly when a slice is created go uses (at least) two memory addresses.
1. One with the length, capacity and the pointer to the head
2. Other one with the values in an array

So, when a slice is passed to a function is the first memory address which is being passed, therefore the 
values beneath can be modified, because is the pointer inside which leads to the memory address with the data.

Despite all of this nonsense seeming hard as hell, whenever this reference values are being passed, we do not have to 
worry about pointers. Win win.