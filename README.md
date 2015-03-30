# uglify
transform multifiles JSON struct into single line JSON
This program transform a multiline JSON like :

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
