# Cron trigger
This trigger provides your flogo application the ability to start a flow via a timer configured by a cron expression.

It uses the library https://github.com/robfig/cron

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
| expression  | True     | The cron expression.  |

## Cron Expression

A cron expression represents a set of times, using 6 space-separated fields.

| Field name   | Mandatory? | Allowed values  | Allowed special characters |
|--------------|:-----------|:----------------|:---------------------------|
| Seconds      | Yes        | 0-59            | * / , -                    |
| Minutes      | Yes        | 0-59            | * / , -                    |
|	Hours        | Yes        | 0-23            | * / , -                    |
|	Day of month | Yes        | 1-31            | * / , - ?                  |
|	Month        | Yes        | 1-12 or JAN-DEC | * / , -                    |
|	Day of week  | Yes        | 0-6 or SUN-SAT  | * / , - ?                  |

See the full documentation : https://godoc.org/github.com/robfig/cron

## Example

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
