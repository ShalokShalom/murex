package docs

var Digest map[string]string = map[string]string{
	`trypipe`: `Checks state of each function in a pipeline and exits block on error`,
	`append`: `Add data to the end of an array`,
	`getfile`: `Makes a standard HTTP request and return the contents as _murex_-aware data type for passing along _murex_ pipelines.`,
	`err`: `'echo' a string to the STDERR`,
	`out`: `'echo' a string to the STDOUT`,
	`>`: `Writes STDIN to disk - overwriting contents if file already exists`,
	`f`: `Lists objects (eg files) in the current working directory`,
	`event`: `Event driven programming for shell scripts`,
	`print`: `Write a string to the OS STDOUT (bypassing _murex_ pipelines)`,
	`g`: `Glob pattern matching for file system objects (eg *.txt)`,
	`rx`: `Regexp pattern matching for file system objects (eg '.*\.txt')`,
	`try`: `Handles errors inside a block of code`,
	`murex-docs`: `Displays the man pages for _murex_ builtins`,
	`tout`: `'echo' a string to the STDOUT and set it's data-type`,
	`if`: `Conditional statement to execute different blocks of code depending on the result of the condition`,
	`alter`: `Change a value within a structured data-type and pass that change along the pipeline without altering the original source input`,
	`get`: `Makes a standard HTTP request and returns the result as a JSON object`,
	`post`: `HTTP POST request with a JSON-parsable return`,
	`set`: `Define a variable and set it's value`,
	`>>`: `Writes STDIN to disk - appending contents if file already exists`,
	`pt`: `Pipe telemetry. Writes data-types and bytes written`,
	`ttyfd`: `Returns the TTY device of the parent.`,
	`prepend`: `Add data to the start of an array`,
	`swivel-datatype`: `Converts tabulated data into a map of values for serialised data-types such as JSON and YAML`,
	`swivel-table`: `Rotates a table by 90 degrees`,
	`catch`: `Handles the exception code raised by 'try' or 'trypipe'`,
	`unset`: `Deallocates an environmental variable (aliased to '!export')`,
}