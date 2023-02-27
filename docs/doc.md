# The Instruction File
 The instruction file is a json file uploaded to the API. It contains the instructions that will be used to train the models. it has information such as the type of model to be trained (regression, decisiontree, etc) and the hyperparameters that will be used to train these models. The fnstruction file options are the following:

   **InstructionType:** the type of model we are going to train it must be one of:
    - linearregression" 
    - decisiontreeclassifier
    - decisiontreeregressor"

   **name:** is the id of the InstructionType. it helps identify to What instructionType a model belongs
   **instructions:** is a Json array that contains a list of models with its training parameters, available parameters are given by the InstructionType
     For linear regression
         **name:** An id for the model, allows its unique identification after training \
         **estimators:** the estimator used to estimate the linear regression coeficients, it have to be one of: \
           *- GD* (for gradicent descent): when this option is speficied it's also required to spedify as a json object the the following properties \
             *- Iterations:* the number of itererations performed by the gradiend descent \
             *- LearningRate:* the learning rate of the GD \
             *- MinStepSize:* how low should the "step size" be before exiting the GD loop \
           *- OLS:* a boolean value indicating if the linear regression close form solution should be used, is set to true any other estimator will be ignored \
         **regularization:** specifies tye type of regularization used, supported types are l1 and l2 regularization and are specified as fallows: \
           *- type:* must by either l1 or l2 \
           *- lambda:* spefifies the regularization rate 
   
   For decision tree classifier \
       **name:** An id for the model, allows its unique identification after training \
       **kind:** the type of desicion tree supported types are "classifier" or "regressor" \
       **criterion:** the criteria on which the tree is trained on.\
         For a tree of kind classifier the supported criterions are *"GINI"* and *"ENTROPY"*\
         For a tree of kind regressor the supported criterions are "RSS" and "MSE"\
       **minLeafSamples:** The minimun amount of leafs the tree should have


### Example of a struction file:

```json
[
  {
    "InstructionType": "linearregression",
    "name": "An id for this instruction set",
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

# Available endpoints

## /train

 Trains the models specified in the instruction file, the data used for training is passed in the train request parameter. the target variable
 must be passed to the API using the target request parameter. Since training can take a while an email must be specified to recieve a report
 containing matrics on how the model performed on the test data.

 **request example:**\
 curl -F 'json=@path/to/instruction/file.json' -F 'trainData=@path/to/train/data.csv' -F 'trainTarget=@path/to/target/data.csv' -F 'testData=@path/to/test/data.csv' -F 'testTarget=@path/to/test/target/data.csv'  <http://localhost:9090/train?email=mail@email.com>

 **response example:**

 *OK*\
 ```json
 {"Status": "OK", "ErrorMessage: ''"}
 ```

 *Fail*\
 ```json
 {"Status": "Fail", "ErrorMessage": 'Failure reason'}
 ```

