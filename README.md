# ChromeDev

`chromedev` is a Go package that allows you to easily launch Google Chrome with remote debugging enabled. It checks if Chrome is already running with the specified debugging port and starts a new instance if it's not.

## CLI Installation

To install the CLI:

```sh
go install github.com/hayeah/chromedev/chromedev@latest
```

It would ensure that the Chrome DevTool is launched, and dump the browser information to `~/.chromedev.json`.

```
cat ~/.chromedev.json
```

```
{
   "Browser": "Chrome/124.0.6367.91",
   "Protocol-Version": "1.3",
   "User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36",
   "V8-Version": "12.4.254.14",
   "WebKit-Version": "537.36 (@51df0e5e17a8b0a4f281c1665dbd1b8a0c6b46af)",
   "webSocketDebuggerUrl": "ws://127.0.0.1:9222/devtools/browser/32956d32-9318-4499-a83e-af1038ebf652"
}
```
