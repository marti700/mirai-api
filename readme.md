# Mirai-API: Simply train ML models

This API allows the training of multiple ML models at once by providing the training data and instructions on how to train the models.
Training data consists on the training data and the target variable that must be provided as csv files (for now) and
the model training instructions is provided as a json file that contains different configuration on how to train the model
for example if the following instruction is provided:
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

two linear regression models will be training one that uses Gradiant descent to estimete the model hyperparameters and one that uses the OLS close form solution.

for more info on how to set up an instruction file for different models you check docs/instructions_examples
for info on how to use the API checn docs/doc.md
