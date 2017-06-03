package main

// TweetTable html template
var TweetTable = `
<h1>{{.Handle}}'s Tweets</h1>
<table>
	<tr style='text-align: left'>
		<th>Text</th>
	</tr>
	{{range .List}}
	<tr>
		<td>{{.Text}}</td>
	</tr>
	{{end}}
</table>
`
