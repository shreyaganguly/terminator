package main

var html = `<!doctype html>
	<html>
	<head>
	  <title>Alert Logs</title>
		<script src="http://localhost:35729/livereload.js"></script>
	</head>
	<body>
    <h3>Logs matching your filtering keywords {{ .Keywords }}<h3>
    <ul>
	  {{range .Logs}}<li>{{ . }}</li>{{ end }}
    </ul>
	</body>
	</html>`
