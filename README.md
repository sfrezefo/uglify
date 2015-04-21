NAME
----
    uglify  
    
SYNOPSIS
--------
    uglify -inputfile='test1.json' -outputfile='out.json' -jsonarray

DESCRIPTION
-----------

    uglify transforms multi-lines JSON struct into single line JSON
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
      
    into a single line JSON :
        {"a":"1","b":"2","c":"3","d":[3,45,{"h":"aze"}],"e":{"f":45,"iiii":["zer","ert",555],"e":3}}

    The options are as follows:
               
    **-inputfile** 
    **-outputfile**
    **-jsonarray**

                        
USAGE
-----
    uglify -inputfile='test1.json' -outputfile='out.json' -jsonarray
    cat test1.json | ./uglify  -outputfile='out.json' -jsonarray
    cat test1.json | ./uglify   -jsonarray
    cat test1.json | ./uglify# uglify
