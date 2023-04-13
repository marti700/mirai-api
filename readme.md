# Mirai-API: Simply train ML models

This API allows the training of multiple ML models at once by providing the training data and instructions on how to train the models,
this "instructions" tells the api how the models are going to be trained, since a model can be trained in different ways an instruction file allows you to specify many training methods at once using a json array.

The model training instructions must be provided as a json file.

E.X

for example the following instruction file asks the api to train three different models via the `InstructionType key`, a Linear regression, a decicion tree classifier and a decision tree regressor. The instruction key then specifies the different models to be trained, for example the "first model" is being trained using gradiant descent, with 1000 interactions and a learning rate of 0.001. and l1 regularization.

```json
[
  {
    "InstructionType": "linearregression",
    "name": "first Instruction",
    "instructions": [
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
  },
  {
    "InstructionType": "decisiontreeclassifier",
    "name": "second Instruction DTC",
    "instructions": [
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
  },
  {
    "InstructionType": "decisiontreeregressor",
    "name": "third Instruction",
    "instructions": [
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
  }
]
```

The data features (variables used to train a model) and the target variable (the variable we want to predict) must be provided in different files.

### Test it with docker

*run the container:*\
```
docker run -e "SENDER_EMAIL=teodoro641@hotmail.com"\
  -e "SMTP_PORT=587" -e "SMTP=smtp.office365.com"\
  -e "SENDER_EMAIL_PASSWORD=BtplSgMb50qK"\
  -p 9090:9090 mirai-api:1.0.1
```

*make a request:*\
```
curl -F 'json=@./exampledata/all.json'\
 -F 'trainData=@./exampledata/LinearReg/x_train.csv'\
 -F 'trainTarget=@./exampledata/LinearReg/y_train.csv'\
 -F 'testData=@./exampledata/LinearReg/x_test.csv'\
 -F 'testTarget=@./exampledata/LinearReg/y_test.csv'\
 http://localhost:9090/train?email=teodoro641@gmail.com

```


For info on how to use the API please check docs/doc.md

