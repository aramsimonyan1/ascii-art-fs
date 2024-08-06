# ascii-art-fs
Go program that processes text inputs as an argument and generates ASCII art representations of the string using predefined patterns from external file. Additionally, users can customise the final output by selecting font-style from multiple ASCII art templates through terminal input.

## Objectives
Graphic representation using ASCII means to write the string received using ASCII characters, as you can see in the example below:
###
    @@@@@@BB@@@@``^^``^^``@@BB$$@@BB$$
    @@%%$$$$^^^^WW&&8888&&^^""BBBB@@@@
    @@@@@@""WW8888&&WW888888WW``@@@@$$
    BB$$``&&&&WWWW8888&&&&8888&&``@@@@
    $$``&&WW88&&88&&&&8888&&88WW88``$$
    @@""&&&&&&&&88888888&&&&&&88&&``$$
    ``````^^``^^^^^^````""^^``^^``^^``
    ""WW^^@@@@^^``````^^BB@@^^``^^&&``
    ^^&&^^@@````^^``&&``@@````^^^^&&``
    ``WW&&^^""``^^WW&&&&""``^^^^&&88``
    ^^8888&&&&&&WW88&&88WW&&&&88&&WW``
    @@``&&88888888WW&&WW88&&88WW88^^$$
    @@""88&&&&&&&&888888&&``^^&&88``$$
    @@@@^^&&&&&&""``^^^^^^8888&&^^@@@@
    @@@@@@^^888888&&88&&&&MM88^^BB$$$$
    @@@@@@BB````&&&&&&&&88""``BB@@BB$$
    $$@@$$$$$$$$``````````@@$$@@$$$$$$

This project should handle an input with numbers, letters, spaces, special characters and \n. 

## Instructions
Your project must be written in Go.
The code must respect the good practices.
Some banner files with a specific graphical template representation using ASCII will be given. The files are formatted in a way that is not necessary to change them.
### 
    shadow
    standard
    thinkertoy

## Banner Format
Each character has a height of 8 lines.
Characters are separated by a new line \n.
Here is an example of ' ', '!' and '"'(one dot represents one space):
###
    ......
    ......
    ......
    ......
    ......
    ......
    ......
    ......

    ._..
    |.|.
    |.|.
    |.|.
    |_|.
    (_).
    ....
    ....

    ._._..
    (.|.).
    .V.V..
    ......
    ......
    ......
    ......
    ......


The second argument of a user input must be the name of the template/banner/font. 
The usage must respect this format go run . [STRING] [BANNER], any other formats must return the following usage message:
    Usage: go run . [STRING] [BANNER]
    EX: go run . something standard


## Usage
###
    $ go run . "hello" standard
     _              _   _          $
    | |            | | | |         $
    | |__     ___  | | | |   ___   $
    |  _ \   / _ \ | | | |  / _ \  $
    | | | | |  __/ | | | | | (_) | $
    |_| |_|  \___| |_| |_|  \___/  $
                                   $
                                   $

    $ go run . "Hello There!" shadow
                                                                                             $
    _|    _|          _| _|                _|_|_|_|_| _|                                  _| $
    _|    _|   _|_|   _| _|   _|_|             _|     _|_|_|     _|_|   _|  _|_|   _|_|   _| $
    _|_|_|_| _|_|_|_| _| _| _|    _|           _|     _|    _| _|_|_|_| _|_|     _|_|_|_| _| $
    _|    _| _|       _| _| _|    _|           _|     _|    _| _|       _|       _|          $
    _|    _|   _|_|_| _| _|   _|_|             _|     _|    _|   _|_|_| _|         _|_|_| _| $
                                                                                             $
                                                                                             $

    $ go run . "Hello There!" thinkertoy
                                                    $
    o  o     o o           o-O-o o                o $
    |  |     | |             |   |                | $
    O--O o-o | | o-o         |   O--o o-o o-o o-o o $
    |  | |-' | | | |         |   |  | |-' |   |-'   $
    o  o o-o o o o-o         o   o  o o-o o   o-o O $
                                                    $
                                                    $

# Allowed packages
Only the standard Go packages are allowed

This project will help you learn about:
The Go file system(fs) API
Data manipulation