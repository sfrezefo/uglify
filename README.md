NAME
----
    uglify 
 
SYNOPSIS
--------
    uglify -inputfile='test1.json' -outputfile='out.json' -jsonarray

DESCRIPTION
-----------

    uglify transforms multi-lines(unformated) JSON struct into single line JSON
    A multi-lines JSON file like :
        {
            "a":"1",
            "b":"2",
            "c":"3",
            "d":[ 3 , 45,{"h":"aze"}],
            "e":{ "e":3 , "f":45,
            "iiii"    :["zer","ert",
            555]}
        }
      
    will be transformed into a single line JSON :
        {"a":"1","b":"2","c":"3","d":[3,45,{"h":"aze"}],"e":{"f":45,"iiii":["zer","ert",555],"e":3}}
    with the -prettyprint option it can also pretty print 

    The options are as follows:
               
    **-inputfile**  is the file name of the json file to uglify or pretty prin. if ommitted standard input will be used
    **-outputfile** is the file name of the output json file. if ommitted standard input will be used
    **-jsonarray** specify if the output json should be surrounded by an array element
    **-prettyprint** with the -prettyprint option it can also pretty print 
    **-debuglevel**


USAGE
-----
    Examples of usage with file names or standard input and output, and with boolean parameter to determine:
    uglify -inputfile='test1.json' -outputfile='out.json' -jsonarray
    cat test1.json | uglify -outputfile='out.json' -jsonarray
    cat test1.json | uglify -jsonarray
    cat test1.json | uglify
    uglify -inputfile='test1.json' -outputfile='out.json' -jsonarray
    uglify -inputfile='test1.json' -outputfile='out.json' -jsonarray -prettyprint
