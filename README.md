<!DOCTYPE html>
<html>
<head>
	<title>Timestamp Microservice</title>
	<style type="text/css">
		body {
			width: 50%;
			margin: 0 auto;
			background-color: #f3f3f3;
			line-height: 1.7;
		}
		h1, h2, .example-usage, footer {
			text-align: center;
		}
		code {
			background-color: white;
		}
	</style>
</head>
<body>
	<h1>API Project: Timestamp Microservice</h1>
	<h2>User Stories (WIP):</h2>

	<ol>
		<li>The API endpoint is <code>GET [project_url]/api/timestamp/:date_string?</code></li>
		<li>
			<p>A date string is valid if can be successfully parsed by <code>new Date(date_string)</code>.</p>
			<p>
			    Note that the unix timestamp needs to be an integer (not a string) specifying milliseconds. <br>
			    In our test we will use date strings compliant with ISO-8601 (e.g. "2016-11-20") because this will ensure an UTC timestamp.
			</p>
		</li>
		<li>
			If the date string is empty it should be equivalent to trigger <code>new Date()</code>, i.e. the service uses the current timestamp.
		</li>
		<li>
			If the date string is valid the api returns a JSON having the structure <br>
    		<code>{"unix": &lt;date.getTime()&gt;, "utc" : &lt;date.toUTCString()&gt; }</code> <br>
    		e.g. <code>{"unix": 1479663089000 ,"utc": "Sun, 20 Nov 2016 17:31:29 GMT"}</code>
		</li>
		<li>
			If the date string is invalid the api returns a JSON having the structure <br>
    		<code>{"error" : "Invalid Date" }</code>.
		</li>
	</ol>
    
    <div class="example-usage">
	    <p>Example Usage:</p>
	    <p><a href="/api/timestamp/2015-12-25">/api/timestamp/2015-12-25</a></p>
	    <p><a href="/api/timestamp/1450137600">/api/timestamp/1450137600</a></p>

		<p>Example Output:</p>
		<p><code>{"unix":1451001600000, "utc":"Fri, 25 Dec 2015 00:00:00 GMT"}</code></p>    
	</div>

	[LICENSE](LICENSE)
</body>
</html>