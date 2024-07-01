package evmUtils

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

/*
parseIntType Helper function for extracting XintYYY[] types from the input string. Returns bitsize (YYY), signed status (X) (if the signed flag is true
this means that they type is a signed type) and
whether the type is an array ([]) as a boolean. If returned bool is true this means that type is an array
*/
func parseIntType(input string) (bitSize int, isSigned bool, isArray bool, err error) {

	// Initialize the returned variables
	rint, rbool, rerr := 0, false, errors.New("")

	// Convert to lower case
	lowerCaseInput := strings.ToLower(input)

	index := strings.Index(lowerCaseInput, "int")

	if index == -1 {
		return rint, rbool, false, errors.New("Type is not int")
	}

	if index == 0 {
		rbool = true
	}

	// uint, int are considered 256 bytes
	if len(input) == index+3 {
		return 256, rbool, false, nil
	}

	endIndex := strings.Index(lowerCaseInput, "[")

	// if end index exists this means that the inputted type string is not an array
	if endIndex != -1 {
		rint, rerr = strconv.Atoi(input[index+3 : endIndex])
		return rint, rbool, true, rerr
	}

	rint, rerr = strconv.Atoi(input[index+3:])

	return rint, rbool, false, rerr
}

/*
parseByteType Helper function for extracting bytesXXX[] types from the input string. Returns bitsize (XXX) and
whether the type is an array ([]) as a boolean. If returned bool is true this means that type is an array
*/
func parseByteType(input string) (bitSize int, isArray bool, err error) {

	// Initialize the returned variables
	rint, rerr := 0, errors.New("")

	// Convert to lower case
	lowerCaseInput := strings.ToLower(input)

	index := strings.Index(lowerCaseInput, "bytes")

	if index == -1 {
		return rint, false, errors.New("type is not byte")
	}

	if len(input) == index+4 {
		return 0, false, errors.New("bitsize is not specified")
	}

	endIndex := strings.Index(lowerCaseInput, "[")

	// if end index exists this means that the inputted type string is not an array
	if endIndex != -1 {
		rint, rerr = strconv.Atoi(input[index+5 : endIndex])
		return rint, true, rerr
	}

	rint, rerr = strconv.Atoi(input[index+5:])

	return rint, false, rerr
}

/*
Helper function for extracting types inside the tuple string eg: tuple(int256, bytes, ...)
*/
func parseTuple(input string) (types []string, hasDynamicType bool, isArray bool, err error) {
	// Init
	types = []string{}

	// Dynamic Type Check
	containsBytes, err := regexp.MatchString("\\bbytes\\b", input)

	if err != nil {
		return
	}

	// Remove the tuple form tuple(...)
	inputCleaned := strings.Replace(input, "tuple", "", 1)

	// Check if this is a tuple[]
	if inputCleaned[0:2] == "[]" {
		isArray = true
	}

	hasDynamicType = strings.Contains(inputCleaned, "[") || strings.Contains(inputCleaned, "string") || containsBytes

	// Make it compatible with the formatter
	inputCleaned = "tuple" + inputCleaned

	formattedInput, err := tupleStringFormatter(inputCleaned)

	if err != nil {
		return
	}

	types, hasNestedTuple, err := tupleParseHelper(strings.Split(formattedInput, " "))

	if err != nil {
		return
	}

	hasDynamicType = hasDynamicType || hasNestedTuple

	return
}

func tupleStringFormatter(input string) (output string, err error) {
	// Format all the "," in the string
	commaFormatter, err := regexp.Compile("\\s*,\\s*")

	if err != nil {
		return
	}

	commaFormatted := commaFormatter.ReplaceAllString(input, " , ")

	// Format all the "(" and ")" in the string

	openningParantesesFormatter, err := regexp.Compile("\\s*\\(\\s*")

	if err != nil {
		return
	}

	closingParantesesFormatter, err := regexp.Compile("\\s*\\)\\s*")

	if err != nil {
		return
	}

	openningFormatted := openningParantesesFormatter.ReplaceAllString(commaFormatted, " ( ")
	closingFormatted := closingParantesesFormatter.ReplaceAllString(openningFormatted, " ) ")

	output = strings.Trim(closingFormatted, " ")

	return
}

func tupleParseHelper(input []string) (output []string, hasNestedTuple bool, err error) {
	currentLevel := -1
	lastWrittenIndex := 0

	for _, token := range input {
		// Understand starting of a nested tuple
		if token == "(" {
			currentLevel += 1

			if currentLevel >= 1 {
				hasNestedTuple = true
			}

			if currentLevel == 0 {
				continue
			}
		}

		// Understand ending of a nested tuple
		if token == ")" {
			currentLevel -= 1

			if currentLevel == -1 {
				continue
			} else {
				output[lastWrittenIndex] = output[lastWrittenIndex] + token
				continue
			}
		}

		if currentLevel == 0 {
			if token != "," && token != "" {
				output = append(output, token)
				lastWrittenIndex = len(output) - 1
				continue
			}
			continue
		}

		if currentLevel > 0 {
			output[lastWrittenIndex] = output[lastWrittenIndex] + token
			continue
		}

	}

	return
}
