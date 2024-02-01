# MLPH Attache DPWH Deployment CLI 

This tool helps automate the deployment of the DPWH web app through a command-line interface.

## Current Features

- Create backup of deployment files.

## To be Added
- Create backup of database
- Copy deployment files to the folder

## Configuration File

In the directory containing the CLI executable, create a `config.yml` file. This serves as the configuration file for the cli tool.

For deployment file backup:
```
backupMap:
  - source:      "/LOCATION/TO/FOLDER/frontend"
    destination: "/LOCATION/TO/DESTINATION/BackupFiles"
  - source:      "/LOCATION/TO/FOLDER/backend"
    destination: "/LOCATION/TO/DESTINATION/BackupFiles"
```
## Usage

To use the tool, the directory of the executable file can be added to the PATH environment variable.

Otherwise, go to the directory of executable file and execute:

On Linux/Mac
```
./dp -s "/location/to/config.yml" create-backup -f "Folder name"
```

On Windows
```
./dp.exe -s "/location/to/config.yml" create-backup -f "Folder name"
```

#### Flags

`\-s` is an optional flag that can be set to a custom location for the `config.yml` file. If the config file does not exist in this location, the tool will search for `./config.yml`. Otherwise, it will return an error.

`\-f` is a required flag used to set the folder name for the new backup files. The result after creating a backup for the deployment files will be:

```
/location/to/backupFiles/
    /folder_name_from_flag
        /file1
        /file2
```

#### Building from source

On Linux/Mac
```sh
go build -o dp
```

On Windows
```sh
go build -o dp.exe
```

