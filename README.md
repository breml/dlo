# dlo - CLI Client for dict.leo.org

dlo is a CLI client for the german - english dictionary at http://dict.leo.org

## Dockerization

If you want to run the app inside a docker container and/or have no clue
how you can compile it, but have docker running, then this simple command
will create a dockerized version of the app in no time:

        $ make -f Makefile.docker

This will build the app with the default `go` build environment as a static
binary and then this binary is used to create the runtime container with
approximately 6 MB. Afterwards you can assign a shell alias to spin up
the container and query the dictionary in a dockerized fashion:

        alias dlo="docker run --interactive --rm --tty dlo-app"

## Ideas 

* Implement some checks (language identifier in xml)
* Allow placeholder (*)
* Support Search Options (provided by leo: Grundwortschatz, Fachwortschatz)
* Support multiword search
* Select type for query (adjadv, verb, subst, example, phrase)

