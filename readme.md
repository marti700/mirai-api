# Mirai-API: Simply train ML models

This API allows the training of multiple ML models at once by providing the training data and instructions on how to train the models,
this "instructions" tells the api how the models are going to be trained, since a model can be trained in different ways an instruction file allows you to specify many training methods at once as a json array.

The model training instructions must be provided as a json file.

for example if the following instruction is provided to a Linear regression model via the /regression endpoint:

```json
[
  {
    "name": "first model",
    "estimators": {
      "GD": {
        "Iteations": 1000,
        "LearningRate": 0.01,
        "MinStepSize": 0.00003
      },
      "OLS": false
    },
    "regularization": {
      "type": "l1",
      "lambda": 20.0
    }
  },
  {
    "name": "third model",
    "estimators": {
      "OLS": true
    }
  }
]
```

two linear regression models will be trained one that uses Gradiant descent to estimete the model hyperparameters and one that uses the OLS close form solution. For more info on how to set up an instruction file for different models you check docs/instructions_examples

The data features (variables used to train a model) and the target variable (the variable we want to predict) must be provided in different files.

For info on how to use the API please check docs/doc.md

