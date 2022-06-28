# Available endpoints

## /regression

 Trains a linear regression model this endpoint returns an array of json models trained based on the provided instructions. To see the posible instruction combinations check the linReg.md file in the instructions_example folder

 request example:
 curl -F 'json=@path/to/instruction/file.json' -F 'train=@path/to/train/data.csv' -F 'target=@path/to/target/data.csv' http://localhost:9090/regression

 curl -F 'json=@./parser/instruction/linReg.json' -F 'train=@./reqhandler/benchmarkdata/x_train.csv' -F 'target=@./reqhandler/benchmarkdata/y_train.csv' http://localhost:9090/regression

 response example:

 ```json
[
  {
    "ModelName": "third model",
    "Model": {
      "Hyperparameters": {
        "Row": 11,
        "Col": 1,
        "Data": [
          0.46489296798558666,
          25.930958811794405,
          85.16096491147886,
          67.37470375925801,
          83.11931230630952,
          60.03642725583279,
          52.74548458336103,
          65.52379401875451,
          98.13034914861448,
          42.04680826386831,
          42.5684777408541
        ]
      },
      "Opts": {
        "Estimator": {},
        "Regularization": {
          "Type": "",
          "Lambda": 0
        }
      }
    }
  },
  {
    "ModelName": "first model",
    "Model": {
      "Hyperparameters": {
        "Row": 11,
        "Col": 1,
        "Data": [
          2.196122825441585e+110,
          6.141375620505818e+109,
          -6.382092921576376e+109,
          -2.4761571161551897e+110,
          -1.3269077812627159e+110,
          1.5543050920674517e+110,
          -4.08397765944916e+109,
          -1.2949826003609184e+110,
          -6.5537604711919086e+109,
          6.797939928786705e+109,
          -1.0692408654206326e+110
        ]
      },
      "Opts": {
        "Estimator": {
          "Iteations": 1000,
          "LearningRate": 0.01,
          "MinStepSize": 3e-05
        },
        "Regularization": {
          "Type": "l1",
          "Lambda": 20
        }
      }
    }
  },
  {
    "ModelName": "second model",
    "Model": {
      "Hyperparameters": {
        "Row": 11,
        "Col": 1,
        "Data": [
          2.5358403045823112e+110,
          7.091382887897631e+109,
          -7.369336664887424e+109,
          -2.8591930027238152e+110,
          -1.5321666863115007e+110,
          1.79474000835524e+110,
          -4.715726748918294e+109,
          -1.4953030102347397e+110,
          -7.56755940828814e+109,
          7.849510901595641e+109,
          -1.2346413645124386e+110
        ]
      },
      "Opts": {
        "Estimator": {
          "Iteations": 100,
          "LearningRate": 0.01,
          "MinStepSize": 0.002
        },
        "Regularization": {
          "Type": "l2",
          "Lambda": 0.01
        }
      }
    }
  }
]
```

## /decisiontree/regression

Trains a dicision tree model for regression. Since tree model can be hard to read as json a .dot files are downloaded instead as representation of the models

request example:

curl -F 'json=@path/to/instruction/file.json' -F 'train=@path/to/train/data.csv' -F 'target=@path/to/target/data.csv' http://localhost:9090/decisiontree/regression > aquielacosar.zip

curl -F 'json=@./parser/instruction/decisionTreeRegressor.json' -F 'train=@./reqhandler/benchmarkdata/x_train.csv' -F 'target=@./reqhandler/benchmarkdata/y_train.csv' http://localhost:9090/decisiontree/regression > aquielacosar.zip

## /decisiontree/classification

Trains a dicision tree model for classification. Since tree model can be hard to read as json a zip of .dot files are downloaded instead as representation of the models

request example:

curl -F 'json=@path/to/instruction/file.json' -F 'train=@path/to/train/data.csv' -F 'target=@path/to/target/data.csv' http://localhost:9090/decisiontree/regression > aquielacosar.zip

curl -F 'json=@./parser/instruction/decisionTreeClassifier.json' -F 'train=@./reqhandler/benchmarkdata/x_train.csv' -F 'target=@./reqhandler/benchmarkdata/y_train.csv' http://localhost:9090/decisiontree/classification > aquielacosac.zip