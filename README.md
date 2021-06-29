# CLI Task Manager

>Implementation of CLI Task Manager from **[Gophercises](https://courses.calhoun.io/courses/cor_gophercises)**  by Jon Calhoun, including the bonus section.

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

**Run Commands:**

- go install .
- go build -o tasks.exe .
- .\tasks `[flag]`

**Features:**

- building a command line application using cobra framework
- creating a central database to store and retrieve tasks
- performing some basic crud operations on boltdb

**Packages explored:**

- encoding/binary - to convert int types into byte stream
- time - set timeout on wait period to get database instance
- log - to simply log application generated errors
- strconv - to convert strings to their integer counterpart
- [github.com/mitchellh/go-homedir](github.com/mitchellh/go-homedir) - detect user's home directory
- [github.com/boltdb/bolt](https://github.com/boltdb/bolt) - database for project
- [github.com/spf13/cobra](https://github.com/spf13/cobra) - library for creating cli applications

**Output:**

``` terminal
PS D:\gophercises\task-manager> go build -o tasks.exe .
PS D:\gophercises\task-manager> .\tasks
Task is a CLI Task Manager

Usage:
  task [command]

Available Commands:
  add         Adds a task to your task list.
  do          Marks a task as completed.
  done        List all completed tasks
  help        Help about any command
  list        Lists all of your tasks.
  rm          Remove the task from list.
Flags:
  -h, --help   help for task

PS D:\gophercises\task-manager> .\tasks list
Nothing to see here, try adding some... 🙂

PS D:\gophercises\task-manager> .\tasks do
You have no tasks to complete, why not take a vacation? ⛵

PS D:\gophercises\task-manager> .\tasks done
You haven't completed any tasks yet, why not do them instead? 🙃

PS D:\gophercises\task-manager> .\tasks rm
You have no tasks to remove, start adding some... 📜

PS D:\gophercises\task-manager> .\tasks add try doing some exercise
Added 'try doing some exercise' to your task list.

PS D:\gophercises\task-manager> .\tasks add read some books        
Added 'read some books' to your task list.

PS D:\gophercises\task-manager> .\tasks add complete gophercises
Added 'complete gophercises' to your task list.

PS D:\gophercises\task-manager> .\tasks add solve cses problem set
Added 'solve cses problem set' to your task list.

PS D:\gophercises\task-manager> .\tasks add explore flutter cookbook
Added 'explore flutter cookbook' to your task list.

PS D:\gophercises\task-manager> .\tasks list
Pending tasks...
1. try doing some exercise
2. read some books
3. complete gophercises
4. solve cses problem set
5. explore flutter cookbook

PS D:\gophercises\task-manager> .\tasks do 1 one 1 two 2 2 10       
'one' is not a valid ID
'two' is not a valid ID
id: '10' out of bounds: '1-5'
You have completed the 'try doing some exercise' task.
You have completed the 'read some books' task.

PS D:\gophercises\task-manager> .\tasks done
Completed tasks...
1. try doing some exercise
2. read some books

PS D:\gophercises\task-manager> .\tasks list
Pending tasks...
1. complete gophercises
2. solve cses problem set
3. explore flutter cookbook

PS D:\gophercises\task-manager> .\tasks rm two 0 2   
'two' is not a valid ID
id: '0' out of bounds: '1-3'
Task 'solve cses problem set' has been deleted successfully.

PS D:\gophercises\task-manager> .\tasks list
Pending tasks...
1. complete gophercises
2. explore flutter cookbook
```
