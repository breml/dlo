# CLI client for dict.leo.org

## Find words based on prefix

Request:

	curl 'http://dict.leo.org/dictQuery/m-query/conf/ende/query.conf/strlist.json?q=short&sort=PLa&shortQuery&noDescription&sideInfo=on&where=0&term=short'

Response:

	["short",["shortage","shortly","short","shortcoming","shortfall","shortcomings","shorts","shortcut","shorten","shortening"],[],["shortage","shortly","short","shortcoming","shortfall","shortcomings","shorts","shortcut","shorten","shortening"],[1,1,1,1,1,1,3,1,1,1]]

1 = english
2 = german
3 = both


http://codebeautify.org/xmlviewer
