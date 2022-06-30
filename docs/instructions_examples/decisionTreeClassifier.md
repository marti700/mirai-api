# Decision tree regressonr instruction fields

**name:** is an id for the model. This will allow you to recognize what model was trined with this instruction

**criterion:** The method used to split the tree. Supported values are "GINI" and "ENTROPY"

## Complete example

```json
[
  {
    "name": "model1",
    "kind": "classifier",
    "criterion": "GINI"
  },
  {
    "name": "model2",
    "kind": "classifier",
    "criterion": "ENTROPY"
  }
]
```
