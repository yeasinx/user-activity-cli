# GitHub User Activity CLI

A simple command-line tool to fetch and display a GitHub user's recent activity.

## Usage

```bash
go run main.go <username>
```

Or build and run:

```bash
go build -o github-activity
./github-activity <username>
```

## Example

```bash
$ go run main.go torvalds
- Pushed 1 commits to torvalds/linux
- Pushed 2 commits to torvalds/GuitarPedal
- Starred octocat/Hello-World
...
```

## Supported Events

| Event Type    | Output                                      |
|---------------|---------------------------------------------|
| PushEvent     | `Pushed X commits to repo/name`             |
| IssuesEvent   | `Opened a new issue in repo/name`           |
| WatchEvent    | `Starred repo/name`                         |
| Other events  | `EventType in repo/name`                    |

## Corner Cases Handled

| Scenario                      | Behavior                                          |
|-------------------------------|---------------------------------------------------|
| No username provided          | Displays usage message and exits                  |
| User not found (404)          | Displays "User X was not found!"                  |
| API rate limit exceeded (403) | Displays rate limit error message                 |
| Network/HTTP error            | Displays error and exits                          |
| Invalid JSON response         | Logs fatal error                                  |
| No recent activity            | Displays "No recent activity found for this user" |
| Unexpected status code        | Displays status code and message                  |

## API

This tool uses the GitHub Events API:

```
https://api.github.com/users/{username}/events
```

> **Note:** Unauthenticated requests are limited to 60 requests per hour.
