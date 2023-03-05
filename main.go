package main

import (
	demo "github.com/saschagrunert/demo"
	"github.com/urfave/cli/v2"
)

func main() {
	d := demo.New()

	d.Name = "A demo of demo"
	d.Usage = "Examples of how to use this"
	d.HideVersion = true

	d.Setup(setup)
	d.Cleanup(cleanup)

	d.Add(simple(), "simple", "Simple commands demo")
	d.Add(longRunning(), "long-run", "Long-running commands demo")
	d.Add(errorProne(), "error-prone", "Error-prone commands demo")
	d.Add(longTyping(), "long-typing", "Commands that require a lot of typing")

	d.Run()
}

func simple() *demo.Run {
	r := demo.NewRun(
		"Simple commands demo",
	)

	r.Step(demo.S(
		"Simple commands:",
	), nil)

	r.Step(nil, demo.S(
		"ps aux | head",
	))

	r.Step(nil, demo.S(
		"ls -l",
	))

	r.Step(nil, demo.S(
		"cat some-file",
	))

	return r
}

func longRunning() *demo.Run {
	r := demo.NewRun(
		"Long-running commands demo",
	)

	r.Step(demo.S(
		"Long running:",
	), demo.S(
		"docker build -t some-app .",
	))

	return r
}

func errorProne() *demo.Run {
	r := demo.NewRun(
		"Error-prone commands demo",
	)

	r.Step(demo.S(
		"Error prone (dependant on external factors like network):",
	), nil)

	r.Step(nil, demo.S(
		"curl --silent -X GET https://httpstat.us/418 -H \"Accept: application/json\" | jq .",
	))

	r.Step(nil, demo.S(
		"curl -i -X POST \"https://httpbin.org/status/204\" --data '{\"some\":\"data\"}'",
	))

	r.Step(nil, demo.S(
		"wget -q --show-progress -O some-file.txt https://raw.githubusercontent.com/octocat/Hello-World/master/README",
	))

	return r
}

func longTyping() *demo.Run {
	r := demo.NewRun(
		"Commands that require a lot of typing",
	)

	r.Step(demo.S(
		"Hard to type:",
	), demo.S(
		"openssl req -x509 -newkey rsa:4096 -sha256 -days 3650 -nodes ",
		"-keyout example.key -out example.crt -subj \"/CN=example.com\"",
		"-addext \"subjectAltName=DNS:example.com,DNS:www.example.net,IP:10.0.0.1\" 2>/dev/null",
	))

	r.Step(nil, demo.S(
		"ls -l | grep example",
	))

	return r
}

func setup(ctx *cli.Context) error {
	// Ensure can be used for easy sequential command execution
	return demo.Ensure(
		"echo 'Here\nis\nsome\ndata\nfor\nyou' > some-file",
	)
}

func cleanup(ctx *cli.Context) error {
	return demo.Ensure(
		"rm some-file some-file.txt || true",
		"docker rmi some-app || true",
		"rm example.crt example.key || true",
	)
}
