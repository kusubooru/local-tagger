// generated by go generate; DO NOT EDIT

package main

import "html/template"

var (
	layoutTmpl = template.Must(template.New("layout").Funcs(fns).Parse(layoutTemplate))

	indexTmpl = template.Must(template.Must(layoutTmpl.Clone()).Parse(indexTemplate))

	layoutTemplate = `
{{ define "layout" }}
<!DOCTYPE html>
<html lang="en">

<head>
<meta charset="utf-8">

<title>Bulk Add CSV UI</title>
<meta name="description" content="Interface for the Bulk Add CSV Shimmie2 extension">
<meta name="author" content="kusubooru">

<style>
html {
    font-family: sans-serif;
}
input {
    margin-bottom: 0.6em;
}
#err {
    background: #f2dede;
    display: block;
    padding: 15px;
    margin-bottom: 10px;
    color: #333;
}
</style>
{{ template "style" . }}

<!--[if lt IE 9]>
	    <script src="http://html5shiv.googlecode.com/svn/trunk/html5.js"></script>
    <![endif]-->
</head>

<body>
{{ template "content" . }}
<script></script>
{{ template "script" . }}
</body>

</html>
{{ end }}
{{ define "style" }}{{end}}
{{ define "script" }}{{end}}
`

	indexTemplate = `
{{ define "content" }}

{{ $inputSize := 60 }}
{{ $taRows := 6 }}
{{ if .Err }}
<div id="err">
	{{ .Err }}
</div>
{{ end }}
<form action="/load" method="POST" enctype="multipart/form-data">
	<label for="loadCSVFile"><b>Load CSV File</b></label>
	<br>
	<input id="loadCSVFile" name="csvFilename" type="file" accept=".csv" required>
	<input type="submit" value="Load from CSV">
	<br>
</form>
<form action="/update" method="POST">
	<label for="csvFilenameInput"><b>CSV Filename</b></label>
	<br>
	<input id="csvFilenameInput" type="text" name="csvFilename" value="{{ .CSVFilename }}" size="{{ $inputSize }}">
	<input id="saveCSVSubmit" type="submit" value="Save to CSV">
	<br>
	<label for="directory"><b>Local Directory</b></label>
	<br>
	<input id="directory" type="text" name="prefix" value="{{ .Dir }}" disabled size="{{ $inputSize }}">
	<br>
	<label for="prefixInput"><b>Server Path Prefix</b> (It will replace local directory prefix)</label>
	<br>
	<input id="prefixInput" type="text" name="prefix" value="{{ .Prefix }}" size="{{ $inputSize }}">

	<section>
		{{ if .Images }}
		<h2>Images</h2>
		{{ else }}
		<h2>No Images found in local directory</h2>
		Add some and then refresh.
		{{ end }}

		{{ range .Images }}

		<article>
			<h3>{{ .Name }}</h3>
			<img src="/img/{{ .ID }}" alt="{{ .Name }}">


			<fieldset>
				<legend>{{ .Name }}</legend>
				<label for="tagsTextArea{{ .ID }}"><b>Tags</b></label>
				<br>
				<textarea id="tagsTextArea{{ .ID }}" name="image[{{ .ID }}].tags" cols="{{ $inputSize }}" rows="{{ $taRows }}">{{ join .Tags " " }}</textarea>
				<br>
				<label for="sourceInput{{ .ID }}"><b>Source</b></label>
				<br>
				<input id="sourceInput{{ .ID }}" type="text" name="image[{{ .ID }}].source" value="{{ .Source }}" size="{{ $inputSize }}">
				<br>
				<label><b>Rating</b></label>
				<br>
				<input id="sRadio{{ .ID }}" type="radio" name="image[{{ .ID }}].rating" value="s" {{ if eq .Rating "s" }}checked{{ end }}>
				<label for="sRadio{{ .ID }}">Safe</label>
				<input id="qRadio{{ .ID }}" type="radio" name="image[{{ .ID }}].rating" value="q" {{ if eq .Rating "q" }}checked{{ end }}>
				<label for="qRadio{{ .ID }}">Questionable</label>
				<input id="eRadio{{ .ID }}" type="radio" name="image[{{ .ID }}].rating" value="e" {{ if eq .Rating "e" }}checked{{ end }}>
				<label for="eRadio{{ .ID }}">Explicit</label>
			</fieldset>
		</article>

		{{ end }}
	</section>
</form>
{{ end }}
`
)
