package main

var twoFactorTopic = &Topic{
	Name:        "twofactor",
	Description: "manage two-factor authentication settings",
}

var twoFactorCmd = &Command{
	Topic:     "twofactor",
	NeedsAuth: true,
	Run: func(ctx *Context) {
		req := apiRequest(ctx.APIToken)
		req.Method = "GET"
		req.Uri = req.Uri + "/account"
		res, err := req.Do()
		ExitIfError(err)
		var doc map[string]bool
		res.Body.FromJsonTo(&doc)
		if doc["two_factor_authentication"] {
			Println("Two-factor authentication is enabled")
		} else {
			Println("Two-factor authentication is not enabled")
		}
	},
}

var twoFactorGenerateCmd = &Command{
	Topic:       "twofactor",
	Command:     "generate-recovery-codes",
	Description: "Generates and replaces recovery codes",
	NeedsAuth:   true,
	Run: func(ctx *Context) {
		req := apiRequest(ctx.APIToken)
		req.Method = "POST"
		req.Uri = req.Uri + "/account/recovery-codes"
		req.AddHeader("Heroku-Password", getPassword())
		req.AddHeader("Heroku-Two-Factor-Code", getString("Two-factor code: "))
		res, err := req.Do()
		ExitIfError(err)
		var codes []interface{}
		res.Body.FromJsonTo(&codes)
		Println("Recovery codes:")
		for _, code := range codes {
			Println(code)
		}
	},
}

var twoFactorDisableCmd = &Command{
	Topic:       "twofactor",
	Command:     "disable",
	Description: "Disable two-factor authentication for your account",
	NeedsAuth:   true,
	Run: func(ctx *Context) {
		req := apiRequest(ctx.APIToken)
		req.Method = "PATCH"
		req.Uri = req.Uri + "/account/"
		req.Body = map[string]interface{}{
			"two_factor_authentication": "false",
			"password":                  getPassword(),
		}
		_, err := req.Do()
		ExitIfError(err)
		Println("disabled two-factor authentication")
	},
}