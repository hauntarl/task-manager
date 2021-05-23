# CLI Task Manager

 Building a CLI tool that can be used to manage your TODOs in the terminal. The basic usage of the tool is going to look roughly like this:

 ``` terminal
$ task
task is a CLI for managing your TODOs.

Usage:
  task [command]

Available Commands:
  add         Add a new task to your TODO list
  do          Mark a task on your TODO list as complete
  list        List all of your incomplete tasks

Use "task [command] --help" for more information about a command.

$ task add review talk proposal
Added "review talk proposal" to your task list.

$ task add clean dishes
Added "clean dishes" to your task list.

$ task list
You have the following tasks:
1. review talk proposal
2. some task description

$ task do 1
You have completed the "review talk proposal" task.

$ task list
You have the following tasks:
1. some task description
 ```

- ```add``` - adds a new task to our list
- ```list``` - lists all of our incomplete tasks
- ```do``` - marks a task as complete

Bonus:

``` terminal
$ task rm 1
You have deleted the "review talk proposal" task.

$ task completed
You have finished the following tasks today:
- wash the dishes
- clean the car
```

>Implementation of CLI Task Manager from **[Gophercises](https://courses.calhoun.io/courses/cor_gophercises)**  by Jon Calhoun, including the bonus section.

<!--
**Run Commands:**

- go run main.go --help (-h)
- go run main.go
- go run main.go --depth ```int``` (-depth=```int```)

**Features:**

- performing http.Get request for given urls and parsing the documents
- using bread-first traversal on child hrefs to generate a sitemap
- encoding generated sitemap into xml format

**Packages explored:**

- encoding/xml - to encode go data structure into xml format
- flag - to get depth of search and root url
- net/http - to perform GET request on urls
- net/url - to access specific parts of a request url
- os - to create new file and store encoded results into
- strings - to check for prefixes in the parsed links
- [github.com/hauntarl/link-parser](https://github.com/hauntarl/link-parser) - to parse the HTML document and extract all hrefs from it

**Output:**

``` terminal
```
