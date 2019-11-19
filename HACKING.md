# Computing History

Alternative names:
 - Historical figures in Computing
 - History of Computing (in faces)


## Short description
Concise history of computing in persons in Human and Machine readable formats.

A curated list of people and their works, who were instrumental for having computers as we know them now.


## Motivation

Help broadening understanding of historical context of "computing", sample "powerfull ideas" from the past. How did we get "here" in our field?


## Details

This is more of a Media project, than a Technology one as it includes:
 - collecting data about important historical figures of computing, starting with Turing Award winners and
 - distilling it to short summary
 - manually curating best sources for further reading.

Each person has a small page with a summary of his contributions and links to software / most influencial papers / videos, lectures, etc any other materials for further reading and firsthand experience.

Can be a single .md document for every person, and then a timeline visualization can be generated. Or a web site. Or a graph database. Or any other form of structure knowlege representation.

Inspired by recent interest in Doublas Engelbart works but there are many more people who's worth is worth studing!

Having it in some moldable human/machine readable format (sematic web?) + all external links perseved from archive.org would help to perserve and promote this heritage.


# People
 - 51 Turing award winners
 - DARPA/Xerox Parc folks
   ["Power of contect"](http://www.vpri.org/pdf/m2004001_power.pdf)
 - https://docs.google.com/spreadsheets/d/1kNkiS2Xxi5TZJ_5yEMpgPOWCDrIlJgmmdpnIGP6GhjI/edit#gid=1933526947


# Sources
 - https://amturing.acm.org/byyear.cfm
 - https://github.com/devinmcgloin/theory/tree/master/worrydream
 - https://www.ibiblio.org/pioneers/index.html
 - https://ethw.org/Category:Computing_and_electronics
 - https://en.wikipedia.org/wiki/List_of_pioneers_in_computer_science
 - http://mrmgroup.cs.princeton.edu/cos583/syllabusS15.pdf

 # Timelines
 - http://www.computerhistory.org/timeline/computers/
 - https://www.denizcemonduygu.com/philo/browse/


---------------------


 - add Alan Perils, John Tukey
 - add Norbert Wiener, Marvin Minsky, John McCarthy
 - Richard Hamming, Ted Nelson, Alan Kay
 - generate xrefs for names in *.md

 ```
   parse all .md, get
    person
    book
    article
 ```
 - automate person template structure refactoring?


 - content in TOML + rendering to .md
 - nice <img> position

 - a graph-based model + visualization
   person
    :name
    :slug
    :period
    studied
    worked
    wrote
   place
   article/book
   fact

A Graph DB + Query language

https://en.wikipedia.org/wiki/Wikipedia:List_of_infoboxes

https://en.wikipedia.org/wiki/Template:Infobox_person
https://en.wikipedia.org/wiki/Template:Infobox_scientist

DBPedia has already scrapped it, using Apache Spark!!!
 https://github.com/dbpedia/extraction-framework#about-dbpedia
 https://github.com/dbpedia/extraction-framework#dbpedia-extraction-framework-now-powered-by-apache-spark

<http://dbpedia.org/ontology/notableStudent>
<http://dbpedia.org/ontology/doctoralStudent>
<http://dbpedia.org/ontology/influenced>

<http://dbpedia.org/resource/Vannevar_Bush>	<http://dbpedia.org/ontology/notableStudent>	<http://dbpedia.org/resource/Frederick_Terman> .

```
g.V().Has("<http://www.w3.org/1999/02/22-rdf-syntax-ns#type>", "<http://dbpedia.org/ontology/Person>").All()

g.V()
 .Has("<http://www.w3.org/1999/02/22-rdf-syntax-ns#type>", "<http://dbpedia.org/ontology/Person>")
 .Tag("source")
 .Out("<http://dbpedia.org/ontology/notableStudent>")
 .Tag("target")
 .All()
```

