# TASKN

A simple tool to read a CSV file and perform a task on each entry read from the file.
The application assumes an `input.csv` file is present in the current directory.

## Usage
Add a command to perform on the filename read from the CSV.

#### Long Form

```bash
./taskn owner --input [FILENAME.csv]
```

#### Short Form
```bash
./taskn owner -i [FILENAME.csv]
```


## Test

```bash
go test
```
