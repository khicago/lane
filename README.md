# lane

A swim lane implementation.

# Usage

```go
import "github.com/khicago/lane"

const LaneType = "test"
var LTest = lane.New(LaneType)


```

## Lane with value

```go
import "github.com/khicago/lane"

const LaneType = "test"
var LTestWithID = lane.New(LaneType, "this_is_a_test_id")

// get
func main() {
    // ...
    id := lane.Value.ReadFrom(LTestWithID)   
}
```

## Lane tags

```go
import "github.com/khicago/lane"

const LaneType = "test"
const TagTestEnv lane.Tag = "env"
const TagTestDesc lane.Tag = "desc"

var LTest1 = lane.New(LaneType, "test1").Apply(TagTestEnv, "development")
var LTest2 = lane.New(LaneType, "test2").Apply(TagTestEnv, "development")

func init() {
    TagTestDesc.WriteTo(LTest1, "desc of test1")
    TagTestDesc.WriteTo(LTest2, "desc of test1")
}


// get
func main() {
    laneCur := LTest1 // or LTest2

    // ...
    id := lane.Value.ReadFrom(LTestWithID)  
    env := TagTestEnv.ReadFrom(laneCur)
    desc := TagTestDesc.ReadFrom(laneCur)
}
```

## Payload methods

```go
import "github.com/khicago/lane"

// ...

func main() {
    laneCur := LTest1 // or LTest2

    targetMap := map[string]string {
        "development": "channel 1",
        "production": "channel 2",
    }

    // ...
    envPayload := TagTestEnv.Of(laneCur)
    if ! envPayload.Is("development") {
        panic("must running in dev environment")
    }
    targetChannel := envPayload.Select(targetMap)
    // ...
}
```