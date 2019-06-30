# slack-notify
cli to send argo workflow notification to a slack channel

### Command line
```
Usage:
  slack-notify [flags]

Flags:
  -t, --api-token string     slack api token
  -c, --channel string       slack channel where the message will be posted
      --doc                  generate markdown documentation
  -e, --endpoint string      argo-ui http endpoint that will be used for visualizing workflow output (required)!!!!
  -h, --help                 help for slack-notify
      --log-file string      file for log output other than standard output, written to a temp folder by default
      --log-format string    format of the logging out, either of json or text (default "json")
      --log-level string     log level for the application (default "error")
  -s, --status string        the status of workflow, will be used to pick up an emoji for displaying in slack
  -w, --workflow-id string   a workflow identifier that has triggered this slack notification
```

required flag(s) "api-token", "channel", "endpoint", "status", "workflow-id" 
