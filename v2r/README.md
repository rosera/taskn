# V2R 

A simple tool to write a Qwiklabs Yaml File 

## Usage
The application expects a folder to be provided as an argument.
`qwiklabs.yaml` file location is derived from the argument.

Ensure the path to the file is provided such that the application can
append the qwiklabs.yaml and locate the required file e.g.

```
[FOLDER PATH]/qwiklabs.yaml
```

#### Long Form

```bash
./v2r [Folder Path]
```

## Test

```bash
go test
```
