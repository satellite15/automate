package config

import (
	"io/ioutil"
	"os"
	"time"

	toml "github.com/pelletier/go-toml"
	log "github.com/sirupsen/logrus"
	rrule "github.com/teambition/rrule-go"

	"github.com/chef/automate/api/interservice/ingest"
	project_update_lib "github.com/chef/automate/lib/authz"
	base_config "github.com/chef/automate/lib/config"
	"github.com/chef/automate/lib/workflow"
)

// TODO @afiune We are unable to use this custom type because the underlying go-toml
// library we use to unmarshal the configuration doesn't support these cases:
// => https://github.com/pelletier/go-toml/blob/master/marshal.go#L341-L345
//
// When using this field and trying to unmarshal we get the following error:
// ```
// ingest-service.default(O): panic: reflect.Set: value of type int is not assignable to type config.job
// ```
//
// type job identifier
//type job int

// Enum jobs
const (
	DeleteNodes int = iota
	NodesMissing
	MissingNodesForDeletion
)

// List of jobs
var JobList = map[int]string{
	DeleteNodes:             "delete_nodes",
	NodesMissing:            "missing_nodes",
	MissingNodesForDeletion: "missing_nodes_for_deletion",
}

// JobConfig is the config for a job
type JobConfig struct {
	// The ID of the job
	ID int `toml:"id"`
	// The threshold time that the job will use internally to do a task
	Threshold string `toml:"threshold"`
	// How often to run the job
	Every string `toml:"every"`
	// Is the job running
	Running bool `toml:"running"`
}

func JobSettingsToUpdateOpts(settings *ingest.JobSettings, oldRecurrence *rrule.RRule) ([]workflow.WorkflowScheduleUpdateOpts, error) {
	err := settings.Validate()
	if err != nil {
		return nil, err
	}

	ret := make([]workflow.WorkflowScheduleUpdateOpts, 0)
	if e := settings.GetEvery(); len(e) > 0 {
		// Convert duration to an rrule
		d, err := time.ParseDuration(e)
		if err != nil {
			// unlikley as validate already checked this
			return nil, err
		}

		r, err := rrule.NewRRule(rrule.ROption{
			Freq:     rrule.SECONDLY,
			Interval: int(d.Seconds()),
			Dtstart:  oldRecurrence.OrigOptions.Dtstart,
		})
		if err != nil {
			return nil, err
		}

		ret = append(ret, workflow.UpdateRecurrence(r))
	}

	if t := settings.GetThreshold(); len(t) > 0 {
		ret = append(ret, workflow.UpdateParameters(t))
	}

	ret = append(ret, workflow.UpdateEnabled(settings.GetRunning()))

	return ret, nil

}

// JobSchedulerConfig - the config for the job scheduler
type JobSchedulerConfig struct {
	// Is the job scheduler running
	Running bool `toml:"running"`
}

func defaultConfig() aggregateConfig {
	return aggregateConfig{
		JobSchedulerConfig: JobSchedulerConfig{
			Running: true,
		},
		JobsConfig: []JobConfig{
			DeleteNodes: JobConfig{
				ID:        DeleteNodes,
				Threshold: "1d",
				Every:     "15m",
				Running:   false,
			},
			NodesMissing: JobConfig{
				ID:        NodesMissing,
				Threshold: "1d",
				Every:     "15m",
				Running:   true,
			},
			MissingNodesForDeletion: JobConfig{
				ID:        MissingNodesForDeletion,
				Threshold: "30d",
				Every:     "15m",
				Running:   true,
			},
		},
		ProjectUpdateConfig: project_update_lib.ProjectUpdateConfig{
			State: project_update_lib.NotRunningState,
		},
	}
}

// Manager - configuration manager for the service
type Manager struct {
	baseConfigManager *base_config.Manager
}

// Config - stores the configuration for the service
type aggregateConfig struct {
	JobsConfig          []JobConfig                            `toml:"jobs_config"`
	JobSchedulerConfig  JobSchedulerConfig                     `toml:"job_scheduler_config"`
	ProjectUpdateConfig project_update_lib.ProjectUpdateConfig `toml:"project_update_config"`
}

// NewManager - create a new config. There should only be one config for the service.
func NewManager(configFile string) (*Manager, error) {
	storedConfig, err := readConfigFromFile(configFile, defaultConfig())
	if err != nil {
		return &Manager{}, err
	}

	// Testing Updating
	baseManager := base_config.NewManager(configFile, storedConfig)
	err = baseManager.UpdateConfig(func(config interface{}) (interface{}, error) {
		return storedConfig, nil
	})
	if err != nil {
		return &Manager{}, err
	}

	return &Manager{
		baseConfigManager: baseManager,
	}, err
}

// Close - close channels
func (manager *Manager) Close() {
	manager.baseConfigManager.Close()
}

// GetProjectUpdateConfig - get the project update config data
func (manager *Manager) GetProjectUpdateConfig() project_update_lib.ProjectUpdateConfig {
	return manager.getConfig().ProjectUpdateConfig
}

// UpdateProjectUpdateConfig - update the project update config
func (manager *Manager) UpdateProjectUpdateConfig(projectUpdateConfig project_update_lib.ProjectUpdateConfig) error {
	return manager.updateConfig(func(config aggregateConfig) (aggregateConfig, error) {
		config.ProjectUpdateConfig = projectUpdateConfig
		return config, nil
	})
}

func (manager *Manager) updateConfig(updateFunc func(aggregateConfig) (aggregateConfig, error)) error {
	return manager.baseConfigManager.UpdateConfig(func(config interface{}) (interface{}, error) {
		return updateFunc(config.(aggregateConfig))
	})
}

func (manager *Manager) getConfig() aggregateConfig {
	aggregateConfig, ok := manager.baseConfigManager.Config.(aggregateConfig)
	if !ok {
		log.Error("baseConfigManager.Config is not of type 'aggregateConfig'")
		os.Exit(1)
	}

	return aggregateConfig
}

func readConfigFromFile(configFile string, defaultConfig aggregateConfig) (interface{}, error) {
	config := defaultConfig

	tomlData, err := ioutil.ReadFile(configFile)
	if os.IsNotExist(err) {
		// config file does not exists use the default config
		return config, nil
	} else if err != nil {
		log.WithFields(log.Fields{
			"config_file": configFile,
		}).WithError(err).Error("Unable to read config file")

		return defaultConfig, err
	}

	err = toml.Unmarshal(tomlData, &config)
	if err != nil {
		log.WithFields(log.Fields{
			"config_file": configFile,
		}).WithError(err).Error("Unable to load manager configuration")

		// Could not load data from config file using the default config.
		return defaultConfig, nil
	}

	return config, nil
}
