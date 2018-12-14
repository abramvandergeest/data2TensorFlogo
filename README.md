---
title: data2TensorFlow
weight: 4717
---

# data2Tensor
This activity converts data to a tensorflow tensor

#below still the data from the copied source
## Installation
### Flogo Web
This activity comes out of the box with the Flogo Web UI
### Flogo CLI
```bash
flogo install github.com/TIBCOSoftware/flogo-contrib/activity/mapper
```

## Metadata
```json
{
  "input":[
    {
      "name": "mappings",
      "type": "array",
      "required": true,
      "display": {
        "name": "Mapper",
        "type": "mapper",
        "mapperOutputScope" : "action"
      }
    }
  ]
}
```
### Details
#### Settings:
| Setting     | Required | Description |
|:------------|:---------|:------------|
| mappings    | true     | An array of mappings that are executed when the activity runs |

## Example
The below example allows you to configure the activity to reply and set the output values to literals "name" (a string) and 2 (an integer).

```json
{
  "id": "mapper_6",
  "name": "Mapper",
  "description": "Simple Mapper Activity",
  "activity": {
    "ref": "github.com/TIBCOSoftware/flogo-contrib/activity/mapper",
    "input": {
      "mappings": [
        {
          "mapTo": "FlowAttr1",
          "type": "assign",
          "value": "$activity[log_3].message"
        }
      ]
    }
  }
}
```
