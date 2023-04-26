package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/gocolly/colly/v2"
)

var (
	converter = md.NewConverter("", true, nil)
)

type LeetcodeGraphQlResponse struct {
	Data struct {
		Question struct {
			ExampleTestcaseList []string `json:"exampleTestcaseList"`
			Content             string   `json:"content"`
			MysqlSchemas        []any    `json:"mysqlSchemas"`
			QuestionTitle       string
		} `json:"question"`
	} `json:"data"`
}

func main() {
	problems := getProblems()

	for i := range problems {
		src := problems[i]
		xs := strings.Split(src, "/")
		slug := xs[len(xs)-1]

		path := fmt.Sprintf("../../%d-%s/", i+1, slug)
		fmt.Println(path)

		response := queryLeetcode(slug)
		createFile(path, response.Data.Question.Content)
		time.Sleep(5 * time.Second)
	}
}

// collects problem names from grind75 list at techinterviewhandbook.org.
func getProblems() (problems []string) {
	c := colly.NewCollector()

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		src := e.Attr("href")
		if strings.HasPrefix(src, "https://leetcode.com/problems/") {
			problems = append(problems, src)
		}
	})

	c.Visit("https://www.techinterviewhandbook.org/grind75")

	return problems
}

// creates a file if it does not exist at path with following content. Intermediate folders are created if they dont exist.
func createFile(path string, content string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(path, 0700) // Create your file
		// PROBLEM.md
		if markdown, err := converter.ConvertString(content); err == nil {
			f, err := os.Create(path + "PROBLEM.md")

			if err != nil {
				log.Fatal(err)
			}

			defer f.Close()

			_, err2 := f.WriteString(markdown)

			if err2 != nil {
				log.Fatal(err2)
			}
		}
	}

}

// query leetcode graphql API for questioncontents.
func queryLeetcode(slug string) LeetcodeGraphQlResponse {
	res, err := http.Post("https://leetcode.com/graphql/", "application/json;charset=UTF-8", makeQuestionContent(slug))
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	response := LeetcodeGraphQlResponse{}
	json.Unmarshal(body, &response)
	return response
}

// returns a reader in leetcode graphql API query format for specific titleSlug.
func makeQuestionContent(titleSlug string) io.Reader {
	questionContent := `{"query":"\n    query questionContent($titleSlug: String!) {\n  question(titleSlug: $titleSlug) {\n    content\n    mysqlSchemas\n  }\n}\n    ","variables":{"titleSlug":"<replace>"},"operationName":"questionContent"}`
	replaced := strings.Replace(questionContent, "<replace>", titleSlug, 1)
	return strings.NewReader(replaced)
}
