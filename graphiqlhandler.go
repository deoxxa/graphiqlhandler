package graphiqlhandler

import (
	"io"
	"net/http"
	"strings"
)

const htmlContent = `
<!DOCTYPE html>
<html>
  <head>
    <style>body {height: 100vh; margin: 0; width: 100%; overflow: hidden;}</style>
    <link rel="stylesheet" href="//cdn.jsdelivr.net/graphiql/0.10.2/graphiql.css" />
    <script src="//cdn.jsdelivr.net/fetch/0.9.0/fetch.min.js"></script>
    <script src="//cdn.jsdelivr.net/react/15.5.4/react.min.js"></script>
    <script src="//cdn.jsdelivr.net/react/15.5.4/react-dom.min.js"></script>
    <script src="//cdn.jsdelivr.net/graphiql/0.10.2/graphiql.min.js"></script>
    <script>
      document.addEventListener('DOMContentLoaded', function () {
        var endpoint = window.location.origin + '/api/v1/graphql';

        var jwt = prompt("Do you have a JWT you'd like to use?");

        function fetcher(params) {
          return fetch(endpoint, {
            method: 'post',
            headers: {
              'Accept': 'application/json',
              'Content-Type': 'application/json',
              'Authorization': jwt ? 'Bearer ' + jwt : null,
            },
            body: JSON.stringify(params),
            credentials: 'include',
          }).then(function (res) { return res.json() });
        }

        var body = React.createElement(GraphiQL, {fetcher: fetcher, query: null, variables: null});

        ReactDOM.render(body, document.body);
      });
    </script>
  </head>
  <body>
  </body>
</html>
`

func Handler(rw http.ResponseWriter, r *http.Request) {
	if _, err := io.Copy(rw, strings.NewReader(htmlContent)); err != nil {
		panic(err)
	}
}
