package main

// TweetTable html template
var TweetTable = `
<h1><a href='{{.URL}}'>{{.Handle}}'s Tweets</a></h1>
<table>
	<tr style='text-align: left'>
		<th>Date</th>
		<th>Text</th>
		<th>Replies</th>
		<th>Upvotes</th>
		<th>Retweets</th>
	</tr>
	{{range .Tweets}}
	<tr>
		<td><a href='{{.URL}}'>{{.Date}}</a></td>
		<td>{{.Text}}</td>
		<td>{{.Replies}}</td>
		<td>{{.Upvotes}}</td>
		<td>{{.Retweets}}</td>
	</tr>
	{{end}}
</table>
`
