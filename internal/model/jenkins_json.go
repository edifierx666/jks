package model

type ParamsBranch struct {
	Class  string        `json:"_class"`
	Errors []interface{} `json:"errors"`
	Values []struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"values"`
}
type JenkinsJobJSON struct {
	Class   string `json:"_class"`
	Actions []struct {
		Class                string `json:"_class,omitempty"`
		ParameterDefinitions []struct {
			Class                 string `json:"_class"`
			DefaultParameterValue struct {
				Class string `json:"_class"`
				Value string `json:"value"`
			} `json:"defaultParameterValue"`
			Description string `json:"description"`
			Name        string `json:"name"`
			Type        string `json:"type"`
		} `json:"parameterDefinitions,omitempty"`
	} `json:"actions"`
	Description       string      `json:"description"`
	DisplayName       string      `json:"displayName"`
	DisplayNameOrNull interface{} `json:"displayNameOrNull"`
	FullDisplayName   string      `json:"fullDisplayName"`
	FullName          string      `json:"fullName"`
	Name              string      `json:"name"`
	URL               string      `json:"url"`
	Buildable         bool        `json:"buildable"`
	Builds            []struct {
		Class  string `json:"_class"`
		Number int    `json:"number"`
		URL    string `json:"url"`
	} `json:"builds"`
	Color      string `json:"color"`
	FirstBuild struct {
		Class  string `json:"_class"`
		Number int    `json:"number"`
		URL    string `json:"url"`
	} `json:"firstBuild"`
	HealthReport []struct {
		Description   string `json:"description"`
		IconClassName string `json:"iconClassName"`
		IconURL       string `json:"iconUrl"`
		Score         int    `json:"score"`
	} `json:"healthReport"`
	InQueue          bool `json:"inQueue"`
	KeepDependencies bool `json:"keepDependencies"`
	LastBuild        struct {
		Class  string `json:"_class"`
		Number int    `json:"number"`
		URL    string `json:"url"`
	} `json:"lastBuild"`
	LastCompletedBuild struct {
		Class  string `json:"_class"`
		Number int    `json:"number"`
		URL    string `json:"url"`
	} `json:"lastCompletedBuild"`
	LastFailedBuild struct {
		Class  string `json:"_class"`
		Number int    `json:"number"`
		URL    string `json:"url"`
	} `json:"lastFailedBuild"`
	LastStableBuild struct {
		Class  string `json:"_class"`
		Number int    `json:"number"`
		URL    string `json:"url"`
	} `json:"lastStableBuild"`
	LastSuccessfulBuild struct {
		Class  string `json:"_class"`
		Number int    `json:"number"`
		URL    string `json:"url"`
	} `json:"lastSuccessfulBuild"`
	LastUnstableBuild     interface{} `json:"lastUnstableBuild"`
	LastUnsuccessfulBuild struct {
		Class  string `json:"_class"`
		Number int    `json:"number"`
		URL    string `json:"url"`
	} `json:"lastUnsuccessfulBuild"`
	NextBuildNumber int `json:"nextBuildNumber"`
	Property        []struct {
		Class                string `json:"_class"`
		ParameterDefinitions []struct {
			Class                 string `json:"_class"`
			DefaultParameterValue struct {
				Class string `json:"_class"`
				Name  string `json:"name"`
				Value string `json:"value"`
			} `json:"defaultParameterValue"`
			Description string   `json:"description"`
			Name        string   `json:"name"`
			Type        string   `json:"type"`
			Choices     []string `json:"choices"`
		} `json:"parameterDefinitions"`
	} `json:"property"`
	QueueItem          interface{}   `json:"queueItem"`
	ConcurrentBuild    bool          `json:"concurrentBuild"`
	Disabled           bool          `json:"disabled"`
	DownstreamProjects []interface{} `json:"downstreamProjects"`
	LabelExpression    interface{}   `json:"labelExpression"`
	Scm                struct {
		Class string `json:"_class"`
	} `json:"scm"`
	UpstreamProjects []interface{} `json:"upstreamProjects"`
}
