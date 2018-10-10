# Cron trigger
This trigger provides your flogo application the ability to start a flow via a timer configured by a cron expression.

## Installation
### Flogo Web

#### Start

Start a container of Flogo Web UI :

```bash
docker run --name flogo -it -d -p 3303:3303 -e FLOGO_NO_ENGINE_RECREATION=false flogo/flogo-docker eula-accept
```
*The environment variable FLOGO_NO_ENGINE_RECREATION=false allows to force import of installed contributions.*

#### Installation of the activity

To install the activity into the started container :

```bash
docker exec -it flogo sh -c 'cd /tmp/flogo-web/build/server/local/engines/flogo-web && flogo install github.com/square-it/flogo-contrib-triggers/cron'
```

Restart the container
```bash
docker restart flogo
```

### Flogo CLI
```bash
flogo install github.com/square-it/flogo-contrib-triggers/cron
```

## Schema
Inputs and Outputs:

```json
{
  "settings": [
    {
      "name": "expression",
      "type": "string"
    }
  ]
}
```
## Settings
| Setting     | Required | Description |
|:------------|:---------|:------------|
| expression  | True     | The cron expression |

## Examples

```json
{
      "id": "cron",
      "ref": "github.com/square-it/flogo-contrib-triggers/cron",
      "name": "Cron",
      "description": "Timer trigger with cron abilities",
      "settings": {
        "expression": "0/10 * * * * *"
      }
}

```

