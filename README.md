# SAReq />>/

SAReq, short for "Send A Request", is a command-line HTTP client to help developers and testers send HTTP requests and analyze responses directly from the terminal. It aims to provide a lightweight yet powerful and intuitive environment that minimizes overhead and complexity of the application while maintaining essential developer functionality.

SAReq is built for:
- Developers who are just getting started and want to experiment with HTTP requests in a straightforward way.
- Developers who prefer command-line tools for quick testing and debugging of APIs.
- Solo developers looking for a minimal tool for testing in a local environment.

## Features

- **⚡ Lightweight & Fast:** A small, fast Go-based CLI tool designed to make sending HTTP requests simple and efficient without any extra dependencies or setup.
- **💻 Simple CLI Requests:** Quickly send HTTP requests directly from your terminal using intuitive commands that let you test APIs in seconds.
- **🎨 Prettified Colorful Output:** Response data is displayed in a readable, color-coded format that highlights JSON structure, headers, and status codes for easy inspection.

## Installation

SAReq can be installed via the command line or by downloading the official binary from GitHub Releases.

### Using `go install`

To install SAReq using this method, ensure you have a latest version of [Go](https://go.dev/) installed on your system. Then, run the following command in your terminal:
```sh
go install github.com/im-varun/sareq@latest
```

To verify the installation, you can run:
```sh
sareq --version
# or sareq version
```
This should display the installed version of SAReq.

### Downloading Pre-built Binaries

Pre-built binaries for SAReq are available for various operating systems (Windows, Linux, macOS) and architectures (amd64, arm64), and can be downloaded directly from the [GitHub Releases](https://github.com/im-varun/sareq/releases) page.

Follow these steps to ensure a successful installation:

1. Navigate to the [GitHub Releases](https://github.com/im-varun/sareq/releases) page of SAReq.
2. Download the appropriate archive file for your operating system and architecture along with its corresponding `.sha256` checksum file.
3. Verify the integrity of the downloaded archive using the SHA256 checksum provided in the `.sha256` file.
4. Once the checksum is verified, extract the contents of the downloaded archive to a directory of your choice.
5. Add the directory containing the extracted `sareq` executable to your system's PATH environment variable. This allows you to run SAReq from any terminal window.
6. To verify the installation, open a new terminal window and run:
    
    ```sh
    sareq --version
    # or sareq version
    ```
    This should display the installed version of SAReq indicating that the installation was successful.

### Other Installation Methods

Support for additional installation methods, such as package managers like Scoop and Homebrew, is currently under development. Stay tuned for updates in future releases!

## Quickstart

This quickstart guide covers how to quickly get started using SAReq.

### Introduction

SAReq, short for "Send A Request", is a CLI-based HTTP client that allows you to make HTTP requests directly from your terminal. It is designed to be simple, fast, and easy to use. It provides a straightforward way to interact with APIs without the need for a graphical interface.

### Understanding the Basics

SAReq is built around the concept of commands and flags (also known as options) that you can use to perform various actions. The basic syntax for using SAReq is as follows:
```sh
sareq [command] [flags]
```
The command may also require additional arguments, such as a URL.

### Making Your First Request

#### Request Semantics

SAReq supports several HTTP methods (also known as verbs) that you can use to make requests. The most commonly used methods are:
- `GET`: Retrieve data from a specified resource.
- `POST`: Send data to a server to create a new resource.
- `PUT`: Update an existing resource with new data.
- `PATCH`: Partially update an existing resource.
- `DELETE`: Remove a specified resource.

Each of these methods corresponds to a command of the same name in SAReq. For example, to make a GET request, you would use the `get` command, and to make a POST request, you would use the `post` command.

A request typically consists of the following components:
- **URL**: The endpoint to which the request is sent.
- **Headers**: Additional information sent with the request (e.g., authentication tokens, content type).
- **Body**: The data sent with the request (mainly used with POST, PUT, PATCH requests).

Each HTTP request command in SAReq requires at least a URL as an argument. Additional components such as headers and body can be specified using the `-H/--header` and `-B/--body` flags, respectively.

#### A simple GET request

Let's start by making a simple GET request to fetch data from a hypothetical API endpoint.
```sh
sareq get https://api.example.com/users
```
In this example, we are using the `get` command to retrieve a list of users from the specified URL. SAReq will send the request and display the response in your terminal like this:
```sh
HTTP/1.1 200 OK

Content-Type: application/json; charset=utf-8

[
    {
        "id": 1,
        "name": "John Doe",
        "email": "john@doe.com"
    },
    {
        "id": 2,
        "name": "Jane Smith",
        "email": "jane@smith.com"
    }
]
```

#### A simple POST request

Now, let's make a simple POST request to create a new user.
```sh
sareq post https://api.example.com/users --body '{"name": "Jane Doe", "email": "jane@doe.com"}'
```
In this example, we used the `post` command to send a request to create a new user with the specified name and email. The `--body` flag was used to define the request body containing the user data in JSON format.

### Using flags

SAReq provides several flags that you can use to customize the behavior of the commands. Flags are generally used to customize requests and responses, such as adding headers, specifying request bodies, or changing output formats.

Here are some commonly used flags:
- `-H, --header <header>`: Add a header to include with the request.
- `-B, --body <body>`: Define the body to send with the request (mainly used with POST, PUT, PATCH requests).
- `--timeout <seconds>`: Specify a timeout to use with the request.

There are some other flags available as a convenience for users. These include:
- `-h, --help`: Display help information for the command (or a subcommand) with which it is used.
- `-v, --version`: Display the current version of SAReq.
- `--no-color`: Disable coloring in the HTTP response output.
- `--no-prettify`: Disable prettification in the HTTP response output.

#### Adding Headers

Headers can be added using the `-H` or `--header` flag followed by the header key-value pair in the format `"key=value"`.
```sh
sareq get https://api.example.com/users --header "Authorization=abc123"

# or using the shorthand version
sareq get https://api.example.com/users -H "Authorization=abc123"
```

You can also add multiple headers by using the flag multiple times (once for each header):
```sh
sareq get https://api.example.com/users -H "Authorization=abc123" -H "Accept=application/json"
```

#### Defining a Request Body

Request bodies can be defined using the `-B` or `--body` flag followed by the body content. This is mainly used with POST, PUT, or PATCH requests.
```sh
sareq post https://api.example.com/users --body '{"name": "John Doe"}'

# or using the shorthand version
sareq post https://api.example.com/users -B '{"name": "John Doe"}'
```

The content of the body is enclosed in single quotes or double quotes, depending on type of the content. For example, JSON content is typically enclosed in single quotes to avoid conflicts with double quotes used for JSON keys and values.

#### Specifying a Timeout

To avoid waiting indefinitely for a response, you can specify a timeout with the request using the `--timeout` flag followed by the number of seconds to wait before timing out.
```sh
sareq get https://api.example.com/users --timeout 10
```
This will set a timeout of 10 seconds for the request. If the server does not respond within this time frame, SAReq will terminate the request and display a timeout error.

#### Making a Request with Multiple Flags

SAReq allows you to combine multiple flags in a single request. You simply specify each flag followed by its corresponding value (if applicable) in the command.
```sh
sareq post https://api.example.com/users -H "Authorization=abc123" -H "Content-Type=application/json" -B '{"name": "John Doe"}' --timeout 15
```
This command makes a POST request to create a new user, includes two headers, defines a JSON body, and sets a timeout of 15 seconds.

#### Formatting the Output

Currently, SAReq offers two features to format the output of HTTP responses:
- **Coloring**: This feature adds color to different parts of the HTTP response (e.g., status code, headers, body) to enhance readability.
- **Prettification**: This feature formats the response body (if it's in JSON format) to make it more human-readable by adding indentation and line breaks.

By default, both coloring and prettification are enabled. However, you can disable them using the `--no-color` and `--no-prettify` flags, respectively.
```sh
# to disable only coloring
sareq get https://api.example.com/users --no-color

# to disable only prettification
sareq get https://api.example.com/users --no-prettify

# to disable coloring and prettification together
sareq get https://api.example.com/users --no-color --no-prettify
```

#### Help and Version Information

SAReq provides built-in help and version information through the `-h/--help` and `-v/--version` flags.

To display help information for SAReq or a specific command, use the `-h` or `--help` flag:
```sh
sareq --help

# help for a specific command
sareq get --help
```

To display the current version of SAReq, use the `-v` or `--version` flag:
```sh
sareq --version
```

You can also use the `version` subcommand to get the version information:
```sh
sareq version
```

### Testing with `localhost`

SAReq is not limited to making requests to public APIs. You can also use it to test your local development servers running on `localhost`. For example, if you have a server running on `http://localhost:8080`, you can make requests to it like this:
```sh
# making a GET request to localhost
sareq get http://localhost:8080/api/test

# making a POST request to localhost
sareq post http://localhost:8080/api/test --body '{"data": "test"}'

# making a DELETE request to localhost
sareq delete http://localhost:8080/api/test/1
```

## Contributing

For information on how to contribute to SAReq, please refer to the [CONTRIBUTING.md](CONTRIBUTING.md) document.

## Roadmap

For a detailed roadmap of planned features and improvements, please refer to the [ROADMAP.md](ROADMAP.md) document.

## License

SAReq is licensed under the [BSD-3-Clause license](LICENSE).

## Maintainers

The project is currently maintained by [Varun Mulchandani](https://github.com/im-varun).