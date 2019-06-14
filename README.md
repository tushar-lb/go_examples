# filesystem-manager

## Problem 1:
Write a program that does the following:
1. Traverses the file system starting from the root directory, or from a path provided by the user via the command line.
2. Finds all regular files and gets their information (a.k.a file stats, either via os.FileInfo or the standard Unix stat(1) command).
3. Sends this information via POST requests, serialised as JSON, to an HTTP address provided by the user via the command line.


## Guidelines:
Traversal may happen with just one thread (or go-routine), but there must be at least 3 threads sending file information to the HTTP address. This means traversal and sending happen in different threads. The program must provide proper usage information when --help is used.


## Problem 2:
Write another program that does the following:
1. Starts an HTTP server listening on a port defined by the user via the command line, as either an HTTP server or an HTTPS (TLS) server, again depending on the user's input.
2. Maintains a structure of the following file statistics:
     i. Number of files received
     ii. Maximum file size (including file path)
     iii. Average file size
     iv. List of file extensions
     v. Most frequent file extension (including number of occurrences)
     vi. List of latest 10 file paths received
3. Accepts POST requests formatted as JSON to a path of your choosing, which are expected to contain file information (the same information sent in stage 1), and updates the statistics according to the information.
4. Accepts GET requests to a path of your choosing, and returns the (current) statistics as JSON.

## Guidelines:
Since multiple POST requests may be received at the same time, care must be taken that the statistics are updated atomically (that is to say, the effects of one request should not be overridden by another one). The server must properly return errors when provided invalid inputs.

## Packaging Soultions:
1. Modify the server so that the POST endpoint may accept both JSON and XML requests, depending on the
client's choice.
2. Write another program the continuously makes GET requests to the server (at a user-provided interval)
and prints the statistics to standard output.
3. Create Dockerfiles for both executables. You can either create two images, or one image that starts the
appropriate executable based on command line arguments.
4. Create a multi-stage Dockerfile that both builds the executables and creates a Docker image
for them.
