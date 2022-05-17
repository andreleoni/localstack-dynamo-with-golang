## Criar tabela no dynamodb

```
aws dynamodb --endpoint-url=http://localhost:4566 create-table \
    --table-name recommendations \
    --attribute-definitions \
        AttributeName=cluster_id,AttributeType=S \
        AttributeName=product_id,AttributeType=S \
    --key-schema \
        AttributeName=cluster_id,KeyType=HASH \
        AttributeName=product_id,KeyType=RANGE \
--provisioned-throughput \
        ReadCapacityUnits=10,WriteCapacityUnits=5
```

`aws --endpoint-url=http://localhost:4566 dynamodb describe-table --table-name recommendations | grep recommendations`

## Sobre partition key

https://aws.amazon.com/blogs/database/choosing-the-right-dynamodb-partition-key/


## Download nosequel workbench dynamodb

https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/workbench.settingup.html
