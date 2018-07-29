package docs

// Digest stores a 1 line summary of each builtins
var Digest map[string]string = map[string]string{
	`swivel-datatype`: `Converts tabulated data into a map of values for serialised data-types such as JSON and YAML`,
	`getfile`:         `Makes a standard HTTP request and return the contents as _murex_-aware data type for passing along _murex_ pipelines.`,
	`>>`:              `Writes STDIN to disk - appending contents if file already exists`,
	`and`:             `Returns 'true' or 'false' depending on whether multiple conditions are met`,
	`export`:          `Define an environmental variable and set it's value`,
	`global`:          `Define a global variable and set it's value`,
	`alter`:           `Change a value within a structured data-type and pass that change along the pipeline without altering the original source input`,
	`>`:               `Writes STDIN to disk - overwriting contents if file already exists`,
	`f`:               `Lists objects (eg files) in the current working directory`,
	`rx`:              `Regexp pattern matching for file system objects (eg '.*\.txt')`,
	`trypipe`:         `Checks state of each function in a pipeline and exits block on error`,
	`append`:          `Add data to the end of an array`,
	`brace-quote`:     `Write a string to the STDOUT without new line`,
	`out`:             `'echo' a string to the STDOUT with a trailing new line character`,
	`ttyfd`:           `Returns the TTY device of the parent.`,
	`read`:            `'read' a line of input from the user and store as a variable`,
	`catch`:           `Handles the exception code raised by 'try' or 'trypipe'`,
	`event`:           `Event driven programming for shell scripts`,
	`prepend`:         `Add data to the start of an array`,
	`post`:            `HTTP POST request with a JSON-parsable return`,
	`if`:              `Conditional statement to execute different blocks of code depending on the result of the condition`,
	`try`:             `Handles errors inside a block of code`,
	`tread`:           `'read' a line of input from the user and store as a user defined typed variable`,
	`murex-docs`:      `Displays the man pages for _murex_ builtins`,
	`get`:             `Makes a standard HTTP request and returns the result as a JSON object`,
	`tout`:            `Print a string to the STDOUT and set it's data-type`,
	`g`:               `Glob pattern matching for file system objects (eg *.txt)`,
	`set`:             `Define a local variable and set it's value`,
	`swivel-table`:    `Rotates a table by 90 degrees`,
	`err`:             `Print a line to the STDERR`,
	`or`:              `Returns 'true' or 'false' depending on whether one code-block out of multiple ones supplied is successful or unsuccessful.`,
	`pt`:              `Pipe telemetry. Writes data-types and bytes written`,
}
