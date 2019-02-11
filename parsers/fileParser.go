package parsers

import (
	"bufio"
	"os"
	"path"
	"strings"

	_ "github.com/AntoineAugusti/moduluschecking/data"
	"github.com/AntoineAugusti/moduluschecking/helpers"
	m "github.com/AntoineAugusti/moduluschecking/models"
)

// Describes the content of a file.
type LineRecord struct {
	// The content of a line in the file
	content string
	// The line number in the file
	lineNumber int
}

// A parser that reads data from the filesystem.
type FileParser struct {
	// The path to the weights file.
	weightsPath string
	// The path to the substitutions file for sort codes.
	substitutionsPath string
	// The actual weights for each sort code.
	weights map[string]m.SortCodeData
}

// Get all known sort code substitutions.
func (fp FileParser) Substitutions() map[string]string {
	substitutions := make(map[string]string)

	jobs := make(chan LineRecord)
	go readFile(fp.substitutionsPath, jobs)

	for lineRecord := range jobs {
		fields := strings.Split(lineRecord.content, " ")
		key, value := fields[0], fields[1]
		substitutions[key] = value
	}

	return substitutions
}

// Get the weights, exception information and algorithm to use for all known sort codes.
func (fp FileParser) Weights() map[string]m.SortCodeData {
	jobs := make(chan LineRecord)
	results := make(chan m.SortCodeRange)

	go readFile(fp.weightsPath, jobs)
	go parseWeightsLine(jobs, results)

	// Process all the sort code ranges
	for result := range results {
		fp.addSortCodeRange(result)
	}

	return fp.weights
}

// Process a sort code range and add it to the weights map.
func (fp *FileParser) addSortCodeRange(scRange m.SortCodeRange) {
	scData := m.SortCodeData{
		Algorithm:      scRange.Algorithm,
		Weights:        scRange.Weights,
		ExceptionValue: scRange.ExceptionValue,
		Next:           nil,
		LineNumber:     scRange.LineNumber,
	}

	// Go over the sort code range and add each sort code in the range in
	// a map to decrease lookup time later. This requires a larger amount of
	// memory, but it seems worth it.
	for sortCode := scRange.Start; sortCode <= scRange.End; sortCode++ {
		key := helpers.AddLeadingZerosToNumber(sortCode)
		val, exist := fp.weights[key]
		if !exist {
			fp.weights[key] = scData
			continue
		}

		// Check that the first data structure was before in the weights file
		if val.LineNumber < scData.LineNumber {
			var tmp = val
			tmp.Next = &scData
			fp.weights[key] = tmp
		} else {
			// We read the second weights first.
			// Put it in the right order in the map
			var tmp = val
			scData.Next = &tmp
			fp.weights[key] = scData
		}
	}
}

// Parse lines from the weights file and put the result
// as a SortCodeRange structure in a channel.
func parseWeightsLine(jobs <-chan LineRecord, results chan<- m.SortCodeRange) {
	var fields []string

	for lineRecord := range jobs {
		lineNumber, data := lineRecord.lineNumber, lineRecord.content
		fields = strings.Split(data, ",")
		// Sort code range
		sortCodeStart, sortCodeEnd := helpers.ToInt(fields[0]), helpers.ToInt(fields[1])
		// Algorithm to use in order to perform the check
		algorithm := fields[2]
		// Weights for sort code and account number
		weights := fields[3:17]

		scRange := m.SortCodeRange{
			Start:          sortCodeStart,
			End:            sortCodeEnd,
			Algorithm:      algorithm,
			Weights:        helpers.StringSliceToIntSlice(weights),
			ExceptionValue: 0,
			LineNumber:     lineNumber,
		}

		// Does this sort code range has got an exception?
		hasException := len(fields) > (2 + 1 + 14)

		// Set the exception value if needed
		if hasException {
			scRange.ExceptionValue = helpers.ToInt(fields[17])
		}

		results <- scRange
	}

	close(results)
}

// Read a file and put the content in a channel.
func readFile(path string, jobs chan<- LineRecord) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	lineNumber := 1

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		jobs <- LineRecord{
			content:    scanner.Text(),
			lineNumber: lineNumber,
		}
		lineNumber += 1
	}

	// We are done with the file, release the channel
	close(jobs)
}

// Create a new instance of a file parser that satisfies
// the parser interface.
func CreateFileParser() m.Parser {
	goPath := path.Clean(os.Getenv("GOPATH"))
	dataPath := goPath + "/src/github.com/AntoineAugusti/moduluschecking/data/"

	return FileParser{
		weightsPath:       dataPath + "weights.txt",
		substitutionsPath: dataPath + "substitutions.txt",
		weights:           make(map[string]m.SortCodeData),
	}
}
