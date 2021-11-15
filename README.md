# MOGEE
**mogee (pronounced ēˈmōjē) is an extremely simple interpreted programming language that consists solely of emojis.** <br />
___
<br />

Similar to [Brainfuck](https://en.wikipedia.org/wiki/Brainfuck), mogee is also a minimal language that uses an instruction pointer. It has a few built in functions to manipulate memory in certain ways. <br />
<br />


### The Tape
One of the key concepts of mogee is the **tape**. Every program initializes with a 'tape' that consists of 10 000 bytes. You can visualize the tape like this:

```
┌───────────────────────────────────>
│  0  0  0  0  0  0  0  0  0   ...
└───────────────────────────────────>
```

Every zero above represents a **byte** that can get a value between 0 and 255. There are 10 000 bytes in every tape in mogee programs.
### Manipulating The Tape
Another important concept in mogee is manipulation of the tape. This is done through the pointer.

#### The Pointer
The pointer just points to a specific byte in the tape:
```
┌───────────────────────────────────>
│  0  0  0  0  0  0  0  0  0   ...
└───────────────────────────────────>
   ▲   <-- The pointer
```
The pointer starts off at index 0 when the program starts. You can move the pointer in both directions using the following commands:

```
👉 <-- move pointer right 
👈 <-- move pointer left 
🚘 <-- move pointer right 10 cells 
🚗 <-- move pointer left 10 cells 
```

For example, 👉 will produce the following:
```
┌───────────────────────────────────>
│  0  0  0  0  0  0  0  0  0   ...
└───────────────────────────────────>
      ▲
```
Similarly you can use 👈 to move the pointer left:
```
┌───────────────────────────────────>
│  0  0  0  0  0  0  0  0  0   ...
└───────────────────────────────────>
   ▲
```
**TIP:** Functions 🚘 and 🚗, move pointer right/left 10 cells respectively, are a shorthand way to move the pointer around. </br>
<br />

#### Incrementing and Decrementing cells
Every cell we looked at until now was at the initial value, 0. To increase or decrease this value, we use the following functions:
```
👍 <-- increment current cell 
👎 <-- decrement current cell 
```
For example, the cell with the pointer on will increase by 3 when the following program is executed:
```
👍👍👍
```
Combining the concept of the pointer and incrementing/decrementing cells, the following program should output this:
```
CODE: 👉👉👍👍👍👎

OUTPUT:
┌───────────────────────────────────>
│  0  0  2  0  0  0  0  0  0   ...
└───────────────────────────────────>
         ▲
```
There are also shorthand ways to do cell manipulation:
```
✋ <-- add 5 to current cell 
🤚 <-- subtract 5 from current cell 
🔵 <-- add 10 to current cell 
🟦 <-- subtract 10 from current cell 
🔴 <-- add 100 to current cell 
🟥 <-- subtract 100 from current cell 
```
These functions behave the same way 👍 and 👎 do, but instead of incrementing and decrementing by 1, they increment and decrement by 5, 10, and 100 respectively. For example:
```
🔴🔵✋👎
```
will be 114:
```
┌───────────────────────────────────>
│  114  0  0  0  0  0  0  0  0   ...
└───────────────────────────────────>
    ▲
```
### Comments
all text (and undefined emojis too) are considered to be comments.
### I/O
To write a string to stdout, or to get user input, we use the following commands:
```
📖 <-- replace current cell with user input 
📝 <-- print current cell content (ascii) 
🧮 <-- print current cell content (bytes) 
```
#### Input
The 📖 function asks for user input and requires a byte. Just enter a number between 0 and 255 to represent a byte. The value provided will replace the byte the pointer is currently on.

### Output
🧮 Simply prints out the current cell directly as a byte and not as an ascii value:
```
🔴🧮
```
will output `100`. <br />
<br />
Similarly, 📝 also prints the current cell, but in ascii format.
```
🔴📝
```
will `d` since the ascii code for lowercase d is 100.

### Printing letters using ASCII
If you are not familiar with ascii codes, they are simply representations of letters in bytes. Read more about ascii [here](https://tr.wikipedia.org/wiki/ASCII). <br />
<br />
Here is an ascii table:
```
ASCII        BINARY           SYMBOL
────────────────────────────────────────
65		    01000001     	Uppercase A
66		    01000010     	Uppercase B
67		    01000011     	Uppercase C
68		    01000100     	Uppercase D
69		    01000101     	Uppercase E
70		    01000110     	Uppercase F
71		    01000111     	Uppercase G
72		    01001000     	Uppercase H

... 
```
> Note that this table is very much incomplete and there are many more ascii characters, please refer to a complete ascii table.
<br />

You will likely refer to [one of these](https://www.ascii-code.com/) when printing letters. Just set a cell to the specific ascii value you want and print it.

#### Hello World!
Considering you have understood the following concepts, you should be able to write your hello world program:
```
CODE:
🔴🟦🟦🤚👎👎👎     <--  H 📝👉
🔴👍                <-- e 📝👉
🔴✋👍👍👍         <--  l 📝👉
🔴✋👍👍👍         <--  l 📝👉
🔴🔵👍              <-- o 📝👉
🔵🔵🔵👍👍          <-- space 📝👉
🔴🟦👎👎👎          <-- W 📝👉
🔴🔵👍              <-- o 📝👉
🔴🔵✋👎            <-- r 📝👉
🔴✋👍👍👍         <--  l 📝👉
🔴                  <-- d 📝👉
🔵🔵🔵✋👎👎        <-- ! 📝👉

OUTPUT:
Hello World!
```

### Helpers
```
🧿 <-- set current cell to 255 
🚫 <-- set current cell to 0 
📼 <-- print the tape 
🚿 <-- clear the whole tape 
❌ <-- exit program 
```
These helpers do exactly what their labels say they do, no real explanation required.
### Functions
Functions in mogee is pretty straight forward. There are no parameters (yet) and functions are just meant to be used to avoid code duplication. Here is an example function that simply adds 11 to the current cell: <br />
```
Example function: add 11
👇🔥
  🔵👍
👆
```
#### Defining Functions
Let's dive into the anatomy of this example function:<br />
Every function has a start and an end defined by 👇 and 👆. Everything in between these two emojis are considered to be the **function body**.
```
👇
This is the function body!
👆
```
To identify and call this function later, we need to give it a **name**. Function names in mogee should be an emoji too as all text is a comment and therefore ignored. Let's name our function 🔥. To do this, we have to add the name right after the function open (👇).
```
👇🔥 The name of this function is (fire emoji)
This is the function body!
👆
```
**TIP:** You *must* assign the function name right after 👇. There should be no space or any other character between 👇 and the function name <br />
<br />
Now that we gave our function a name, we can write the actual body of the function and make it do something useful. Just write any mogee expression in the function body.
```
👇🔥
  🔵👍 Add 10 and 1
👆
```
Note that you cannot define other functions inside a mogee function and all functions should be top-level
#### Calling Functions
Calling functions in mogee is very easy. Just use the 📞 keyword followed by the function name.
```
📞🔥 call the (fire emoji) function
```
