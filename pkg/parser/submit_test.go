package parser

import (
	"strings"
	"testing"
)

func TestCheckIfPageHasLoginForm(t *testing.T) {
	t.Run("it returns true if page contains a Login Form", func(t *testing.T) {
		example := `<form action="action_page.php" method="post">
						<label for="uname"><b>Username</b></label>
						<input type="text" placeholder="Enter Username" name="uname" required>
					
						<label for="psw"><b>Password</b></label>
						<input type="password" placeholder="Enter Password" name="psw" required>
					
						<button type="submit">Login</button>
					</form>`

		contains, _ := CheckIfPageHasLoginForm(strings.NewReader(example))

		assertPageContainsLoginForm(t, true, contains)
	})

	t.Run("it returns false if page doesn't contain a Login Form", func(t *testing.T) {
		example := `<html>
					<title>Home Page</title>
					<body></body>
					</html`

		contains, _ := CheckIfPageHasLoginForm(strings.NewReader(example))

		assertPageContainsLoginForm(t, false, contains)
	})
}
