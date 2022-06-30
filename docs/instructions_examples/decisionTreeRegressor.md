
# Decision tree regressonr instruction fields

**name:** is an id for the model. This will allow you to recognize what model was trined with this instruction

**criterion:** The method used to split the tree. Supported values are "MSE" (Root mean square error) and RSS (Resudial sum of squares)

**minLeafSamples:** indicates how many samples should be in a leaf to generate a prediction

## Complete example

```json
[
  {
    "name": "model3",
    "kind": "regressor",
    "criterion": "RSS",
    "minLeafSamples": 20
  },
  {
    "name": "model4",
    "kind": "regressor",
    "criterion": "MSE",
    "minLeafSamples": 20
  }
]
```
