## graph-sdl ##

The schema definition language component of GraphQL. 

All documents and types are stored in AWS Dynamodb.

Implements specification:  https://spec.graphql.org/June2018/

## Testing ##

Over 170 test functions are include.

```
cd parser
go test  -v > test.all.log &
tail -10f test.all.log
```


