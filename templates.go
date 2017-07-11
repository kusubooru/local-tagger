// generated by go generate; DO NOT EDIT

package main

import "html/template"

var (
	layoutTmpl = template.Must(template.New("layout").Funcs(fns).Parse(layoutTemplate))

	indexTmpl = template.Must(template.Must(layoutTmpl.Clone()).Parse(indexTemplate))

	uploadTmpl = template.Must(template.Must(layoutTmpl.Clone()).Parse(uploadTemplate))

	layoutTemplate = `
{{ define "layout" }}
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Tagaa {{printv .Version}}</title>
    <meta name="description" content="Interface for the 'Bulk Add CSV' Shimmie2 extension">
    <meta name="author" content="kusubooru">

    <style>
      html {
        font-family: sans-serif;
      }
      input {
        margin-bottom: 0.6em;
      }
      .block {
        display: block;
        padding: 15px;
        margin-bottom: 10px;
      }
      .block-danger {
        background: #f2dede;
        color: #333;
      }
      .block-success {
        background: #dff0d8;
        color: #333;
      }
      h1 small {
        font-size:65%;
        color:#777;
      }
      nav {
        margin-bottom: 1em;
      }

      .loader {
        display: none;
        border: 5px solid #f3f3f3;
        border-radius: 50%;
        border-top: 5px solid #006FFA;
        border-right: 5px solid #006FFA;
        width: 32px;
        height: 32px;
        -webkit-animation: spin 1s linear infinite;
        animation: spin 1s linear infinite;
        will-change: transform;
      }
      .loader-small {
        width: 8px;
        height: 8px;
        border-width: 3px;
      }
      @-webkit-keyframes spin {
        0% { -webkit-transform: rotate(0deg); }
        100% { -webkit-transform: rotate(360deg); }
      }
      @keyframes spin {
        0% { transform: rotate(0deg); }
        100% { transform: rotate(360deg); }
      }
    </style>
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/awesomplete/1.1.2/awesomplete.min.css" />
    {{ template "style" . }}

    <!--[if lt IE 9]>
      <script src="http://html5shiv.googlecode.com/svn/trunk/html5.js"></script>
    <![endif]-->
  </head>

  <body>
    <h1>Tagaa <small>{{printv .Version}}</small></h1>
    {{ template "content" . }}
    <script src="https://cdnjs.cloudflare.com/ajax/libs/awesomplete/1.1.2/awesomplete.min.js"></script>
    {{ template "script" . }}
  </body>
</html>
{{ end }}
{{ define "style" }}{{end}}
{{ define "script" }}{{end}}
`

	indexTemplate = `
{{ define "style" }}
  <style>
    #advanced {
      display: none;
    }
    .image {
      max-width: 100%;
    }
    /* Roughly equivalent to cols= 60 and rows = 6 original sizes. */
    .tags-textarea {
      width: 37em;
      height: 7em;
    }
    .medium-input {
      width: 37em;
    }
    .chomp {
      width: 37em;
    }
    .tag {
      color: #0073ff;
    }
    .tag-tk {
      color: #ee5542;
    }
    .tag-artist {
      color: #a00;
    }
    .tag-series {
      color: #a0a;
    }
    .tag-character {
      color: #0a0;
    }
    .tag-count {
      float: right;
    }
    .tag-favicon {
      float: left;
      margin-right: 3px;
      height: 16px;
      width: 16px
    }

    @media (max-width: 768px) {
      /* TODO: Why need this? On mobile view with width 100% the inputs in
       * advanced div exceed max width and scrolling appears. It should be
       * possible to keep all inputs on width: 100% without them exceeding the
       * screen. */
      .chomp {
        width: 98%;
      }
      .tags-textarea {
        width: 100%;
      }
      .awesomplete {
        width: 100%;
      }
      .medium-input {
        width: 100%;
        padding: 0;
      }
    }
  </style>
{{ end }}
{{ define "content" }}
  <nav>
    <a href="/upload">Upload</a>
  </nav>

  {{ if .Err }}
    <div class="block block-danger">
      {{ .Err }}
    </div>
  {{ end }}
  <form action="/load" method="POST" enctype="multipart/form-data">
    <label for="loadCSVFile"><b>Load CSV File</b></label>
    <br>
    <input id="loadCSVFile" name="csvFilename" type="file" accept=".csv" required>
    <input type="submit" value="Load from CSV">
    <button id="toggleButton" type="button">Advanced +</button>
    <br>
  </form>
  <form action="/update" method="POST">
    <div id="advanced">
      <label for="csvFilenameInput"><b>CSV Filename</b></label>
      <br>
      <input id="csvFilenameInput" type="text" name="csvFilename" value="{{ .CSVFilename }}" class="chomp">
      <input id="saveCSVSubmit" type="submit" value="Save to CSV">
      <br>
      <label for="directory"><b>Working Directory</b></label>
      <br>
      <input id="directory" type="text" name="prefix" value="{{ .WorkingDir }}" disabled class="chomp">
      <br>
      <label for="prefixInput"><b>Server Path Prefix</b> (It will replace working directory path prefix)</label>
      <br>
      <input id="prefixInput" type="text" name="prefix" value="{{ .Prefix }}" class="chomp">
      <br>
      <label for="useLinuxSepInput"><b>Use Linux Separator "/" when saving to CSV</b> </label>
      <br>
      <input id="useLinuxSepInput" type="checkbox" name="useLinuxSep" {{if eq .UseLinuxSep true}}checked{{end}}>
      (Check, if working on a windows machine and want to upload to a Linux machine)
      <br>
      <input id="deleteCacheKey" type="text">
      <button id="deleteCacheButton" type="button">Delete Tag from Cache</button>
      <input id="scroll" type="hidden" name="scroll" value="">
    </div>

    <section>
      {{ if .Images }}
        <h2>Images</h2>
      {{ else }}
        <h2>No Images found in local directory</h2>
        Add some and then refresh.
      {{ end }}

      {{ range .Images }}
        <article>
          <fieldset>
            <a id="tags{{ .ID }}"></a>
            <a id="img{{ .ID }}"></a>
            <legend>{{ .Name }}</legend>
            <a href="#img{{ .ID }}"><img class="image" src="/img/{{ .ID }}" alt="{{ .Name }}"></a>
            <br>
            <label for="tagsTextArea{{ .ID }}"><b>Tags</b></label>
            <div id="loader{{ .ID }}" class="loader loader-small"></div>
            <br>
            <textarea id="tagsTextArea{{ .ID }}" data-loader="loader{{ .ID }}" name="image[{{ .ID }}].tags" class="tags-textarea awesomeplete" data-multiple >{{ join .Tags " " }}</textarea>
            <br>
            <label for="sourceInput{{ .ID }}"><b>Source</b></label>
            <br>
            <input id="sourceInput{{ .ID }}" class="medium-input" type="text" name="image[{{ .ID }}].source" value="{{ .Source }}" >
            <br>
            <label><b>Rating</b></label>
            <br>
            <input id="sRadio{{ .ID }}" type="radio" name="image[{{ .ID }}].rating" value="s" {{ if eq .Rating "s" }}checked{{ end }}>
            <label for="sRadio{{ .ID }}">Safe</label>
            <input id="qRadio{{ .ID }}" type="radio" name="image[{{ .ID }}].rating" value="q" {{ if eq .Rating "q" }}checked{{ end }}>
            <label for="qRadio{{ .ID }}">Questionable</label>
            <input id="eRadio{{ .ID }}" type="radio" name="image[{{ .ID }}].rating" value="e" {{ if eq .Rating "e" }}checked{{ end }}>
            <label for="eRadio{{ .ID }}">Explicit</label>
            <br>
            <input class="save-to-csv" type="submit" value="Save to CSV" data-scroll="#tags{{.ID}}">
          </fieldset>
        </article>
        <br>
      {{ end }}
    </section>
  </form>
{{ end }}
{{ define "script" }}
  <script>
    (function(){
      "use strict";

      var csvButtons = document.getElementsByClassName("save-to-csv");
      for(var i=0; i < csvButtons.length; i++) {
          csvButtons[i].onclick = setScroll;
      }
      function setScroll() {
        var scroll = this.getAttribute("data-scroll");
        document.getElementById("scroll").value = scroll;
      }

      var toggleButton = document.getElementById("toggleButton");
      toggleButton.onclick = toggleAdvanced;

      function toggleAdvanced() {
        var b = document.getElementById("toggleButton");
        var div = document.getElementById("advanced");
        // Empty display reverts to CSS rule, in this case none.
        if (div.style.display == '') {
          div.style.display = 'block';
          b.innerHTML = "Advanced -";
        } else {
          div.style.display = '';
          b.innerHTML = "Advanced +";
        }
      }

      // Autocomplete

      var map = {};
      var tas = document.querySelectorAll('textarea[data-multiple]');
      tas.forEach(function(ta){
        var ap = makeAwesomplete(ta);
        map[ta.id] = ap;
        ta.onkeyup = getTagsEventHandler;
      });
      function makeAwesomplete(ta) {
        return new Awesomplete(ta, {
          minChars: 3,
          filter: function(text, input) {
            return Awesomplete.FILTER_CONTAINS(text, input.match(/[^ ]*$/)[0]);
          },

          item: function(text, input) {
            // We have previously stored the tag returned by the server as JSON
            // text in the label in order to access the extra information like
            // count, category and old.
            var item = JSON.parse(text.label);
            var li = document.createElement('li');
            if (item.board && item.board == "kusubooru") {
              var img = document.createElement('img')
              img.src = "/img/kusubooru.ico"
              img.className = "tag-favicon";
              li.appendChild(img);
            }
            if (item.board && item.board == "danbooru") {
              var img = document.createElement('img');
              img.src = "/img/danbooru.ico"
              img.className = "tag-favicon";
              li.appendChild(img);
            }
            if (item.old) {
              var cclass = categoryClass(item.category);
              var old = document.createElement('span');
              old.innerHTML = item.old;
              old.className = cclass;
              var name = document.createElement('span');
              name.innerHTML = item.name;
              name.className = cclass;
              var arrow = document.createElement('span');
              arrow.innerHTML = " → ";
              li.appendChild(old);
              li.appendChild(arrow);
              li.appendChild(name);
            } else {
              var span = document.createElement('span');
              span.innerHTML = item.name;
              span.className = categoryClass(item.category);
              li.appendChild(span);
            }
            if (item.count) {
              var count = document.createElement('span');
              count.innerHTML = item.count;
              count.className = "tag-count";
              li.appendChild(count);
            }
            return li;
          },

          replace: function(text) {
            var before = this.input.value.match(/^.+ \s*|/)[0];
            this.input.value = before + text.value + " ";
          },
          // Set sort function to false to disable sorting. Our backend handler
          // returns items sorted by count (first kusubooru then danbooru).
          sort: false
        });
      }
      function categoryClass(category) {
        switch (category) {
          case "tk":
            return "tag-tk";
          case "artist":
            return "tag-artist";
          case "series":
            return "tag-series";
          case "character":
            return "tag-character";
          default:
            return "tag";
        }
      }

      var timeout = null;
      function getTagsEventHandler(e) {
        var code = (e.keyCode || e.which);
        // https://github.com/LeaVerou/awesomplete/issues/16802#issuecomment-303124988
        if (code !== 37 && code !== 38 && code !== 39 && code !== 40 && code !== 27 && code !== 13) {
          var input = this.value;
          var id = this.id;
          var loaderID = this.getAttribute('data-loader');
          // Wait for user to stop typing before getting tags:
          // https://schier.co/blog/2014/12/08/wait-for-user-to-stop-typing-using-javascript.html
          clearTimeout(timeout);

          timeout = setTimeout(function () {
              getTags(input.match(/[^ ]*$/)[0], id, loaderID);
          }, 500);
        }
      }

      var cache = localStorage;

      var deleteCacheButton = document.getElementById("deleteCacheButton");
      deleteCacheButton.onclick = function(e) {
        var key = document.getElementById("deleteCacheKey").value;
        cache.removeItem(key);
      }

      function getTags(query, apid, loaderID) {
        if (query == "" || query.length < 3) {
          return;
        }
        var hit = cache.getItem(query);
        if (hit) {
          var obj = JSON.parse(hit);
          var now = new Date().getTime();
          if (now < obj.expires) {
            updateTags(obj.value, map, apid);
            return
          }
        }
        var loader = document.getElementById(loaderID);
        loader.style.display = "inline-block";
        var xhr = new XMLHttpRequest();
        xhr.onreadystatechange = function(response) {
          if (xhr.readyState === 4 && xhr.status === 200) {
            loader.style.display = "none";
            var tags = JSON.parse(xhr.responseText);
            var now = new Date().getTime();
            // Get the current time and add (day * hour * min * sec * msec).
            var inOneWeek = now + (7 * 24 * 60 * 60 * 1000);
            var object = {
              value: tags,
              timestamp: now,
              expires: inOneWeek
            }
            cache.setItem(query, JSON.stringify(object));
            updateTags(tags, map, apid);
          }
        };
        xhr.open("GET", "tags?q="+query, true);
        xhr.send();
      }

      function updateTags(tags, apmap, apid) {
        var list=[];
        tags.forEach(function(item) {
          var label = JSON.stringify(item);
          list.push({"label": label, "value": item.name});
        });
        apmap[apid].list = list;
      }

    })();
  </script>
{{ end }}
`
	uploadTemplate = `
{{ define "style" }}
  <style>
    .thumbnail {
      width: 50px;
      height: 30px;
      display: inline-block;
    }
    .thumbnail img {
      width: 100%;
      height: auto;
    }
    .upload-table {
      width: 100%;
    }
    .upload-table textarea {
      width: 95%;
    }
    .upload-button {
      display: inline-block;
      padding: 0.5em;
    }
  </style>
{{end}}

{{ define "content" }}
  <nav>
    <a href="/">Back</a>
  </nav>

  {{ if .Err }}
    <div class="block block-danger">
      {{ .Err }}
    </div>
  {{ else if .Success }}
    <div class="block block-success">
     {{ .Success }}
    </div>
  {{ end }}

  <form action="/upload" method="POST" enctype="multipart/form-data" onsubmit="showLoader()">
    <table class="upload-table">
      <thead>
        <tr>
          <th></th>
          <th>Name</th>
          <th>Tags</th>
          <th>Source</th>
          <th>Rating</th>
        </tr>
      </thead>
      <tbody>
        {{ range .Images }}
          <tr>
            <td>
              <div class="thumbnail">
                <a href="#img{{ .ID }}"><img src="/img/{{ .ID }}" alt="{{ .Name }}" width=150 height=100></a>
              </div>
            </td>
            <td width="10%">
              {{ .Name }}
            </td>
            <td width="65%">
              <textarea id="tagsTextArea{{ .ID }}" name="image[{{ .ID }}].tags" cols="20" rows="2" readonly>{{ join .Tags " " }}</textarea>
            </td>
            <td width="25%">
              {{ .Source }}
            </td>
            <td>
              {{ if eq .Rating "s" }} Safe
              {{ else if eq .Rating "q" }} Questionable
              {{ else if eq .Rating "e" }} Explicit
              {{ else }} Unknown
              {{ end }}
            </td>
          </tr>
        {{ end }}
      </tbody>
    </table>

    <span>The images above are going to be:</span>
    <ul>
      <li>Compressed to a .zip archive</li>
      <li>Uploaded to server Kusubooru.com</li>
      <li>Manually reviewed before posted</li>
    </ul>
    <p>Please make sure that all images have adequate tags, a source and a rating before uploading.</p>

    <p>Use your Kusubooru account to upload:</p>
    <label for="username">Username</label>
    <input id="username" type="text" name="username" placeholder="Username" required>
    <label for="password">Password</label>
    <input id="password" type="password" name="password" placeholder="Password" required>
    <button type="button" onclick="testCredentials()">Test</button>
    <label id="result"></label>
    <div id="testLoader" class="loader loader-small"></div>

    <p><small>(Max file size for a single upload is 50MB and you may upload a total of 200MB per day.)</small></p>
    <input id="uploadButton" class="upload-button" type="submit" value="Upload">
    <div id="uploadLoader" class="loader"></div>
  </form>
{{ end }}
{{ define "script" }}
  <script>
    function showLoader() {
      var loader = document.getElementById("uploadLoader");
      loader.style.display = "inline-block";
      var button = document.getElementById("uploadButton");
      button.style.display = "none";
    }

    function testCredentials() {
      var loader = document.getElementById("testLoader");
      loader.style.display = "inline-block";
      var username = document.getElementById("username").value;
      var password = document.getElementById("password").value;
      var resultLabel = document.getElementById("result");
      resultLabel.innerHTML = "";
      var xhr = new XMLHttpRequest();
      var url = "https://kusubooru.com/suggest/login/test";
      var params = "username="+username+"&password="+password;
      xhr.open("POST", url, true);
      xhr.setRequestHeader("Content-type", "application/x-www-form-urlencoded");
      xhr.onreadystatechange = function() {
        loader.style.display = "none";
        if(xhr.readyState == 4 && xhr.status == 200) {
          resultLabel.innerHTML = "Ok!";
	      } else if(xhr.readyState == 4 && xhr.status != 200) {
	        var reason = "";
	        if (xhr.responseText) { reason = ": " + xhr.responseText }
          resultLabel.innerHTML = "Failed" + reason;
        }
      }
      xhr.send(params);
    };
  </script>
{{ end }}
`
)
