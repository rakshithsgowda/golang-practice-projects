# JSON

- package used
  ---> ("encoding/json")

## Data Types

default golang data types for decoding and emcoding json are:-

- bool for JSON booleans
- Int/float fro JSON numbers.
- string for JSON strings.
- nil for JSON null
- array as JSON array
- map or struct as JSON object

### parsing a JSON in golang

We change the JSON code which is in key value pairs to struct type in golang.using,

- Marshal(encode):convert golang struct into json format.
- Unmarshal(Decode): convert json into golang struct.

### NOTE-----> ONLY DATSTRUCTURES THAT CAN BE REPRESENTED AS **VALID JSON** will be encoded..

### Cyclical data structurea are not supported they will cause Marshal to go into an infinte loop.

### JSON only supports String as KEYS, to encode a go MAP type it must be in the form of "map[string]T"

### only the exported fields(those that begin with an uppercase letter) of the struct can be encoded in json.

### The POINTERS will be encoded as the values they are pointed to (or "null if the pointer in nil)

---

## Converting a "go type" into JSON----->(Marshal function)
