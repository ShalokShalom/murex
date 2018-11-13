package docs

//Synonym is used for builtins that might have more than one internal alias
var Synonym map[string]string = map[string]string{
	`(`:               `brace-quote`,
	`echo`:            `out`,
	`!and`:            `and`,
	`!or`:             `or`,
	`!catch`:          `catch`,
	`!global`:         `global`,
	`!set`:            `set`,
	`!if`:             `if`,
	`!export`:         `export`,
	`unset`:           `export`,
	`!event`:          `event`,
	`if`:              `if`,
	`trypipe`:         `trypipe`,
	`post`:            `post`,
	`tout`:            `tout`,
	`f`:               `f`,
	`read`:            `read`,
	`prepend`:         `prepend`,
	`murex-docs`:      `murex-docs`,
	`out`:             `out`,
	`append`:          `append`,
	`>>`:              `>>`,
	`ttyfd`:           `ttyfd`,
	`g`:               `g`,
	`event`:           `event`,
	`get`:             `get`,
	`getfile`:         `getfile`,
	`pt`:              `pt`,
	`global`:          `global`,
	`alter`:           `alter`,
	`swivel-table`:    `swivel-table`,
	`err`:             `err`,
	`tread`:           `tread`,
	`brace-quote`:     `brace-quote`,
	`catch`:           `catch`,
	`swivel-datatype`: `swivel-datatype`,
	`>`:               `>`,
	`and`:             `and`,
	`set`:             `set`,
	`rx`:              `rx`,
	`or`:              `or`,
	`try`:             `try`,
	`export`:          `export`,
}
