package actions

import (
	"regexp"
)

func (as *ActionSuite) Test_AssetURL_AddsContentVersion() {
	url := assetURL("assets/app.css")
	as.Regexp(regexp.MustCompile(`^/assets/app\.css\?v=[0-9a-f]{12}$`), url)
	as.Equal(url, assetURL("/assets/app.css"))
}

func (as *ActionSuite) Test_AssetURL_MissingFile() {
	as.Equal("/assets/no-existe.css", assetURL("assets/no-existe.css"))
}

func (as *ActionSuite) Test_Layout_UsesVersionedAssets() {
	res := as.HTML("/").Get()
	as.Equal(200, res.Code)
	as.Regexp(regexp.MustCompile(`href="/assets/app\.css\?v=[0-9a-f]{12}"`), res.Body.String())
	as.Regexp(regexp.MustCompile(`src="/assets/app\.js\?v=[0-9a-f]{12}"`), res.Body.String())
}
