# herald

Notify someone from some place about your success

## Install

```bash
go install github.com/fullpipe/herald@latest
```

## Usage

```bash
herald help
```

For example we want to notify grafana from local git repository

```bash
export GRAFANA_API_KEY=abirvalg
herald git grafana --host https://grafana.example.com:3000 --api-key ${GRAFANA_API_KEY} --tag TASK-123 --tag release --tag wtf-project
# or you could modify message
herald git grafana --host https://grafana.example.com:3000 --api-key ${GRAFANA_API_KEY} --tag TASK-123 --tag release --tag wtf-project -m "Project {{.project}} deployed. Going home."
```

| Add anotations                    | Feel the difference         | 
| ----------                        | :----------                 | 
| ![!](/assets/grafana-2.png)       | ![!](/assets/grafana-1.png) | 

## Places

Place tries to generate project metadata if possible.
You could find available place options by `herald PLACE help`

| Place      | Description                              | 
| ---------- | :----------                              | 
| nowhere    | set meta as you want                     | 
| git        | reads meta from local git repository     |
| gitlab     | reads meta from gitlab ci env vars       |

### Metadata

You could use metadata in message templates. For example:

```gotpl
Project {{.project}} deployed

Project: {{ or .project "none"}}
URL: {{ or .url "none"}}
Branch: {{ or .branch "none"}}
CommitSHA: {{ or .sha "none"}}
Author: {{ or .author "none"}}
```

## Targets

It's where you will see your success. 
You could find available target options by `herald nowhere TARGET --help`.

| Target        | Description                              | Required options       |
| ----------    | :----------                              | :-----------           |
| cli           | just prints out to your terminal         |                        |
| grafana       | creates grafana annotation with tags     | `--host`, `--api-key`  |
| slack         | send message to Slack channel            | `--webhook`            |

## TODO:

- ci
- wrap in docker container
- error handling
- lint
- commit message as meta.Description?
- more targets
  - telegram
  - ...
- more places
  - github actions
  - ...


