# Questions

If you are having difficulties using the APIs or have a question about the IBM Watson Services,
please ask a question on [dW Answers][dw] or [Stack Overflow][stackoverflow].

# Issues

If you encounter an issue with the Go SDK, you are welcome to submit a [bug report](https://github.com/watson-developer-cloud/go-sdk/issues).
Before that, please search for similar issues. It's possible somebody has encountered this issue already.

# Pull Requests

If you want to contribute to the repository, here's a quick guide:
  1. Fork the repository
  2. The GOPATH environment variable is used to specify directories outside of $GOROOT that contain the source for Go projects and their binaries. Example:
  ```sh
  export GOPATH="$HOME/workspace/go-workspace"
  export PATH=$PATH:$GOPATH/bin
  ````
  3. Have the following directory layout
  ```sh
  mkdir $GOPATH/{src,bin,pkg}
  ```
  4. Clone the respository into the `src` directory
  5. `go get` all the package dependencies
  ```sh
  go get -d ./...
  ```
  6. Check your code for lint issues
  ```sh
  go get -u golang.org/x/lint/golint
  golint ./...
  ```
  * ignore error `don't use MixedCaps in package name` and `should not use dot imports`
  7. Develop and test your code changes `go test`
  * Check for unnecessary whitespace with git diff --check before committing
  8. Make the tests pass
  9. Commit your changes
  10. Push to your fork and submit a pull request to the **master** branch

# Developer's Certificate of Origin 1.1

By making a contribution to this project, I certify that:

(a) The contribution was created in whole or in part by me and I
   have the right to submit it under the open source license
   indicated in the file; or

(b) The contribution is based upon previous work that, to the best
   of my knowledge, is covered under an appropriate open source
   license and I have the right under that license to submit that
   work with modifications, whether created in whole or in part
   by me, under the same open source license (unless I am
   permitted to submit under a different license), as indicated
   in the file; or

(c) The contribution was provided directly to me by some other
   person who certified (a), (b) or (c) and I have not modified
   it.

(d) I understand and agree that this project and the contribution
   are public and that a record of the contribution (including all
   personal information I submit with it, including my sign-off) is
   maintained indefinitely and may be redistributed consistent with
   this project or the open source license(s) involved.

## Additional Resources
+ [General GitHub documentation](https://help.github.com/)
+ [GitHub pull request documentation](https://help.github.com/send-pull-requests/)

[dw]: https://developer.ibm.com/answers/questions/ask/?topics=watson
[stackoverflow]: http://stackoverflow.com/questions/ask?tags=ibm-watson