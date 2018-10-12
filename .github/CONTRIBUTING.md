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
  4. Clone the respository into `src/github.com/watson-developer-cloud` directory
  5. Use [dep][dep] dependency manager.
  On MacOS you can install or upgrade to the latest released version of [dep][dep] with Homebrew
  ```sh
    $ brew install dep
    $ dep ensure
  ```
  6. Check your code for lint issues
  ```sh
  go get -u gopkg.in/alecthomas/gometalinter.v2
  go get -u golang.org/x/lint/golint
  gometalinter.v2 --errors ./...
  ```
  7. Develop and test your code changes `go test ./...`
  8. Commit your changes:
  * Commits should follow the [Angular commit message guidelines](https://github.com/angular/angular/blob/master/CONTRIBUTING.md#-commit-message-guidelines). This is because our release tool uses this format for determining release versions and generating changelogs. To make this easier, we recommend using the [Commitizen CLI](https://github.com/commitizen/cz-cli) with the `cz-conventional-changelog` adapter.
  9. Push to your fork and submit a pull request to the **master** branch

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
[dep]: https://github.com/golang/dep