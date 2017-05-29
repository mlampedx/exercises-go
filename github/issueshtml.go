package main

// IssueList html template
var IssueList = `
<h1>{{.TotalCount}} issues</h1>
<table>
	<tr style='text-align: left'>
		<th>#</th>
		<th>State</th>
		<th>User</th>
		<th>Title</th>
		<th>Age</th>
	</tr>
	{{range .Items}}
	<tr>
		<td><a href='{{.HTMLURL}}'>{{.Number}}</td>
		<td>{{.State}}</td>
		<td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
		<td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
		<td>{{.CreatedAt | daysAgo}} days</td>
	</tr>
	{{end}}
</table>
`
