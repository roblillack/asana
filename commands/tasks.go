package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/thash/asana/api"
	"github.com/thash/asana/utils"
)

const (
	CacheDuration = "5m"
)

func Tasks(c *cli.Context) {
	var tasks []api.Task_t

	if c.Bool("no-cache") {
		tasks = fromAPI(false)
	} else {
		if utils.Older(CacheDuration, utils.CacheFile()) || c.Bool("refresh") {
			tasks = fromAPI(true)
		} else {
			txt, _ := ioutil.ReadFile(utils.CacheFile())
			_ = json.Unmarshal(txt, &tasks)
			if tasks == nil {
				tasks = fromAPI(true)
			}
		}
	}

	for i, t := range tasks {
		project := ""
		for _, m := range t.Memberships {
			if v, ok := m["project"]; ok {
				project = v.Name
			}
		}
		for _, m := range t.Memberships {
			if v, ok := m["section"]; ok {
				project += fmt.Sprintf(" (%s)", v.Name)
			}
		}
		if len(project) > 0 {
			project += ": "
		}
		fmt.Printf("%3d [ %10s ] %s%s\n", i, t.Due_on, project, t.Name)
	}
}

func fromAPI(saveCache bool) []api.Task_t {
	tasks := api.Tasks(url.Values{}, false)
	if saveCache {
		cache(tasks)
	}
	return tasks
}

func cache(tasks []api.Task_t) {
	f, _ := os.Create(utils.CacheFile())
	defer f.Close()

	utils.Check(json.NewEncoder(f).Encode(tasks))
}
