# CS8803-O10-final

This repository is used to transform the dataset used in my CS8803-O10 Summer
2020 project, the 2018 Federal Housing Finance Agency (FHFA) Single-Family
Mortgage-Level Owner-Occupied 1-Unit Properties (.zip file can be downloaded
[here](https://www.fhfa.gov/DataTools/Downloads/Pages/Single-Family-Mortgage-Level-Owner-Occupied-1-Unit-Property-(National-File-A).aspx)).

## Dependencies

In order to run this script, you must have Docker installed.

## Usage

To run this script simply run `cleanup.sh`. This will do the following:

1. Download the zip file containing the raw data
1. Unzip the zipped file
1. Run a Golang script in Docker which transforms the raw data into a CSV file
1. Produce a CSV file called `./fhlmc_sf2018a_loans.csv`
