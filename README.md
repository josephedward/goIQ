# AWS IQ Automation

This project helps automate the process of interacting with AWS IQ requests.

**Local Start:**

```bash
task devtools
task auth
task cli -- <DEVTOOLS-STRING-FROM-STEP-1>
```

This project uses a [taskfile](https://taskfile.dev/) to run tasks:

- `devtools`:  
  launches a headless browser with a custom user agent

- `auth`:  
  authenticates the browser console with the Amazon Web Services

- `cli`:
  runs the main script for the command line interface

- `scrape`:
  tests use of Rod library for scraping

- `launcher`
  tests use of Rod library for launching a browser

**X not implemented yet X** <br>
`test`:  
runs all integration and unit tests
- zap vs zerolog 