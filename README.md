# recsv filter
A command line utility that reads a CSV file specified in the argument or stdin, if unspecified, sorts the input lines and outputs the sorted line to a specified output file or stdout. This tool is useful for normalizing CSV output formats of different programs that generate CSV reports so that the outputs can easily be compared with a visual diff tool such as WinMerge.
## Usage
### Under Linux bash
`cat 1167_el_report_jun24.csv | recsv > 1167_el_report_jun24_sorted.csv`

or

`recsv 1167_el_report_jun24.csv 1167_el_report_jun24_sorted.csv`
### Under Window cmd
`type 1167_el_report_jun24.csv | recsv > 1167_el_report_jun24_sorted.csv`

or

`recsv 1167_el_report_jun24.csv 1167_el_report_jun24_sorted.csv`

