# Linear regression model instructions

Instructions tells the api how the models are going to be trained. There are many methods to estimate a model hyperparameters
the instruction file allows you to specify many training methods at once as an array of intruction in a json file

## Linear regression instruction fields

name: is an id for the model. This will allow you to recognize what model was trined with this instruction
estimators: specify what estimator should be used to train the linear regression model. Supported options are:

- GD for gradient descent this one must be specified as a json object (with "GD key") with the following properties:
    Iterations: the max number of iteration the GD loop will make
    LearningRate: the model learning rate
    MinStepSize: the size the Gradient descent steps should be before finishing the estimation. E.X.: In the below example if the step size is 0.00003 the GD loop will end\
    EX:

```json
    "GD": {
      "Iteations": 1000,
      "LearningRate": 0.001,
      "MinStepSize": 0.00003
    }
```

- OLS: If set to true estimates the linear model hyperparameters from the Ordinary least squares close form solution
 regularization: Specifies if regularization should be applyed to the model. Value should be a javascript object with keys type and lambda

  EX:

```json
   "regularization": {
      "type": "l1",
      "lambda": 0.01
      }
```

supported types are "l1" for lasso regression and "l2" for ridge regression\
the lambda value must be a float value

## Example

```json
[
  {
    "name": "first model",
    "estimators": {
      "GD": {
        "Iteations": 1000,
        "LearningRate": 0.001,
        "MinStepSize": 0.00003
      },
      "OLS": true
    },
    "regularization": {
      "type": "l1",
      "lambda": 0.01
    }
  },
  {
    "name": "second model",
    "estimators": {
      "GD": {
        "Iteations": 100,
        "LearningRate": 0.01,
        "MinStepSize": 0.002
      },
      "OLS": false
    },
    "regularization": {
      "type": "l2",
      "lambda": 0.01
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
